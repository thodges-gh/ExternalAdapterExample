# External Adapter Example

This example originated from [this](https://read.acloud.guru/serverless-golang-api-with-aws-lambda-34e442385a6a) Medium post that explains how to set up a Go function using AWS Lambda.

This repository should serve as an example for making external adapters available for the ChainLink network. At a minimum, developers should provide basic instructions for installing the external adapter. Additional information, like what is contained in the headings listed below, will further help developers and users of the external adapter.

## API Resource

Place information on account creation for the API here.

## Testing

Testing instructions go here.

## Build & Package

First, clone this repository. Then, enter into the directory

```bash
cd ExternalAdapterExample
```

Build the binary

```bash
GOOS=linux go build -o bin/main src/main.go
```

Zip the binary

```bash
zip deployment.zip bin/main
```

## Installation

Instructions for provider-specific (if adapter is not generic) installation goes here.
