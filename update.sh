#create the network
# docker network create authRedisNet

docker pull fymadigan/gateway_studybuddy
docker pull fymadigan/db_studybuddy

docker rm -f redisServer
# docker network rm authRedisNet

docker run -d --name redisServer --network authRedisNet redis

docker rm -f mysqlauth
docker rm -f gateway_studybuddy
docker rm -f summary

export SQL_PWD=1234
export SQL_DB=auth
export SESSIONKEY=1234
export REDISADDR=redisServer:6379


docker run -d \
-p 3306:3306 \
--name mysqlauth \
-e MYSQL_ROOT_PASSWORD=$SQL_PWD \
-e MYSQL_DATABASE=$SQL_DB \
--network authRedisNet \
fymadigan/db

export DSN=root:$SQL_PWD@tcp\(mysqlauth:3306\)/$SQL_DB

docker run -d \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=/etc/letsencrypt/live/studybuddy.kaylalee.me/fullchain.pem \
-e TLSKEY=/etc/letsencrypt/live/studybuddyapi.kaylalee.me/privkey.pem \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e DSN=$DSN \
--name gateway_studybuddy \
--network authRedisNet \
fymadigan/gateway_studybuddy  