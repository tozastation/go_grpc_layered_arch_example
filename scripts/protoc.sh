SRC="$HOME/go/src"
# Generate User 
protoc -I idl/ --plugin=$HOME/go/bin/protoc-gen-go --go_out=plugins=grpc:$SRC idl/user.proto 
# Generate Weather
protoc -I idl/ --plugin=$HOME/go/bin/protoc-gen-go --go_out=plugins=grpc:$SRC idl/weather.proto 
# Generate Ping 
protoc -I idl/ --plugin=$HOME/go/bin/protoc-gen-go --go_out=plugins=grpc:$SRC idl/ping.proto 