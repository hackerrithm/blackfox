FROM golang:1.11.2-alpine3.8 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/hackerrithm/blackfox/services/backend/api/pkg
COPY vendor ../../vendor
COPY services/backend/user ../../user
COPY services/backend/space ../../space
COPY services/backend/profile ../../profile
COPY services/backend/task ../../task
COPY services/backend/post ../../post
COPY services/backend/geography ../../geography
COPY services/backend/goal ../../goal
COPY services/backend/match ../../match
COPY services/backend/group ../../group


COPY ./services/backend/api/pkg ./
RUN go build -o /go/bin/app ./server/

FROM alpine:3.8
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]