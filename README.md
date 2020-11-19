# Golang calculator example
Build out a calculator in golang to show structure and containerization.

## Build the code
From the root dir:

`./scripts/build.sh ./calculator`

## Build docker image
`docker build -t 'calc:latest' .`

## Run it...
`docker run --rm calc:latest <action> <initial value> <values...>`

where:
- `<action>` represents one of `'+'`, `add`, `'-'`, `subtract`, `'*'`, `multiply`, `'/'`, `divide` (case insensitive)
- `<initial value>` represents the initial value to start from
- `<values...>` represents an arbitrary number of values to apply to the `<initial value>` using the `<action>`

```
  docker run --rm calc:latest add 0 1 2         // => 3
  docker run --rm calc:latest subtract 10 1 1   // => 8
  docker run --rm calc:latest multiply 1 1 2 3  // => 6
  docker run --rm calc:latest divide 1 2        // => 0.5
```

### Alias
Probably easiest to alias:
```
  alias calc='docker run --rm calc:latest'
```
Then it's as simple as:
```
  calc '+' 0 1 2    // => 3
  calc '-' 10 1 1   // => 8
  calc '*' 1 1 2 3  // => 6
  calc '/' 1 2      // => 0.5
```
# Based on...
This is meant to be an idiomatic go implementation of [the java version](https://github.com/jbariel/example-java-calculator)

# License
This is licensed under Apache-2.0.
