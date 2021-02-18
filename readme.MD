# golang tools

## go run 
* go run main.go

## go env

## go mod
* go mod init number-game 

## go test
* go test ./game
* go test ./... 

## go build
* GOOS=linux GOARCH=386 go build
* GOOS=windows GOARCH=amd64 go build -o ./number-guess

## 參考資料
* https://golang.org/doc/install/source

## BDD behavior-driven development
* Gherkin語法：https://cucumber.io/docs/gherkin/reference/
* cucumber支援各種主流語言 Ruby、JAVA、javascript、.NET、PHP
* cucumber in golang: https://github.com/cucumber/godog
* .feature檔案方便閱讀、傳遞知識