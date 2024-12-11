#!/bin/bash

go build
echo -l
./myWc -l go.mod
echo -w
./myWc -w go.mod
echo -m
./myWc -m go.mod
echo "noflag (-w)"
./myWc go.mod
echo "2 files"
./mywc go.mod main.go
rm myWc