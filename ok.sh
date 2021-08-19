#!/bin/sh
# 30000になる
BIN=./goflock-ex1
INIT="$BIN init"
INC="$BIN flock_inc"

$INIT ; $INC & $INC & $INC
