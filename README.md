# goflock-ex1

GoLang で複数プロセスで flock を使って排他制御を行うサンプルコード。
flock()を使っているので UNIX のみ動く。詳しくは`man 2 flock`参照。

[heiwa4126/python_flock_ex1: Python 3で複数プロセスでflockを使って排他制御を行うサンプルコード](https://github.com/heiwa4126/python_flock_ex1)
がちょっとおもしろかったので、Go 版を書いてみました。

# 動かし方

まずビルド

```sh
go build
```

Go なのでサブコマンド式になってます。

```sh
BIN=./goflock-ex1
INIT="$BIN init"      # カウンターを0に初期化
INC="$BIN inc"        # カウンターを10000増やす(排他制御なし)。並行で複数動かすと死ぬ。
FINC="$BIN flockinc"  # カウンターを10000増やす。並行で動かしても死なない。

# 期待通りに動いて、10000になる
$INIT ; $INC

# 死ぬ。または30000にならない
$INIT ; $INC & $INC & $INC

# 30000になる
$INIT ; $FINC & $FINC & $FINC
```

# ex1,2 ...

- ex1 - 素朴版。もとの Python と同じ。ただしバッファリングなし。ファイルは数字文字列
- ex2 - バッファリングなし。ファイルは UINT64 をリトルエンディアンで
- ex3 - 無理やりバッファ I/O にしてみた。ファイルは ex2 と同じ
- ex4 - ex2 を [gofrs/flock](https://github.com/gofrs/flock) にしてみたもの。ロックファイルがカウンタファイルと別

# そのほかメモ

Python とちがって
カウンターファイルを IO バッファリングしてないので、
ファイルそのものにロックがかけられる。

まあ効率を考えるとファイルの規模が大きくなるとバッファ使わないわけにいかないので
Python 同様、別ロックファイルにするべきだろう。
-> それはどうだろうか。

あと、Python 版と等しくカウンターファイルを文字列にしてるんだけど
効率よくないので UINT64 そのもの書く版も学習用に作ってみる。
-> ex2 として作ってみました。

Rust 版も作ってみるかなあ。

UNIX 以外もサポートする場合は
[gofrs/flock: Thread\-safe file locking library in Go \(originally github\.com/theckman/go\-flock\)](https://github.com/gofrs/flock)
を参照。

終わったら

```sh
rm /tmp/goflock-ex1-count*
```

すること。
