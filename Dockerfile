FROM alpine:3.12.3

CMD mkdir /opt/bblog
WORKDIR /opt/bblog
COPY ./cmd/server/main .
COPY ./cmd/server/config.yaml .
COPY ./migrations ./migrations

CMD ["/opt/bblog/main", "-config", "/opt/bblog/config.yaml"]