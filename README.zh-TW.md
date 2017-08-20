# Gin Scaffold

[![Godoc Reference](https://godoc.org/github.com/dcu/gin-scaffold?status.svg)](http://godoc.org/github.com/dcu/gin-scaffold)
[![Build Status](https://travis-ci.org/dcu/gin-scaffold.svg?branch=master)](https://travis-ci.org/dcu/gin-scaffold)
[![Coverage Status](https://coveralls.io/repos/github/mattn/go-colorable/badge.svg?branch=master)](https://coveralls.io/github/dcu/gin-scaffold?branch=master)
[![Go Report Card](https://goreportcard.com/badge/dcu/gin-scaffold)](https://goreportcard.com/report/dcu/gin-scaffold)

For now the project only supports `mongodb` and `mgo` as database.
`Gin Scaffold` 是用於為 `gin` 框架生成腳手架的CLI。
現在該項目只支持 `mongodb` 和 `mgo` 作為數據庫。

## 安裝 

	go get github.com/dcu/gin-scaffold

## 創建一個模型

	gin-scaffold init <project path>

## 創建控制器

	gin-scaffold model <model name> <field name>:<field type>

## 創建一個腳手架

	gin-scaffold controller <controller name>

## 創建一個腳手架

	gin-scaffold scaffold <controller name> <field name>:<field type>

## 賽跑

	go run <project name>.go

## 訪問

	Open browser, and access to http://localhost:4000. (Default port:4000)

# 執照

MIT

# 作者

David Cuadrado (dcu)

