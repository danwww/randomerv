randomerv
=========
Serves a random file from a directory

TODO: Serve files other than images

Configurable using config.txt 
|Option|Function|
|-----:|:-------|
|FileDir|Directory from which random file is chosen|
|StaticPath|URL from which files can be accessed directly|
|RandPath|URL from which random file will be served|
|ListenAddr|Sever will listen on this address (Leave at 0.0.0.0 for all addesses)|
|ListenPort|Server will listen on this port (Note: ports below 1000 are often locked for users other than root)|
|Exts|Allowed file extensions|
