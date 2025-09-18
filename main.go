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

func main() {

	// pointer
	// new
	// var p *int = new(int)
	// fmt.Println(p)
	// fmt.Println(*p) // 初期値0
	// *p++
	// fmt.Println(*p) // 初期値1

	// var p2 *int
	// fmt.Println(p2)

	// makeとnewの共通点：どちらもメモリ確保してる。
	// ポインタを返すかそうじゃないか。
	// []int
	// map[string]int
	// *int
	// s := make([]int, 0)
	// fmt.Printf("%T\n", s)
	
	// m := make(map[string]int)
	// fmt.Printf("%T\n", m)

	// var p *int = new(int)
	// fmt.Printf("%T\n", p)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)
	
	var st = new(struct{})
	fmt.Printf("%T\n", st)

}
