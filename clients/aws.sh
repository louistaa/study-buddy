docker rm -f sb-client

docker pull aryasaatvik/sb-client

docker run -d \
 --name sb-client \
 -p 80:80 -p 443:443 \
 -v /etc/letsencrypt:/etc/letsencrypt:ro \
 aryasaatvik/sb-client