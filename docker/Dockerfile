#------------------------------------------------------------------------------------------------
# Build image stage
FROM golang:1.21.3-alpine3.18 AS build

WORKDIR /build

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./source ./source

RUN go build -o /application source/cmd/service/main.go

#------------------------------------------------------------------------------------------------
# Production image stage
FROM alpine:3.18

ARG LISTEN_PORT

COPY --from=build /application /
COPY ./source/.env /
COPY ./deploy/external /external

EXPOSE ${LISTEN_PORT}

ENTRYPOINT ["/application"]