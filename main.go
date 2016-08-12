package main

/////////////
// Main.go //
/////////////

/*
Usage: Start a web server to display current usage of labs.

Usage:
	dashboard [options]
	dashboard --version
	dashboard -h | --help

Options:
	--debug
  -b, --bind=(<interface>:<port>|<interface>|:<port>) Set the interface and port for the server.
  -c, --config=<file>                                 Specify a configuration file.
  -i, --interval=(<sec>s|<min>m|<hr>h)`
*/

func main() {
	settings := new(Config)
	settings.Configure()
	runServer(settings)
}
