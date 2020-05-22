#!/bin/sh

protoc -I $HOME/go/src/ $HOME/go/src/github.com/lemmerelassal/PBFlutter/proto/pbflutter.proto --go_out=plugins=grpc:$HOME/go/src/
mv $HOME/go/src/github.com/lemmerelassal/PBFlutter/pbflutter.pb.go $HOME/go/src/github.com/lemmerelassal/PBFlutter/proto/pbflutter.pb.go
#protoc -I $HOME/go/src/ --dart_out=grpc:$HOME/lemmerworld/lemmerworld/lib/proto $HOME/go/src/lemr/lemr-backend/lemr-proto/lemr.proto