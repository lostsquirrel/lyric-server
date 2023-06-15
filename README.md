# lyric-server
self hosted lyric server for synology lyrics plugin

## test
curl "localhost:8000?artist=a&song=b"
curl "localhost:8000/lyric?id=a-b.lrc"

## build
VERSION=v1.6
docker build -t lostsquirrel/lyric-server:$VERSION .
docker push lostsquirrel/lyric-server:$VERSION

docker run --rm -p 8008:8000 -v /tmp/lyrics:/lyrics -e LYRCS_PATH=/lyrics lostsquirrel/lyric-server:v1.4
## run

docker run -d --name lyric-server \
-v /volume2/music/lyrics:/lyrics \
-e  LYRCS_PATH=/lyrics \
-p 18000:8000 \
lostsquirrel/lyric-server:$VERSION