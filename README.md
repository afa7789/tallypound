# TallyPound - CompoundAPI usage
![!Theme Image](resources/compound.png)

## Take-Home Programn Assignment for Tally using Compound API

A take home assignment in Golang.
API interaction with Compound.

### The Problem: 

We would like to make an API that uses the Compound proposals endpoint to fetch data
about their governance proposals.
Compound is an Ethereum lending protocol. Users lock some amount of one asset, like
ETH, to borrow some amount of another one, like DAI. Compound uses a token called
COMP and a smart contract called Governor to manage changes to their lending
protocol. COMP tokenholders can make changes by sending proposals to Governor
smart contract. If they pass a vote, the smart contract implements the changes.

### Goal: 

Build an API server that returns stats about Compound's proposals using the data from their API. 
Create a proposal stats endpoint that returns aggregate stats about all of Compound's proposals.
Return the number of proposals in each state. 

The states are pending , active , canceled , defeated , succeeded , queued , expired , executed.

To keep things simple, avoid adding additional dependencies on stateful services like
databases.

## Install & Run

```sh
    git clone https://github.com/afa7789/tallypound
    cd tallypound
    go run .
```

## Testing project

``` sh
    make test
```

## Linting the project

``` sh
    make lint
```


## Tree, Project Structure

```
.
├── cmd // initial flow and start of the programn
│   └── server.go
├── go.mod
├── go.sum
├── internal
│   ├── cache // caching package
│   │   └── cache.go
│   ├── compound // bridge that communicates with compound api package
│   │   ├── compound.go
│   │   └── compound_test.go
│   ├── domain // shared data between packages
│   │   ├── compound.go
│   │   └── flags.go
│   └── server // routes
│       ├── compound.go
│       ├── compound_test.go
│       ├── server.go
│       └── server_test.go
├── main.go
├── Makefile
├── README.md
└── resources
    └── compound.png

7 directories, 16 files

```

## Tools used:

- Markdown , for readme, rofl
- Lint
- Unit Test 
- MakeFile
- [CI , github actions](https://github.com/afa7789/tallypound/actions)
