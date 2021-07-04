FROM golang:1.15-alpine AS build
WORKDIR /go/src
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY . ./

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o potato .

FROM scratch AS runtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/potato/sonolus-uploader-core ./
EXPOSE 8080/tcp
ENTRYPOINT ["./sonolus-uploader-core"]
