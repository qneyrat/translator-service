```
> $ # install protoc https://github.com/google/protobuf/releases
> $ go get -u github.com/golang/protobuf/protoc-gen-go
> $ go get -u github.com/uber/prototool/cmd/prototool
> $ prototool init
> $ # edit prototool.yaml

> $ cat proto/input.json | prototool grpc proto/translator.proto 0.0.0.0:4000 proto.Translator/Translate -
```
