docker rm -f sb-client

docker pull aryasaatvik/sb-client

docker run -d \
 --name client \
 -p 80:80 -p 443:443 \
 -v /etc/letsencrypt:/etc/letsencrypt:ro \
 aryasaatvik/sb-client