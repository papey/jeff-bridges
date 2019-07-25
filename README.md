# Jeff-Bridges

Multiply bridges, for fun and profit

## Getting Started

### Prerequisites

- [Golang](https://golang.org)

### Installing

Get deps using go mod

```sh
go mod vendor
```

then build

```sh
cd cmd
go build jb.go
```

### Usage

Jeff-Bridges can add or delete bridge

#### Add

For example :

```sh
jb add test 127.17.0.1/16
```

Will add a new bridge name test with next available ip. In this case, ip will be 127.18.0.1 if not used by another iterface.

#### Delete

For example :

```sh
jb delete test
```

Will delete bridge named test

## Authors

- **Wilfried OLLIVIER** - *Main author* - [Papey](https://github.com/papey)

## License

[LICENSE](LICENSE) file for details

## Acknowledgments

- This project uses gitflow, only for learning purpose
