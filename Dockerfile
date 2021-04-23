FROM golang:alpine as build
WORKDIR /code
ENV GOPROXY https://goproxy.cn
COPY myapp .
RUN go build hello.go

FROM alpine
COPY --from=build /code/hello /usr/local/bin/
EXPOSE 8080
CMD ["hello"]
