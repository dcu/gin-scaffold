# Gin Scaffold

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

