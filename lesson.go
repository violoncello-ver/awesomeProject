// goはパッケージ単位のソースコードを管理するため、コードがどのパッケージに属しているか明示する必要がある。
package main

/*	
 fmt（formatの略）というパッケージをインポートしている。(fmt.Println関数を呼び出したいため。)
 fmt.Println関数の大文字から始まるのは、別のパッケージの関数由来かどうかを一目で分かるようにするため。
 逆に小文字の関数はパッケージ内でしか呼び出せない。
*/ 
import "fmt"

// main関数より優先される特別な関数
// 変数が多いコードで初期設定が必要な時などに使う。
func init() {
	fmt.Println("Init!")
}

// 別関数
func bazz() {
	fmt.Println("Bazz!")
}

// 特別な関数で、コードを実行するとき最初に呼び出される。
func main() {
	/*
	 1PJ = 1main func
	 必ず初期化する（go mod init <PJName>）
	 複数のgoコードの実行はmain関数を書き換える or 別PJ作成
	*/
	fmt.Println("Hello world", "golang")

	// 内部関数呼び出し
	// bazz()
	
}