./build_servers.sh
docker push fymadigan/gateway_studybuddy
docker push fymadigan/db_studybuddy

ssh ec2-user@studybuddy.kaylalee.me 'bash -s' < update.sh