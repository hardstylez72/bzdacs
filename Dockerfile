FROM alpine:3.12.3

CMD mkdir /opt/bzdacs
WORKDIR /opt/bzdacs
COPY ./cmd/server/main .
#COPY ./cmd/server/config.yaml .
COPY ./migrations ./migrations

CMD ["/opt/bzdacs/main", "-config", "/opt/cfg/config.yaml"]