FROM golang:1.11.0-alpine3.8 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/hackerrithm/blackfox/services/backend/reaction
COPY vendor ../vendor
COPY services/backend/profile ../profile
COPY services/backend/user ../user
COPY ./services/backend/reaction ./
RUN go build -o /go/bin/app ./cmd/reaction/main.go

FROM alpine:3.8
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
