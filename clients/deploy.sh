./build.sh
docker push aryasaatvik/sb-client
ssh -i "~/sb_client.pem" saatvik@studybuddy.kaylalee.me < aws.sh
