#!/bin/sh
# 30000になる
BIN=./goflock-ex1
INIT="$BIN init3"
INC="$BIN flockinc3"

$INIT ; $INC & $INC & $INC
