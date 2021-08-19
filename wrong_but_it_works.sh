#!/bin/sh
# 並行に動かさなければ、ちゃんと10000になる。
BIN=./goflock-ex1
INIT="$BIN init"
INC="$BIN inc"

$INIT ; $INC
