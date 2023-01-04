# pullpigo ![CI badge](https://github.com/nicokosi/pullpigo/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/nicokosi/pullpigo)](https://goreportcard.com/report/github.com/nicokosi/pullpigo)

Pullpigo is a command-line that displays Pull Request counters for GitHub repositories.

## Pre-requisite

Install the [`go command`](https://golang.org/cmd/go/).

## Install

    go install github.com/nicokosi/pullpigo

## Run

    pullpigo --repo=nicokosi/pullpigo

For instance, here is an output example for the GitHub repository <https://github.com/vidal-community/atom-jaxb>:

    pullpigo --repo=vidal-community/atom-jaxb
    GitHub repository 'vidal-community/atom-jaxb'
    Pull requests
        opened per author
            amairi: 1
        commented per author
            AElMehdiVidal: 1
            jcgay: 1
        closed per author

The available command options can be listed this way:

    pullpigo --help

## Code

* run a command: `go run pullpigo.go --repo=nicokosi/pullpigo`
* run the tests: `go test`
* format the code: `go fmt`
