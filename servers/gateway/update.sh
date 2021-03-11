echo "is updating"
docker rm -f gateway_studybuddy
docker pull fymadigan/gateway_studybuddy
docker pull fymadigan/db_studybuddy