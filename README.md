<h1>randomserv</h1>
<p>Picks a random file out of a folder and serves it over HTTP</p>
<h5>TODO: Serve files other than images</h5>  
<p>Configurable using config.txt</p>
| Option | Function|
|-------:|:--------|
| *FileDir* | Files in this folder will be served | 
| *StaticPath* | Static access to folder |
| *RandPath* | Random file from folder |
| *ListenPort* | Sever will listen on this port (Note: ports below 1000 are often locked for users other than root) |
| *ListenAddr* | Server will listen on this address (Leave at 0.0.0.0 for all addesses) |
| *Exts* | Allowed File types that will be served (To serve anything but images, change mime type header on line 84) |
