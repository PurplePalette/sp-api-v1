FROM golang:1.15 AS build
WORKDIR /go/src
COPY . ./

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o potato .

FROM scratch AS runtime
COPY --from=build /go/src/potato/sonolus-uploader-core ./
EXPOSE 8080/tcp
ENTRYPOINT ["./sonolus-uploader-core"]
