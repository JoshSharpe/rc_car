PROTO_DIR="./api/v1"
protoc -I $PROTO_DIR/src $PROTO_DIR/src/motor.proto --go_out=plugins=grpc:$PROTO_DIR
# python -m grpc_tools.protoc -I $PROTO_DIR --python_out=$PROTO_DIR
# python -m grpc_tools.protoc -I $PROTO_DIR/src --python_out=$PROTO_DIR --grpc_python_out=$PROTO_DIR $PROTO_DIR/src/motor.proto
