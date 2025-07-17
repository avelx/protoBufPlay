Go example of how to work with ProtoBuf

Ref: https://protobuf.dev/getting-started/gotutorial/


Generate source code:
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

====
protoc -I=/Users/pavel/devcore/GoLang/protoBufPlay/protobuf/ --go_out=/Users/pavel/devcore/GoLang/protoBufPlay /Users/pavel/devcore/GoLang/protoBufPlay/protobuf/addressbook.proto