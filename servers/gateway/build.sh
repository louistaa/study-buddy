echo "is building"
GOOS=linux go build
docker build -t fymadigan/gateway_studybuddy .
go clean
# GOOS=linux go build
# docker build -t aryasaatvik/sb-gateway .
# go clean
# docker push aryasaatvik/sb-gateway
