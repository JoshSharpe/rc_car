# protoc  -I "./api/v1/src" --go_out=plugins=grpc:"./api/v1" ./api/v1/src/motor.proto
# python -m grpc_tools.protoc -I $PROTO_DIR --python_out=$PROTO_DIR
# python -m grpc_tools.protoc -I $PROTO_DIR/src --python_out=$PROTO_DIR --grpc_python_out=$PROTO_DIR $PROTO_DIR/src/motor.proto

protoc \
    -I ./api/v1/src \
    --go_out=plugins=grpc:./api/v1 \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./src/client/api/v1/ \
    --js_out=import_style=commonjs,binary:./src/client/api/v1/ \
    ./api/v1/src/motor.proto