docker rm -f sb-mysql

docker pull aryasaatvik/sb-mysql

docker run -d -p 3306:3306 --name sb-mysql --network sb-network -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -e MYSQL_DATABASE=$DB_NAME aryasaatvik/sb-mysql