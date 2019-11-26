FROM golang:1.13-alpine

RUN mkdir -p /opt/code/
WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/roster cmd/roster/main.go

FROM alpine

WORKDIR /app

COPY --from=0 /opt/code/bin/roster /app/
COPY --from=0 /opt/code/configs/config.yml /app/configs/config.yml

ENTRYPOINT ["./roster"]