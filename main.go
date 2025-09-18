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
	"fmt"
	"os"
)

// func getOsName() string {
// 	return "xxx"
// }

// func foo() {
// 	// defer延期する
// 	defer fmt.Println("defer foo")
// 	fmt.Println("hello foo")
// }

func main() {

	// switch os := getOsName(); os {
	// case "mac":
	// 	fmt.Println("mac!!")
	// case "windows":
	// 	fmt.Println("Windows!!")
	// default:
	// 	fmt.Println("Default!!", os)
	// }

	// fmt.Println("main world")
	// foo()
	// defer fmt.Println("main defer 1")
	// defer fmt.Println("main defer 2")
	// defer fmt.Println("main defer 3")
	// fmt.Println("main hello")

	// deferの動き
	// 関数、mainそれぞれにdeferがある場合。
	// 関数のdefer -> mainのdefer（複数あったら、古い(上から定義された)ものは最後に来る。）

	// deferどこで使うん？ -> ファイルの閉じ忘れとか
	// file, _ := os.Open("./main.go")
	// defer file.Close()
	// data := make([]byte, 100)
	// file.Read(data)
	// fmt.Println(data)
	// fmt.Println(string(data))
}
