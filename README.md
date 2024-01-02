# Golang に入門する

Hello World をする

```go
package main

import (
    "fmt"
)

func main() {
	fmt.Println("Hello World!")
  }
```

## ブレークポイントを使う

`launch.json`を設定して、`interface`を格納した変数を用意する。

[こちらの記事を参照](https://zenn.dev/jboy_blog/articles/9799a68a79c829)

```go
package main

import (
    "fmt"
)

func greeting() {
	  fmt.Println("こんにちは☀️")
}

func main() {
	// if文でnilだったらエラーを返しnilじゃなかったら挨拶を返すfunctionを実行する
	var val interface{} = 1
	if val == nil {
	  fmt.Println("エラーです")
	} else {
	  greeting()
	}
}
```

## for文を使ってみる
for文で、数字を増やして、breakする処理を書いてみる。10までfor文でloopするが、6になったら分岐処理で抜ける。

```go
package main

import (
    "fmt"
)

func main() {
	// forで値を１０まで増やして6でbreakする処理を書いてlogを出す
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			fmt.Println("breakした")
			break
		}
  }
}
```

実行結果:
```bash
hashimotojunichi@hashimotojunichinoMacBook-Pro HelloWorld % go run main.go
0
1
2
3
4
5
6
breakした
```

for文で数字を10まで増やして、`continue`で偶数の数値だけlogに出す処理。

```go
package main

import (
    "fmt"
)

func main() {
	// forで値を１０まで数値を増やして、偶数でcontinueを使ってスキップしてlogを出力する
	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			// 出力されているのは偶数のみ
			fmt.Println("continueして偶数を表示:", i)
			continue
		}
		fmt.Println(i)
	}
}
```

実行結果:
```bash
hashimotojunichi@hashimotojunichinoMacBook-Pro HelloWorld % go run main.go
continueして偶数を表示: 0
1
continueして偶数を表示: 2
3
continueして偶数を表示: 4
5
continueして偶数を表示: 6
7
continueして偶数を表示: 8
9
continueして偶数を表示: 10
```

## switchを使ってみる
switch文で分岐処理を書いてみる。複数の値をcaseで比較することもできるみたいです。

```go
package main

import "fmt"

func main() {
    x := "apple"
    switch x {
    case "apple", "banana", "peach":
        fmt.Println("x is a fruit")
    case "carrot", "celery", "beet":
        fmt.Println("x is a vegetable")
    default:
        fmt.Println("x is not a fruit or a vegetable")
    }
}
```

実行結果:
```bash
x is a fruit
```