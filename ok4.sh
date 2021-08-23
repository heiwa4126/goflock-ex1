#!/bin/sh
# 30000になる
BIN=./goflock-ex1
INIT="$BIN init4"
INC="$BIN flockinc4"

$INIT ; $INC & $INC & $INC
