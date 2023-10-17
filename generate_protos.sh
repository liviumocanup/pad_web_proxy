#!/bin/bash

# Base directory for proto files
PROTO_DIR="proto"

# Generating Go code for common proto in playback_service
protoc --go_out=playback_service/proto/common --go_opt=paths=source_relative \
       --go-grpc_out=playback_service/proto/common --go-grpc_opt=paths=source_relative \
       -I${PROTO_DIR} \
       ${PROTO_DIR}/common.proto

# Generating Go code for common proto in track_service
protoc --go_out=track_service/proto/common --go_opt=paths=source_relative \
       --go-grpc_out=track_service/proto/common --go-grpc_opt=paths=source_relative \
       -I${PROTO_DIR} \
       ${PROTO_DIR}/common.proto

# Generating Go code for playback_service
protoc --go_out=playback_service/proto --go_opt=paths=source_relative \
       --go-grpc_out=playback_service/proto --go-grpc_opt=paths=source_relative \
       -I${PROTO_DIR} \
       ${PROTO_DIR}/playback.proto ${PROTO_DIR}/track.proto

# Generating Go code for track_service
protoc --go_out=track_service/proto --go_opt=paths=source_relative \
       --go-grpc_out=track_service/proto --go-grpc_opt=paths=source_relative \
       -I${PROTO_DIR} \
       ${PROTO_DIR}/track.proto

# Generating Go code for user_service
protoc --go_out=user_service/proto --go_opt=paths=source_relative \
       --go-grpc_out=user_service/proto --go-grpc_opt=paths=source_relative \
       -I${PROTO_DIR} ${PROTO_DIR}/user.proto

# Generating Python code for gateway
python3 -m grpc_tools.protoc \
    --python_out=gateway/proto \
    --grpc_python_out=gateway/proto \
    -I${PROTO_DIR} \
    ${PROTO_DIR}/playback.proto ${PROTO_DIR}/track.proto ${PROTO_DIR}/user.proto ${PROTO_DIR}/common.proto

# Modify imports in playback.pb.go
sed -i 's/_ \".\"/common \"playback_service\/proto\/common\"/g' playback_service/proto/playback.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' playback_service/proto/playback.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' playback_service/proto/playback.pb.go

# Modify imports in playback_grpc.pb.go
sed -i 's/_ \".\"/common \"playback_service\/proto\/common\"/g' playback_service/proto/playback_grpc.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' playback_service/proto/playback_grpc.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' playback_service/proto/playback_grpc.pb.go

# Modify imports in track.pb.go
sed -i 's/_ \".\"/common \"track_service\/proto\/common\"/g' track_service/proto/track.pb.go
sed -i 's/_ \".\"/common \"playback_service\/proto\/common\"/g' playback_service/proto/track.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' track_service/proto/track.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' playback_service/proto/track.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' track_service/proto/track.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' playback_service/proto/track.pb.go

# Modify imports in track_grpc.pb.go
sed -i 's/_ \".\"/common \"track_service\/proto\/common\"/g' track_service/proto/track_grpc.pb.go
sed -i 's/_ \".\"/common \"playback_service\/proto\/common\"/g' playback_service/proto/track_grpc.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' track_service/proto/track_grpc.pb.go
sed -i 's/_.TrackMetadata/common.TrackMetadata/g' playback_service/proto/track_grpc.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' track_service/proto/track_grpc.pb.go
sed -i 's/_.StatusResponse/common.StatusResponse/g' playback_service/proto/track_grpc.pb.go

# Update imports in Python generated files
# Modify imports in playback_pb2.py
sed -i 's/import common_pb2 as common__pb2/from proto import common_pb2 as common__pb2/g' gateway/proto/playback_pb2.py

# Modify imports in playback_pb2_grpc.py
sed -i 's/import playback_pb2 as playback__pb2/from proto import playback_pb2 as playback__pb2/g' gateway/proto/playback_pb2_grpc.py
sed -i 's/import common_pb2 as common__pb2/from proto import common_pb2 as common__pb2/g' gateway/proto/playback_pb2_grpc.py

# Modify imports in track_pb2.py
sed -i 's/import common_pb2 as common__pb2/from proto import common_pb2 as common__pb2/g' gateway/proto/track_pb2.py

# Modify imports in track_pb2_grpc.py
sed -i 's/import track_pb2 as track__pb2/from proto import track_pb2 as track__pb2/g' gateway/proto/track_pb2_grpc.py
sed -i 's/import common_pb2 as common__pb2/from proto import common_pb2 as common__pb2/g' gateway/proto/track_pb2_grpc.py

# Modify imports in user_pb2_grpc.py
sed -i 's/import user_pb2 as user__pb2/from proto import user_pb2 as user__pb2/g' gateway/proto/user_pb2_grpc.py


echo "Proto files generated successfully."