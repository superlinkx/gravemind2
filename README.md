#Gravemind API service for Keystroke

## Getting Started
You must have golang installed and configured for your system. See golang.org for details.
Currently Gravemind2 is in active development. To set up your dev environment:
- In your `$GOPATH/src/github.com/superlinkx/` directory: `git clone -b gravemind2 https://github.com/superlinkx/gravemind`
- In the `gravemind` directory: `go get`
- Dependencies should now be ready to go, modify as needed
- Create a copy of `gravemind.json` called `configs/gravemind-dev.json` and populate with test settings
- When ready to run, `go install` and `gravemind -config configs/gravemind-dev.json`
- Server will be running on `localhost:9999` by default

## Running Tests
`go test`
