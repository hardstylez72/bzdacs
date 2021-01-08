FROM nginx:1.19.2-alpine

COPY ./ui/nginx.conf /etc/nginx/conf.d/nginx.conf
COPY ./ui/dist /usr/share/nginx/html

CMD mkdir /opt/bblog
WORKDIR /opt/bblog
COPY ./cmd/server/main .
COPY ./cmd/server/config.yaml .

EXPOSE 6667 5432

CMD ["/opt/bblog/main", "-config", "/opt/bblog/config.yaml"]