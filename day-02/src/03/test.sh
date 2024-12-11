#!/bin/bash

touch test_log_dir/test1.log
echo "123123123123\njesnkjvnkdfjnvdn\nfjsdkfndskjf|1231231231231\nflsjnfkjsdnfjk\n" >> test_log_dir/test1.log

touch test_log_dir/test2.log
echo "123123123123\njesnkjvnkdfjnvdn\nfjsdkfndskjf|1231231231231\nflsjnfkjsdnfjk\n" >> test_log_dir/test2.log

go build
echo "running with no flags and save to the log folder"
./myRotate test_log_dir/test1.log
echo "running with -a and save to a specific folder (./)"
./myRotate -a ./ test_log_dir/test2.log
rm myRotate