#!/bin/bash
rm -rf demo.db
go run main.go
sqlite3 demo.db ".schema"
sqlite3 demo.db "select * from user"