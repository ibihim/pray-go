# Pray-go

Get the next prayer time from https://api.aladhan.com/

This means it doesn't calculate anything, but picks times from a free API.

Don't trust it blindly, bugs are possible.
I don't take blame for missed prayers.

## Usage

```bash
  prayer next [flags]

Flags:
      --city string     City name (default "Berlin")
  -h, --help            help for next
      --list-files      List the location of the cache file
      --nation string   Country name (default "Germany")
      --no-cache        Do not cache the prayer times
      --no-newline      Do not print a newline at the end
```
Usage: make [TARGET]

## Install

You currently need a local Go development environment.

```
Targets:
  help        Show this help message
  build       Build the prayer binary
  install     Install the prayer binary to /home/ibihim/go/bin
  test        Run tests
  run         Run the application
  watch-test  Watch for file changes and run tests
  watch-run   Watch for file changes and run the application
  clean       Clean up the prayer binary
```

