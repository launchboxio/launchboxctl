module github.com/launchboxio/launchboxctl

go 1.21.1

replace github.com/launchboxio/launchbox-go-sdk => ../launchbox-go-sdk

require (
	github.com/dghubble/sling v1.4.1
	github.com/launchboxio/launchbox-go-sdk v0.0.0-20231106152915-ad0a439ede29
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/go-resty/resty/v2 v2.10.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.17.0 // indirect
)
