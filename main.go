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
	// "os/user"
	// "time"
	// "strings"
	// "strconv"
)

/* init:
main関数より優先される特別な関数。
変数が多いコードで初期設定が必要な時などに使う。
*/
// func init() {
// 	fmt.Println("Init!")
// }

// func bazz() {
// 	fmt.Println("Bazz!")
// }

// var (
// 		i 	 int 	 = 1
// 		f64  float64 = 1.2
// 		s 	 string  = "test"
// 		t, f bool 	 = true, false
// )

// func foo() {
// 	xi := 1
// 	fmt.Println(xi)
// 	xi = 2
// 	fmt.Println(xi)
// } 

/* main:
特別な関数で、コードを実行するとき最初に呼び出される。
1PJ = 1main func
必ず初期化する（go mod init <PJName>）
複数のgoコードの実行はmain関数を書き換える or 別PJ作成
*/

// constとvarのoverflowの違い:
// var big_var int = 9223372036854775807 + 1
const big_const = 9223372036854775807 + 1

func main() {

	// constとvarのoverflowの違い:
	// fmt.Println(big_var - 1)
	fmt.Println(big_const - 1)

	/* 複数の変数を宣言する方法:
	それぞれ初期化しない場合のデフォルトの値が存在する（0 0 "" false false）
	constもできる。
	*/

	//  変数宣言方法毎の値の変更の挙動:
	// fmt.Println(i, f64, s, t, f)
	// i = 5
	// fmt.Println(i, f64, s, t, f)
	// foo()

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
	論理演算子は&& || ! を使い、他の言語とそこまで違いがなさそう。
	*/
	// t, f := true, false
	// fmt.Printf("%T %v\n", t, t)
	// fmt.Printf("%T %v\n", f, f)

	/* cast:
	データ型のcast:
		型変換はtype(変換したい変数)を指定する。
		※下記の1をキャストして%vの時に%fと同じ出力にならないのは、よしなに1.0よりも１として出力する方が自然であるため。
	文字列のcast:
		strconv.Atoi -> string conversion(文字列変換).A(ASCII) to i(Integer)
	*/
	// var x int = 1
	// xx := float64(x)
	// fmt.Printf("%T %v %f\n", xx, xx, xx)
	// var s string = "8"
	// i, _ := strconv.Atoi(s)  // 変数-> 1.変換結果の整数値 (`int`),　2.エラー情報 (`error`) -> 「整数値だけ受け取り、エラーは無視」必ず2つ返すが、2.に_は推奨されていない。
	// fmt.Printf("%T %v", i, i)
	// 推奨される書き方↓
	// var s string = "8"
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println("変換エラー:", err)
	// 	return
	// }
	// fmt.Printf("%T %v", i, i)

	/* goのArrayデータ構造
	Array(配列):
		固定長（後から長さを変えられない）
		宣言のみ:
			a[2]int{value1, value2}
			arrayName[elements]dataType
		初期化:
			arrayName := [elements]{value1, value2...}
	slice(スライス):
		可変長（後から長さを変えられる）
		sliceName := []int{1, 2, 3, 4, 5}から特定の範囲の呼び出し方:
			[欲しい何番目かの値 = 欲しいIndex = 番号 - 1]
			範囲は、その「欲しい何番目かの値」以上欲しいIndex値+1でとってくる。
		要素の追加:
			append(sliceName, value1, value2...)
		多次元スライス:
			var sliceName = [][]int{
			[]int{value1, value2...},
			[]int{value1, value2...},
			...
				}
		make関数:
			要素0で初期化されたsliceを作れる。
			n := make([]int, lengthNum, capacityNum)
			fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
			len()は現状の要素数を表す。
			cap()は事前に分かっているメモリ(要領)を確保することを表す。appendで要領範囲の要素を追加すると高速で処理できる。
			capacityNumを省略すると cap=len になる
	Map(マップ):
		型が異なる複数の値をKey, Valueの組み合わせで管理する。

	共通点:
		Array, sliceは値をインデックス指定して変更できる。
		ByteArray:
			Array, Sliceで使える。
			引数をByteで値を受け取り、ASCIIコードとして扱えるのでstring(ByteArray)で文字列を出力できる。（文字列 ↔︎ Byte(ASCIIコード)で扱える）
			Byte配列にbyte or 文字列を入れる→文字列出力:
				※書き方注意 []byte{Value1, Value2...} or []byte("文字列")
				Byte配列にどちらかいれても、出力はValueになる。
				string(value)すれば文字列が出力する。
			用途:
				ネットワークやファイル操作などで利用できる。

	固定長 or 可変長の見分け方:
		初期化時に[elements]か[]しているか。
	*/
	// var a [2]int
	// a[0] = 100
	// a[1] = 200
	// fmt.Println(a)
	// a := [2]int{100, 200}
	// fmt.Println(a)
	// 初期化:
	// var board = [2]int{1, 2}
	// fmt.Println(board)

	// // スライスの呼び出し方:
	// n := []int{1, 2, 3, 4, 5}
	// fmt.Println(n[2])
	// fmt.Println(n[1:4])  // 2 3 4
	// fmt.Println(n[:])    // 1 2 3 4 5
	// fmt.Println(n[:4])   // 1 2 3 4
	// fmt.Println(n[3:5])  // 4 5

	// n := []int{1, 2, 3, 4, 5}
	// fmt.Println(n)
	// n = append(n, 100)
	// fmt.Println(n)
	// n = append(n, 200, 300, 400, 500)
	// fmt.Println(n)

	// var board = [][]int{
	// 	[]int{0, 1, 2},
	// 	[]int{3, 4, 5},
	// 	[]int{6, 7, 8},
	// }
	// fmt.Println(board)
	// fmt.Println(board[1])
	// fmt.Println(board[1][2])

	// n := make([]int, 3, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 0, 0)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 1, 2, 3, 4, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// a := make([]int, 3)
	// fmt.Printf("len=%d cap=%d value=%v", len(a), cap(a), a)
	
	// a := []byte{72, 73}
	// fmt.Println(a)
	// fmt.Printf(string(a))

	var a, b string = "Hello", "Go"

	fmt.Println(a, b)

}
