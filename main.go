package main

/////////////
// Main.go //
/////////////

const (
	Usage = `Start a web server to display current usage of labs.

Usage:
  dashboard [options]
  dashboard --version
  dashboard -h | --help

Options:
  --debug                                             Turn on debugging output.
  -b, --bind=(<interface>:<port>|<interface>|:<port>) Set the interface and port for the server.
  -c, --config=<file>                                 Specify a configuration file.
  -i, --interval=(<sec>s|<min>m|<hr>h)`

	Version = "dashboard 1.0.0"
)

func main() {

	var configs Config
	configs.Configure()

	RunServer(configs)
}

