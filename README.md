# Dashboard

Start a web server to display current usage of labs.

# Preconditions
1. `dashboard` requires `go` to be installed.

# Installation
`...`





# Usage
`dashboard` is the command used to run the server. It requires a configuration file called `config.json` that is specified
as an environment variable or with command-line flags.


### `config.json` Format

```json
{
  "interface": "localhost",
  "port": "8080",
  "interval": "10s",
  "machineRanges": [
    {
      "prefix": "***REMOVED***",
      "start": 0,
      "end": 25
    },
    {
      "prefix": "***REMOVED***",
      "start": 0,
      "end": 25
    },
    {
      "prefix": "***REMOVED***",
      "start": 1,
      "end": 20
    },
    {
      "prefix": "***REMOVED***",
      "start": 0,
      "end": 21
    },
    {
      "prefix": "***REMOVED***",
      "start": 1,
      "end": 20
    },
    {
      "prefix": "***REMOVED***",
      "start": 1,
      "end": 25
    },
    {
      "prefix": "linux",
      "start": 1,
      "end": 12
    }
  ],
  "machineIdentifiers": [
    {
      "name": "linux",
      "port": "***REMOVED***",
      "color":"MediumSeaGreen"
    },
    {
      "name": "windows",
      "port": "***REMOVED***",
      "color": "MediumSlateBlue"
    }
  ]
}
```

The settings in `config.json` are the defaults for the program. These can be overwritten on a session-by-session or invokation-by-invokation
basis, using environment variables and command-line flags respectively. The `machineRanges` list specifies the ranges of machines that the
program will treat as one grouping. `machineIdentifiers` determine which ports the program will check and what status each port is associated with,
as well as what color will show up on the page.

### Command-line Flags
```
	dashboard [options]
	dashboard --version
	dashboard -h | --help

Options:
	--debug
  -b, --bind=(<interface>:<port>|<interface>|:<port>) Set the interface and port for the server.
  -c, --config=<file>                                 Specify a configuration file.
  -i, --interval=(<sec>s|<min>m|<hr>h)
```

### Environment Variables
```
DASHBOARD_DEBUG = false
DASHBOARD_CONFIG=/path/to/config
DASHBOARD_BIND=localhost:8080 | localhost | :8080
DASHBOARD_INTERVAL=10s|10m|10h
```
