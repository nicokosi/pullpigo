# pullpigo ![CI badge](https://github.com/nicokosi/pullpigo/workflows/CI/badge.svg)

Pullpigo is a command-line that displays Pull Request statistics for GitHub repositories.

## Build

    go build pullpigo.go

## Run

Via the [`go command`](https://golang.org/cmd/go/):

    go run pullpigo.go --repo=nicokosi/pullpigo

Or via the executable previously generated:

    ./pullpigo --repo=nicokosi/pullpigo

For instance, here is an output example for the GitHub repository https://github.com/vidal-community/atom-jaxb:

    ./pullpigo --repo=vidal-community/atom-jaxb
    GitHub repository 'vidal-community/atom-jaxb'
    Pull requests
        opened per author
            amairi: 1
        commented per author
            AElMehdiVidal: 1
            jcgay: 1
        closed per author

The available command options can be listed this way:

    ./pullpigo --help

## Code

After code changes, format the code:

    go fmt

Run tests:

    go test
