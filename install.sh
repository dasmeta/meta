#!/bin/bash

go mod init meta
go build ./meta.go
go list -f "{{.Target}}"
PATH=/bin:/usr/bin:/usr/local/bin:/`go list -f "{{.Target}}"`:${PATH}
export PATH=$PATH:/`go list -f "{{.Target}}"`
go env -w GOBIN=/`go list -f "{{.Target}}"`
go install
alias meta=`go list -f "{{.Target}}"`/meta