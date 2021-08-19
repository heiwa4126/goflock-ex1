#!/bin/sh
# 死ぬ。または30000にならない
BIN=./goflock-ex1
INIT="$BIN init"
INC="$BIN inc"

$INIT ; $INC & $INC & $INC
