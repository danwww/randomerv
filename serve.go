package main

import (
    "os"
    "io"
    "log"
    "path"
    "time"
    "strconv"
    "net/http"
    "math/rand"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
)

type Config struct {
    FileDir    string
    StaticPath string
    RandPath   string
    ListenAddr string
    ListenPort string
    Exts       []string
}

// Let's get this server running
func main() {

    // Let's call our config set
    set := LoadConfig()    
    print("Starting server...\n")

    // Static file server handle
    http.Handle(set.StaticPath, http.StripPrefix(set.StaticPath, http.FileServer(http.Dir(set.FileDir))))

    // Random file server handle
    http.HandleFunc(set.RandPath, random_file)

    // Begin
    print("Server started. Listening on: " + set.ListenAddr + set.ListenPort + "\n")
    log.Fatal(http.ListenAndServe(set.ListenAddr + set.ListenPort, nil))
}

// Loads JSON Config file from config.txt
func LoadConfig() *Config {
    config_data, err := ioutil.ReadFile("config.txt")

    // Load defaults first
    c := &Config{"images", "/image/", "/random", "0.0.0.0", ":9090", []string{".jpg", ".jpeg"}}

    // If config.txt doesn't exist, err will be something other than nil
    // Let's create one
    if err != nil {
        print("Creating config file...\n")
        config_json, err := json.MarshalIndent(c, "", "    ")
        handle_err(err)
        new_config, err := os.Create("config.txt")
        _, e := new_config.Write(config_json)
        handle_err(e)
        new_config.Close()

    // err = nil therefore config.txt exists. Let's load it
    // TODO: Check for legal params in config.txt
    } else {
        print("Loading config file...\n")
        err = json.Unmarshal(config_data, c)
        handle_err(err)
    }
    return c
}

// This function writes a requet from raw data
// TODO: Handle more MIME types than image/...
func random_file(w http.ResponseWriter, r *http.Request) {

    // Load config before each request.
    // Config can be edited hot
    set := LoadConfig()
    file := pick_random(set.FileDir, set.Exts)
    filepath := path.Join(set.FileDir, file.Name())
    print("Request received for: " + filepath + "\n")

    // Set proper HTTP headers for raw files
    // TODO: autodetect MIME type
    w.Header().Set("Content-type", "image/jpeg")
    w.Header().Set("Content-length", strconv.FormatInt(file.Size(), 10))
    f, err := os.Open(filepath) 
    handle_err(err)
    data := make([]byte, 128)
    for {
        count, err := f.Read(data)
        if err != nil && err != io.EOF {
            handle_err(err)
        }
        if count == 0 {
            break
        }
        w.Write(data[:count])
    }
    return
}

// Returns random number from (min, max) range
func rand_range(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// Checks filelist against list of file extensions
// Uses values from config
func check_ext(file os.FileInfo, exts []string) (is_match bool) {
    for _, ext := range exts {
        if filepath.Ext(file.Name()) != ext {
            is_match = false
        } else {
            is_match = true
            break
        }
    }
    return
}

// Gets list of files fom FileDir in config
func get_files(dir string, exts []string) (files []os.FileInfo) {
    dirlist, _ := ioutil.ReadDir(dir)
    for _, item := range dirlist {
        if item.IsDir() == false {
            // Checks files against Exts in config
            if check_ext(item, exts) == true {
                files = append(files, item)
            }
        }
    }
    return
}

// Picks a random file fom dir
func pick_random(dir string, exts []string) os.FileInfo {
    filelist := get_files(dir, exts)
    max := len(filelist)
    rand := rand_range(0, max)
    return filelist[rand]
}

// Generic error handler. Quits on error
// TODO: Improve error handling
func handle_err(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}
