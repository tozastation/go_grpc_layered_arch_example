# INIT Variables
PROJECT_NAME=github.com/tozastation/go_grpc_layered_arch_example
DOMAINS='ping user'

# shellcheck disable=SC2006
LENGTH=`echo "$DOMAINS" | tr ' ' '\n' | wc -l`
# shellcheck disable=SC2006
for i in `seq "$LENGTH"`
do
  DOMAIN=`echo "$DOMAINS" | cut -d ' ' -f "$i"`
  protoc -I "$GOPATH"/src --go_out="$GOPATH"/src "$GOPATH"/src/"$PROJECT_NAME"/internal/grpc/rpc/"$DOMAIN"/generated.proto
done