FROM alpine:3.12.3

CMD mkdir /opt/bzdacs
WORKDIR /opt/bzdacs

COPY ./cmd/backend/backend .
COPY ./generated/swagger.json ./generated/swagger.json
COPY ./ui/dist ./ui
COPY ./migrations ./migrations
COPY config/config.build.yaml /opt/cfg/

EXPOSE 8080

CMD [ "/opt/bzdacs/backend", "-config", "/opt/cfg/config.build.yaml"]
