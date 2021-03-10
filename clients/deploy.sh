./build.sh
docker push aryasaatvik/sb-client
ssh -i "~/sb_client.pem" saatvik@ec2-3-90-214-253.compute-1.amazonaws.com < aws.sh
