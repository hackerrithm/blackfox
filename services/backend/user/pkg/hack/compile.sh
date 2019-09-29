#!/usr/bin/env sh

protoc -I=../pb --go_out=plugins=grpc:../model ../pb/user.proto