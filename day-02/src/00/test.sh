#!/bin/bash

go build
echo -f
./myFind -f 1

echo -sl
./myFind -sl 1

echo -d
./myFind -d 1
echo -f-ext
./myFind -f -ext "txt" 1
rm myFind