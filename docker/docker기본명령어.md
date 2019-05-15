# docker 기본 명령어

pull : 저장소에서 이미지 가져오기

```
$ sudo docker pull redis
```

ps : 컨테이너 리스트 보기

```
$ sudo docker ps -a
```

run : 새 컨테이너 실행

```
$ sudo docker run -it --name -hello ubuntu /bin/bash
```

images : 이미지 리스트 보기

```
$ sudo docker images
```

rm : 컨테이너 삭제

```
$ sudo docker rm some-container
```

rmi : 이미지 삭제

```
$ sudo docker rmi some-image
```

restart : 컨테이너 재시작
```
$ sudo docker restart some-container
```