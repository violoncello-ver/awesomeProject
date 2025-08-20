/* goの特徴
 書いたコード使ってないと怒られる。
*/

/* package:
 goはパッケージ単位のソースコードを管理するため、コードがどのパッケージに属しているか明示する必要がある。
*/
package main

/* import:
 fmt（formatの略）というパッケージをインポートしている。(fmt.Println関数を呼び出したいため。)
 fmt.Println関数の大文字から始まるのは、別のパッケージの関数由来かどうかを一目で分かるようにするため。
 逆に小文字の関数はパッケージ内でしか呼び出せない。
 複数のパッケージを呼び出す時は()を使う。
 標準パッケージ(https://pkg.go.dev/std)
 パッケージの詳細を知るコマンド　go doc <packageName> [<funcName>]
*/ 
import (
	"fmt"
	// "os/user"
	// "time"
)

/* init:
 main関数より優先される特別な関数。
 変数が多いコードで初期設定が必要な時などに使う。
*/
// func init() {
// 	fmt.Println("Init!")
// }

/* Function:
 同じpackage内だと小文字で呼び出せる。
*/
// func bazz() {
// 	fmt.Println("Bazz!")
// }

/* main:
 特別な関数で、コードを実行するとき最初に呼び出される。
 	 1PJ = 1main func
	 必ず初期化する（go mod init <PJName>）
	 複数のgoコードの実行はmain関数を書き換える or 別PJ作成
*/
func main() {
	// // 別packageの関数呼び出し
	// fmt.Println("Hello world", time.Now())
	// fmt.Println(user.Current())
	// // 内部関数呼び出し
	// bazz()

	// 同じ型の変数宣言時に一行にまとめて記述できるけど、読みにくくなる可能性がある。
	// var i int = 1
	// var t, f bool = true, false
	// fmt.Println(i, t, f)

	/* 複数の変数を宣言する方法:
	 それぞれ初期化しない場合のデフォルトの値が存在する（0 0 "" false false）
	 constもできる。
	*/
	// var (
	// 	i 	 int 	 = 1
	// 	f64  float64 = 1.2
	// 	s 	 string  = "test"
	// 	t, f bool 	 = true, false
	// )
	// fmt.Println(i, f64, s, t, f)

	/* 短縮変数宣言(short variable declaration):
	 %Tのデータ型確認（書式指定子）はfmt.Printfを使って表示させる。
	 明示的にデータ型を指定したい場合、varを使う。
	 constは const VarName = valueで定義できる。よしなに、コンパイラがデータ型を判断してくれる。
	*/ 
	xi := 1
	xf64 := 1.2
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xf64)
}