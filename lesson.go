package main

import "fmt"

func main() {
	// 1PJ = 1main func
	// 必ず初期化する（go mod init <PJName>）
	// 複数のgoコードの実行はmain関数を書き換える or 別PJ作成
	fmt.Println("Hello world")
}