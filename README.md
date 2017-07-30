# Gin Scaffold

`Gin Scaffold` is CLI to generate scaffolds for the `gin` framework.
For now the project only supports `mongodb` and `mgo` as database.

## Installation

	go get github.com/dcu/gin-scaffold

## Initializing a project

	git-scaffold init <project path>

## Creating a model

	git-scaffold model <model name> <field name>:<field type>

## Creating a controller

	gin-scaffold controller <controller name>

## Creating a scaffold

	gin-scaffold scaffold <controller name> <field name>:<field type>

