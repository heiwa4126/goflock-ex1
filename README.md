# goflock-ex1

GoLangで複数プロセスでflockを使って排他制御を行うサンプルコード。
flock()を使っているのでUNIXのみ動く。詳しくは`man 2 flock`参照。

[heiwa4126/python_flock_ex1: Python 3で複数プロセスでflockを使って排他制御を行うサンプルコード](https://github.com/heiwa4126/python_flock_ex1)
がちょっとおもしろかったので、Go版を書いてみました。


# 動かし方

まずビルド
```sh
go build
```

Goなのでサブコマンド式になってます。

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


# そのほかメモ

Pythonとちがって
カウンターファイルをIOバッファリングしてないので、
ファイルそのものにロックがかけられる。

まあ効率を考えるとファイルの規模が大きくなるとバッファ使わないわけにいかないので
Python同様、別ロックファイルにするべきだろう。

あと、Python版と等しくカウンターファイルを文字列にしてるんだけど
効率よくないのでUINT64そのもの書く版も学習用に作ってみる。

Rust版も作ってみるかなあ。
