# brick

This is a simple command line weather application. The name `brick` comes from the weatherman, Brick Tamland, from the Anchorman movies.

### Usage
```console
➜  brick git:(master) $ ./brick 
A weather forecaster

Usage:
  brick [command]

Available Commands:
  forecast    Show the weather forecast
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/brick.yaml)
  -h, --help            help for brick
      --imperial        Use imperial system instead of metric
      --textual         Show textual description instead of icons

Use "brick [command] --help" for more information about a command.

➜  brick git:(master) $ ./brick forecast
Show the weather forecast

Usage:
  brick forecast [command]

Available Commands:
  now         Show the current weather forecast

Flags:
  -h, --help   help for forecast

Global Flags:
      --config string   config file (default is $HOME/brick.yaml)
      --imperial        Use imperial system instead of metric
      --textual         Show textual description instead of icons

Use "brick forecast [command] --help" for more information about a command.
```

### Examples
```console
# Icons use nerd fonts
➜  brick git:(master) $ ./brick forecast now
️ 12.83°C

# Also available as a textual output
➜  brick git:(master) $ ./brick forecast now --imperial --textual
Clouds 55.09°F
```
