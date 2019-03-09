# Implement Mock
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping CheckClient > implements/mock/ping_implement_mock.go 
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/weather WeathersClient > implements/mock/weather_implement_mock.go
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user UsersClient > implements/mock/user_implement_mock.go
# Service Mock
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/domain/service IUserService  > domain/service/mock/user_service_mock.go
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/domain/service IWeatherService  > domain/service/mock/weather_service_mock.go
# Repository Mock
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/domain/repository IUserRepository  > infrastructure/persistence/mssql/mock/user_repository_mock.go
$GOPATH/bin/mockgen -package=mock github.com/tozastation/gRPC-Training-Golang/domain/repository IWeatherRepository  > infrastructure/persistence/mssql/mock/weather_repository_mock.go

