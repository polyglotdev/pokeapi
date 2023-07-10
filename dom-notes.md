
# Dom's Notes

## Resources 🧰

> 🧠 Use a code-generator to generate the SDK.

- Language-agnostic API Design(DSL, JSON, YAML)
  - Language Specific generators
    - Dropbox Stone
      - [Stone](https://domslinks.com/3NRUbCW)
    - Google
      - [JSON Spec](https://domslinks.com/3DcJMfP)
      - [Codegen](https://domslinks.com/44x4xP4)
    - AWS SDK
      - [Git Repo](https://domslinks.com/43qQBp6)
    - OpenAPI
      - [OpenAPI Code Gen](https://domslinks.com/3D7hxPS)

## Steps for getting this "right" 📝

1. Generate the SDK
2. Avoid Surprises 🥳
   1. Principle of **LEAST** Surprise
   2. Avoid external dependencies
   3. Avoid vendored dependencies
   4. What you need is what you get
      1. Scoped sub-packages for larger API’s
      2. Make it `go get` friendly
3. Make the Configuration Straightforward 🤓
   1. Use a `Config` struct
   2. Use a `New` function
   3. Use a `With` function
   4. Use a `Client` struct
   5. Don't use a `cmdline` flags
   6. Environment variables are ok but 'thar be dragons 🐉'
   7. Persistance outside SDK
      1. Let application choose
      2. dbxcli load/stores json
      3. configs
4. Provide Visibility
   1. Allow verbose logging
   2. Allow configurable logging targets
   3. Limited by `stdlib`
5. Use Go idioms for unsupported types
6. Inherited Types
7. Do not authenticate in the SDK
   1. Apps authenticate, SDK accepts accepts OAuth token
   2. Beware of OAuth pitfalls
      1. changed endpoints
      2. app failing on token refresh
      3. Gem from `oauth/internal.token.go`
8. Auto Generate Tests
   1. Tests should be comprehensive and up-to-date
   2. Model after AWS SDK
   3. Not a panacea, still need to test. _Especially_ negative test cases
9. Handle Errors the Go way
   1. Have SDK errors implement `error` interface
   2. Use type assertions to get more info and extract error codes
   3. Good packages exist for combining errors

```go
type Config struct {
    // OAuth2 token
    Token string
    // Enable verbose output
    Verbose bool
    // For testing purposes
    Client *http.Client
}
```
```go
// You need to check AWS SDK for Go on their logging interface
// It provides a better logging level! 🛠️

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new config with verbose logging enabled
  c := &Config{
    Verbose: true,
  }
  // Create a new logger that writes to stdout
  c.Logger = log.New(os.Stdout, "", 0)

  // Try to log something
  c.TryLog("Hello, world!\n")
}

type Config struct {
	// Enable verbose logging
	Verbose bool `yaml:"verbose"`
	// optional logging target
	Logger *log.Logger `yaml:"-"`
}

// TryLog will, if Verbose is set, log to the configured logger.
// If no logger is configured, it will log to os.Stderr.
func (c *Config) TryLog(format string, args ...interface{}) {
	if c.Verbose {
		if c.Logger != nil {
			c.Logger.Printf(format, args...)
		} else {
			fmt.Fprintf(os.Stderr, format, args...)
		}
	}
}

```
```output
Hello, world!
```

## TODO'S 📝

- [ ] Create V2 of SDK
  - [ ] Move the `pokeapi.go` file to `v2` directory
  - [ ] Create the rest of the SDK in the `v2` directory based on [AWS Go SDK](https://domslinks.com/43qQBp6)
- [ ] Create a [self-managed SDK](https://domslinks.com/44iVS2W) and then deploy it to a private repo
- [ ] Clean up the docs and get ready show off the SDK

After you are done, send email with info on how to access the SDK to the team.

