# Gin Scaffold

[![Godoc Reference](https://godoc.org/github.com/dcu/gin-scaffold?status.svg)](http://godoc.org/github.com/dcu/gin-scaffold)
[![Build Status](https://travis-ci.org/dcu/gin-scaffold.svg?branch=master)](https://travis-ci.org/dcu/gin-scaffold)
[![Coverage Status](https://coveralls.io/repos/github/mattn/go-colorable/badge.svg?branch=master)](https://coveralls.io/github/dcu/gin-scaffold?branch=master)
[![Go Report Card](https://goreportcard.com/badge/dcu/gin-scaffold)](https://goreportcard.com/report/dcu/gin-scaffold)
`Gin Scaffold` is CLI to generate scaffolds for the `gin` framework.
For now the project only supports `mongodb` and `mgo` as database.

## Cài đặt

	go get github.com/dcu/gin-scaffold

## Khởi tạo dự án

	gin-scaffold init <project path>

## Tạo mô hình

	gin-scaffold model <model name> <field name>:<field type>

## Tạo bộ điều khiển

	gin-scaffold controller <controller name>

## Tạo một giàn giáo

	gin-scaffold scaffold <controller name> <field name>:<field type>

## Đang chạy

	go run <project name>.go

## Truy cập

	Open browser, and access to http://localhost:4000. (Default port:4000)

## Cài đặt

```
$ go get github.com/dcu/gin-scaffold
```

# Giấy phép

MIT

# Tác giả

David Cuadrado (dcu)