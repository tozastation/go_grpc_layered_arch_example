# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/tozastation/gRPC-Training-Golang
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/tozastation/gRPC-Training-Golang/app /app
ENV CONNECTION_STRING=Data Source=127.0.0.1:1433;Initial Catalog=Weather;User ID=SA;Password=Test@1234
ENV VENDER=mssql
ENV OPENWEATHER_URL=http://api.openweathermap.org/data/2.5/weather
ENV OPENWEATHER_CREDENTIAL=&APPID=1e16e8941ce99bdd844d129d5179d98a
EXPOSE 3001
ENTRYPOINT ["/app"]