# ポインターを学ぶ
ポインターってそもそもなんなのか?

基本的なポインターの使い方:
```go
package main

import "fmt"

func main() {
    var num int = 10

    fmt.Println("numのアドレス:", &num)
    fmt.Println("numの値:", num)

    var ptr *int = &num

    fmt.Println("ptrのアドレス:", &ptr)
    fmt.Println("ptrの値:", ptr)
    fmt.Println(num)

    fmt.Println("numの新しい値:", num)
}
```

&を使うことでこの値のメモリのアドレス、メモリ上のアドレスを取得することができる。

ポインタ型を使うにはこのように書く:
```bash
var ptr *int = &num
```

## ポインターの定義の仕方は２つある
1. 型定義は`*`だけを使うことができる。
2. `&`変数は格納されている値を取ってくることができる.

## Goがメモリ管理をしている方法について
スタックとヒープがある。

基本的に静的、つまりこのコードを見ただけで判別できるものについては`スタック`という領域に入れられる。
コードを見ただけでは、いつまでにどのくらいの大きさを確保すべきなのかどうかわからないというものについては`ヒープ`に入れられる。

スタックかヒープに入れられたのか判別するコマンドがある。
```bash
go build -gcflags '-m -l' main.go
```

実行結果:
```bash
# command-line-arguments
./main.go:6:9: moved to heap: num
./main.go:11:9: moved to heap: ptr
./main.go:8:16: ... argument does not escape
./main.go:8:17: "numのアドレス:" escapes to heap
./main.go:9:16: ... argument does not escape
./main.go:9:17: "numの値:" escapes to heap
./main.go:9:31: num escapes to heap
./main.go:13:16: ... argument does not escape
./main.go:13:17: "ptrのアドレス:" escapes to heap
./main.go:14:16: ... argument does not escape
./main.go:14:17: "ptrの値:" escapes to heap
./main.go:15:16: ... argument does not escape
./main.go:15:17: num escapes to heap
./main.go:17:16: ... argument does not escape
./main.go:17:17: "numの新しい値:" escapes to heap
./main.go:17:40: num escapes to heap
```

## ガベージコレクター
Goだとプログラムが動的に確保したメモリ領域、スタックとヒープのどちらかというと、ヒープの方。
そのデータの中身が必要なくなったとき、そのメモリ領域を指すポインタがないときに自動的にそのメモリを解放するといったことをやってくれる。

Goのガベージコレクターではマークアンドスイープ方式というものが採用されている。このアルゴリズムは２つのマークとスウィープから成り立っている。

### マークフェーズについて
マークフェーズでは、すべてのオブジェクトをマークとして置いて、その後ROOTとなるオブジェクトから開始して、到達可能なすべてのオブジェクをマークしていく。

ここでのルートオブジェクトとは、例えばグローバル変数や現在実行中の関数のローカル変数、そしてCPUレジスタなどを指す。

### スイープフェーズについて:
スイープというのは掃除するといったような意味。マークされていない全てのオブジェクト、つまりガベージを解放する。

これによって不要なメモリが解放されて、再利用可能な形になる。