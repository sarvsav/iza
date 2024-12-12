# 2. Implement cobra-cli for command line utility

Date: 2024-12-12

## Status

Accepted

## Context

The command line framework cobra-cli is a popular choice for building command line utilities in Go. It provides a simple and easy to use interface for building command line applications.

## Decision

The cobra-cli framework will be used to build the command line utility for the iza project.

```shell
go get -u github.com/spf13/cobra
GOWORK=off cobra-cli init --author sarvsav -l MIT --viper .
cobra-cli add touch
```

## Consequences

The special effects of bubbletea won't be available and the interface will be more traditional.
