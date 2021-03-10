GOOS=linux go build
docker build -t aryasaatvik/gateway .
go clean
docker push aryasaatvik/gateway