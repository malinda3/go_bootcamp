#!/bin/bash

go build
echo "executing: "
echo "$ echo "-la" | ./myXargs ls"
echo "-la" | ./myXargs ls

rm myXargs