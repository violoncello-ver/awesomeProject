/*
goの特徴
書いたコードで使ってないのがあるとエラーになる。
コーディングスタイル（読みやすさ重視）：
	・読みやすさのために、演算子などの位置は適宜インデントやスペースを使って、『一番長い部分に合わせて書く。』
	・同じデータ型の変数を1行で宣言すると、コードが長くて分かりにくいこともあるから、『1行で書くべきか調節する。』
*/

/*
package:
goはパッケージ単位のソースコードを管理するため、コードがどのパッケージに属しているか明示する必要がある。
*/
package main

/* import:
標準パッケージ(https://pkg.go.dev/std)
パッケージの詳細を知るコマンド　go doc <packageName> [<funcName>]
*/
import (
	// "fmt"
	"log"
	"os"
)

func main() {
	// log
	// log.Println("logging!")
	// log.Printf("%T %v", "test", "test")
	// FatallnやFatalfを使うとプログラムが強制終了するから、その先実行されない。
	// つまり、logのエラーハンドリングがある場合にプログラムを終了させるために使う。
	// log.Fatalf("%T %v", "test", "test")

	_, err := os.Open("fdafada")
	if err != nil {
		log.Fatalln("Exit\n", err)
	}
	log.Println("logging!")
}
