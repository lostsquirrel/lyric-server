# lyric-server
self hosted lyric server for synology lyrics plugin

## test
curl "localhost:8000/search?artist=a&title=b"
curl "localhost:8000/lyric?id=a-b.lrc"

## build
VERSION=v1.0
docker build -t lostsquirrel/lyric-server:$VERSION .
docker push lostsquirrel/lyric-server:$VERSION


## run

docker run -d --name lyric-server \
-v /volume2/music/lyrics:/lyrics \
-e  LYRCS_PATH=/lyrics \
-p 18000:8000 \
lostsquirrel/lyric-server:$VERSION