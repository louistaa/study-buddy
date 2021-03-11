echo "is building"
GOOS=linux go build
docker build -t fymadigan/gateway_studybuddy .
go clean
