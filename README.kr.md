# Gin Scaffold

`Gin Scaffold`는 CLI `gin`체제를 위한 스캐포르딘그 생성을 위한 CLI(커멘드 라인 인터페이스)입니다.
현재 이 프로젝트에서는 그러면 데이타베이스로서 `mongodb` 와 `mgo` 마셔 서포트하고 있습니다.

## 인스톨

	go get github.com/dcu/gin-scaffold

## 프로젝트의 초기화

	gin-scaffold init 

## 모델의 생성

	gin-scaffold model <모델명> <필드명>:<타입명>

## 콘트롤러의 생성

	gin-scaffold controller <콘트롤러명>

## 스캐포르드의 생성

	gin-scaffold scaffold <콘트롤러명> <필드명>:<타입명>

## 액세스

	브라우저를 기동하고, http://localhost:4000 에 액세스.(디폴트 포토:4000)
