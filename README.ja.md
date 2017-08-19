# Gin Scaffold

[![Godoc Reference](https://godoc.org/github.com/dcu/gin-scaffold?status.svg)](http://godoc.org/github.com/dcu/gin-scaffold)
[![Build Status](https://travis-ci.org/dcu/gin-scaffold.svg?branch=master)](https://travis-ci.org/dcu/gin-scaffold)
[![Coverage Status](https://coveralls.io/repos/github/mattn/go-colorable/badge.svg?branch=master)](https://coveralls.io/github/dcu/gin-scaffold?branch=master)
[![Go Report Card](https://goreportcard.com/badge/dcu/gin-scaffold)](https://goreportcard.com/report/dcu/gin-scaffold)

`Gin Scaffold` は CLI `gin` フレームワークのためのスキャフォルディング生成のためのCLI（コマンド・ライン・インターフェース）です。
現在このプロジェクトではそれではデータベースとして `mongodb` と `mgo` のみサポートしています。

## インストール

	go get github.com/dcu/gin-scaffold

## プロジェクトの初期化

	gin-scaffold init <project path>

## モデルの生成

	gin-scaffold model <モデル名> <フィールド名>:<タイプ名>

## コントローラの生成

	gin-scaffold controller <コントローラ名>

## スキャフォルドの生成

	gin-scaffold scaffold <コントローラ名> <フィールド名>:<タイプ名>

## プロジェクトの起動

	gin run <project path>.go

## アクセス

	ブラウザを起動して、http://localhost:4000へアクセス。（デフォルトポート:4000）

