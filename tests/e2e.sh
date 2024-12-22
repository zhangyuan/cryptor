#! /bin/bash

set -euo pipefail

mkdir -p tmp/tests

go run main.go e -p a -s b -m gcm -i tests/fixtures/data.txt -o tmp/tests/encrypted.gcm.txt
go run main.go d -p a -s b -m gcm -i tmp/tests/encrypted.gcm.txt -o tmp/tests/decrypted.gcm.txt
cmp tests/fixtures/data.txt tmp/tests/decrypted.gcm.txt

go run main.go e -p a -s b -m cfb -i tests/fixtures/data.txt -o tmp/tests/encrypted.cfb.txt
go run main.go d -p a -s b -m cfb -i tmp/tests/encrypted.cfb.txt -o tmp/tests/decrypted.cfb.txt
cmp tests/fixtures/data.txt tmp/tests/decrypted.gcm.txt