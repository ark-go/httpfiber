SHELL := /bin/bash
arklibgo := ~/ProjectsGo/arkAlias.sh
version = ~/ProjectsGo/arkAlias.sh getlastversion
.PHONY: 

.SILENT: run build buildwin


build:
	$(info +Компиляция Linux)
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o bin/main/httpfiber cmd/main/main.go

buildwin:
	$(info +Компиляция Windows)
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o ./bin/main/httpfiber.exe -tags static -ldflags "-s -w -X 'main.versionProg=$$($(version))'" cmd/main/main.go

run: build buildwin
	$(info +Запуск)
	bin/main/httpfiber