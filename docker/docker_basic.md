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

리눅스에서 관리자 권한을 위해 sudo를 항상 입력해야 합니다. 입력하지 않는 방법은 두가지가 있습니다.

1. root 계정으로 로그인 한다.
2. 사용하는 계정을 docker 그룹에 추가한다.

보안상 root 계정으로 로그인 하는 방법은 좋지 않습니다. docker 그룹에 해당 계정을 추가해보겠습니다.

```
$ sudo usermod -aG docker ${USER}
$ sudo service docker restart
```

위와 같이 계정을 docker그룹에 추가 후 docker 서비스를 재시작 합니다. 그리고 현재 계정을 로그아웃 한 후 로그인 합니다.