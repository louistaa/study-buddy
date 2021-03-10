docker network rm sb-network
docker network create sb-network

docker rm -f redis
docker run -d --name redis --network sb-network redis

export MYSQL_ROOT_PASSWORD=$(openssl rand -base64 18)
export DB_NAME=sb-sqlserver

docker rm -f sb-mysql

docker pull aryasaatvik/sb-mysql

docker run -d -p 3306:3306 --name sb-mysql --network sb-network -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -e MYSQL_DATABASE=$DB_NAME aryasaatvik/sb-mysql

docker rm -f sb-gateway

docker pull aryasaatvik/sb-gateway

export TLSCERT=/etc/letsencrypt/live/studybuddy-api.kaylalee.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/studybuddy-api.kaylalee.me/privkey.pem
export SESSIONKEY=$(openssl rand -base64 18)
export ADDR=":443"
export REDISADDR=redis:6379
export DSN=root:$MYSQL_ROOT_PASSWORD@tcp\(sb-mysql:3306\)/$DB_NAME

docker run -d --name sb-gateway \
--network sb-network \
--hostname studybuddy-api.kaylalee.me \
--domainname studybuddy-api.kaylalee.me \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e ADDR=$ADDR \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e SESSIONKEY=$SESSIONKEY \
-e REDISADDR=$REDISADDR \
-e DSN=$DSN aryasaatvik/sb-gateway

exit