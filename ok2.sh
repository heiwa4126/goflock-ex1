#!/bin/sh
# 30000になる
BIN=./goflock-ex1
INIT="$BIN init2"
INC="$BIN flockinc2"

$INIT ; $INC & $INC & $INC
