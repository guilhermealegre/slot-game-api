FROM golang:1.23-alpine as builder

COPY .. /go/src/slot-games-api
WORKDIR /go/src/slot-games-api

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o slot-games-api cmd/server/main.go

RUN chown root:root slot-games-api
RUN chown 755 slot-games-api

FROM alpine:latest

COPY --from=builder --chown=root:root /go/src/slot-games-api/slot-games-api .

RUN apk --no-cache add ca-certificates \
    curl \
    bash

RUN mkdir -p migrations
RUN mkdir -p conf
RUN mkdir -p internal/swagger/docs

COPY conf conf
COPY migrations migrations
COPY internal/swagger/docs internal/swagger/docs


CMD ["./slot-games-api"]
