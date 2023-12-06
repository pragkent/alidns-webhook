#!/bin/sh

go run main.go \
	--tls-cert-file=/tls/tls.crt \
	--tls-private-key-file=/tls/tls.key
