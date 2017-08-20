# Gin Scaffold

[![Godoc Reference](https://godoc.org/github.com/dcu/gin-scaffold?status.svg)](http://godoc.org/github.com/dcu/gin-scaffold)
[![Build Status](https://travis-ci.org/dcu/gin-scaffold.svg?branch=master)](https://travis-ci.org/dcu/gin-scaffold)
[![Coverage Status](https://coveralls.io/repos/github/mattn/go-colorable/badge.svg?branch=master)](https://coveralls.io/github/dcu/gin-scaffold?branch=master)
[![Go Report Card](https://goreportcard.com/badge/dcu/gin-scaffold)](https://goreportcard.com/report/dcu/gin-scaffold)

`Gin Scaffold` 是用于为 `gin` 框架生成脚手架的CLI。
现在该项目只支持 `mongodb` 和 `mgo` 作为数据库。

##安装

	go get github.com/dcu/gin-scaffold


##初始化项目

	gin-scaffold init <project path>

##创建一个模型

	gin-scaffold model<model name> <field name>：<field type>

##创建控制器

	gin-scaffold controller <controller name>

##创建一个脚手架

	gin-scaffold scaffold <controller name> <field name>：<field type>

## Running

	go run <project name>.go

## Accessing

	Open browser, and access to http://localhost:4000. (默认 port:4000)

# 执照

MIT

# 作者

David Cuadrado (dcu)

