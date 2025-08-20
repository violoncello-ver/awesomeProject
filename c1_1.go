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
	"os/user"
	"time"
)

/* init:
 main関数より優先される特別な関数。
 変数が多いコードで初期設定が必要な時などに使う。
*/
func init() {
	fmt.Println("Init!")
}

/* Function:
 同じpackage内だと小文字で呼び出せる。
*/
func bazz() {
	fmt.Println("Bazz!")
}

/* main:
 特別な関数で、コードを実行するとき最初に呼び出される。
 	 1PJ = 1main func
	 必ず初期化する（go mod init <PJName>）
	 複数のgoコードの実行はmain関数を書き換える or 別PJ作成
*/
func main() {
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())

	// 内部関数呼び出し
	bazz()

}