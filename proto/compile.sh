#!/usr/bin/env bash

protoc -I=. --go_out=./ ./service.proto ./errors.proto