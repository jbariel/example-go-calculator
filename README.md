# Golang calculator example
Build out a calculator in golang to show structure and containerization.

## Build the code
From the root dir:

`./scripts/build.sh`

## Build docker image
`docker build -t 'calc:latest' .`

## Run it...
`docker run --rm calc:latest`

# Based on...
This is meant to be an idiomatic go implementation of [the java version](https://github.com/jbariel/example-java-calculator)

# License
This is licensed under Apache-2.0.