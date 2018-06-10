#!/bin/sh

echo 'Previous test results :'
grep '=== RESULT' ./test.log

echo
echo 'Testing ...'
go test > ./test.log
echo

echo 'New test results :'
grep '=== RESULT' ./test.log

echo
echo 'Press any key to close this program.'
read