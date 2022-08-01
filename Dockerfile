FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
//Need close CGO, docker run error: standard_init_linux.go:228: exec user process caused: no such file or directory
ENV CGO_ENABLED=0 
#ENV GO111MODULE=on 
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8080
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
