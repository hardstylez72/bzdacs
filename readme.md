CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t hardstylez72/bblog .
docker login
docker push hardstylez72/bblog:latest
 