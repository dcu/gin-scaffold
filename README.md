# Gin Scaffold

[![Godoc Reference](https://godoc.org/github.com/dcu/gin-scaffold?status.svg)](http://godoc.org/github.com/dcu/gin-scaffold)
[![Build Status](https://travis-ci.org/dcu/gin-scaffold.svg?branch=master)](https://travis-ci.org/dcu/gin-scaffold)
[![Coverage Status](https://coveralls.io/repos/github/mattn/go-colorable/badge.svg?branch=master)](https://coveralls.io/github/dcu/gin-scaffold?branch=master)
[![Go Report Card](https://goreportcard.com/badge/dcu/gin-scaffold)](https://goreportcard.com/report/dcu/gin-scaffold)

`Gin Scaffold` is CLI to generate scaffolds for the `gin` framework.
For now the project only supports `mongodb` and `mgo` as database.

## Installation

	go get github.com/dcu/gin-scaffold

## Initializing a project

	gin-scaffold init <project path>

## Creating a model

	gin-scaffold model <model name> <field name>:<field type>

## Creating a controller

	gin-scaffold controller <controller name>

## Creating a scaffold

	gin-scaffold scaffold <controller name> <field name>:<field type>

## Running

	go run <project name>.go

## Accessing

	Open browser, and access to http://localhost:4000. (Default port:4000)

## Installation

```
$ go get github.com/dcu/gin-scaffold
```

# License

MIT

# Author

David Cuadrado (dcu)

