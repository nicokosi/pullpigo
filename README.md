# pullpigo ![CI badge](https://github.com/nicokosi/pullpigo/workflows/CI/badge.svg)

Pullpigo is a command-line that displays Pull Request statistics for GitHub repositories.

## Build

    go build pullpigo.go

## Run

Either:

    go run pullpigo.go -repo=nicokosi/pullpigo

Or if `go build` has already been called:

    ./pullpigo -repo=nicokosi/pullpigo

If an [access token](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line) is required:

    ./pullpigo -repo=nicokosi/pullpigo -token=$GITHUB_ACCESS_TOKEN

For instance, here is output example for the GitHub repository https://github.com/vidal-community/atom-jaxb:

    ./pullpigo -repo=vidal-community/atom-jaxb
    GitHub repository 'vidal-community/atom-jaxb'
    1 events created by fchetouani
    6 events created by AElMehdiVidal
    3 events created by jcgay
    2 events created by amairi

Display the "usage" (the available command options):

    ./pullpigo --help

## Code

After code changes, format the code:

    go fmt

Run tests:

    go test
