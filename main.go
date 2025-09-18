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
)

func one(x *int) {
	*x = 1
}

func main() {

	// pointer
	// &は値が入っているアドレスを出す
	// *はアドレス先の値を出す
	var n int = 100
	one(&n) // pointer
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*&n)
	fmt.Println(&*&n)
	
	
	// var n int = 100
	// fmt.Println(n)
	// fmt.Println(&n)
	
	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

}
