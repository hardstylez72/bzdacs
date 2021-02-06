### site http://bzdacs.eblog.cyou/


### build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o static .

docker build -t hardstylez72/bzdacss .
docker login
docker push hardstylez72/bzdacss:latest

docker build -t hardstylez72/bzdacsui .
docker login
docker push hardstylez72/bzdacsui:latest
 