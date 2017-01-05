#YodaSays-Slack
A slack bot that dispenses yodaisms.

## Motivation
The goal of the bot is to be a simple exploration of building a slack bot with the nlopes/slack api wrapper.
It is easily modified to produce quotes of other wise guys e.g., old coworkers.


## Build and run
Make sure that GOPATH is set and the go dependencies are installed:

```
go get "github.com/nlopes/slack"
go get "github.com/spf13/viper" 
```

Put your slack api token the bot should use in the `config.toml` file. 
Run the standard `go build` and run the produced binary.

## Contributions
Any pull requests are welcome. 

## Licence
This code is provided under the MIT licence.
