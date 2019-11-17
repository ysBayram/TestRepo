#!/bin/bash

go test -v ./tests

go test -cover ./tests

# go test -cover ./tests -coverprofile=coverage.out

# go tool cover -html=coverage.out -o coverage.html
