# Shodan Query Tool
Basic tool to query shodan, original code taken from "blackhat go" by Tom Steele, Chris Patten, and Dan Kottmann.

## Options
- `-i` print api-usage and exit
- `-s <server>` host/server to query shodan about

## Install
Modify the API key var in the `main.go` file.
```go
const apiKey string = "{Your API Key Here}"
```
Then run `sudo make install` and the binary will be placed at `/usr/bin/shodan`.
