GOOS=linux go build
docker build -t aryasaatvik/sb-gateway .
go clean
docker push aryasaatvik/sb-gateway