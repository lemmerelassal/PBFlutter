#!/bin/sh

protoc -I $HOME/go/src/ $HOME/go/src/github.com/lemmerelassal/PBFlutter/proto/pbflutter.proto --go_out=plugins=grpc:$HOME/go/src/
mv $HOME/go/src/github.com/lemmerelassal/PBFlutter/pbflutter.pb.go $HOME/go/src/github.com/lemmerelassal/PBFlutter/proto/pbflutter.pb.go

protoc -I $HOME/go/src/ --dart_out=grpc:$HOME/lemmerworld/lemmerworld/lib/proto $HOME/go/src/github.com/lemmerelassal/PBFlutter/proto/pbflutter.proto
mv $HOME/lemmerworld/lemmerworld/lib/proto/github.com/lemmerelassal/PBFlutter/proto/* $HOME/lemmerworld/lemmerworld/lib/proto
sed -i '' 's/..\/..\/..\/..\/google\/protobuf\/any.pb.dart/any.pb.dart/' $HOME/lemmerworld/lemmerworld/lib/proto/pbflutter.pb.dart

#../../../../google/protobuf/


rm -r $HOME/lemmerworld/lemmerworld/lib/proto/github.com/

