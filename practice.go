/* goの特徴
書いたコード使ってないと怒られる。
コーディングスタイル：読みやすさのために、演算子などの位置は一番長い部分に合わせて書く。
*/

/*
package:
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
	// "strings"
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
	 %Tのデータ型確認や%v値の確認（書式指定子）等はfmt.Printfを使って表示させる。
	 明示的にデータ型を指定したい場合、varを使う。
	 constは const VarName = valueで定義できる。よしなに、コンパイラがデータ型を判断してくれる。
	*/
	// xi := 1
	// xf64 := 1.2
	// fmt.Printf("%T\n", xi)
	// fmt.Printf("%T\n", xf64)

	/* 数値型の基本:
	 signed Integer（符号付き整数）, Unsigned Integer（符号なし整数）と、それぞれちゃんと明確に区別している。
	 例）uint8は2^8=256となり、0~255の範囲だから256にするとoverflowになる。
	*/
	// var (
	// 	u8  uint8     = 255
	// 	i8  int8      = 127
	// 	f32 float32   = 0.2
	// 	c64 complex64 = -5 + 12i
	// )
	// fmt.Println(u8, i8, f32, c64)
	// fmt.Printf("type=%T value=%v\n", c64, c64)
	// // インクリメント++、デクリメント--
	// i8++
	// fmt.Println(i8)
	// i8--
	// fmt.Println(i8)
	// i8++
	// fmt.Println(i8)
	// // シフト演算(倍に変化してる)
	// fmt.Println(1 << 0)	 // 0001<-0000 (0番目から左に1シフト) 2^0
	// fmt.Println(1 << 1)	 // 0010<-0001 (1番目から左に1シフト) 2^1
	// fmt.Println(1 << 2)  // 0100<-0010 (2番目から左に1シフト) 2^2
	// fmt.Println(1 << 3)  // 1000<-0100 (3番目から左に1シフト) 2^3
	
	/* string:
		特定の文字を出力:
		 Goの場合、特定の文字列を出力しよう "Name"[0] とするとASCII Codeで出力される。
		 文字として取り出したい時は string("Name"[0]) のように囲む必要がある。
		文字列内の特定の文字の置き換え:
		 import strings して、文字列操作の関数一覧からstrings.Replace関数を使う。
		一見すると別扱いの文字列:
		 文字列リテラル："", ''で囲まれている文字列。配列と違って、何かしら文字を自由に代入することができない。他の言語と若干ルール固いかも？
		改行:
		 \n または ''を使って改行
		意味を持っている記号を文字として出力するには:
		 \ を使うか、または``で囲う
	*/
	// fmt.Println("Go practice"[0])
	// fmt.Println(string("Go practice"[0]))
	// var s string = "Go Go practice"
	// fmt.Println(s)
	// s = strings.Replace(s, "G", "X", 1)  // （どの文字列、変換元、変換後、何回分？）
	// fmt.Println(s)
	// var s string = "Go practice"
	// fmt.Println(s)
	// fmt.Println(strings.Contains(s, "practice"))
	
	/* bool:
		
	*/
	t, f := true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)

}
