# AWS EC2에서 redis 시작

## 1. EC2 생성하기

AWS console에서 EC2를 만듭니다. AMI 리눅스로 생성합니다.

## 2. EC2에 docker 설치

생성된 후 EC2 인스턴스가 실행되었다면 ssh로 접근합니다.
그리고 yum을 통해 docker를 설치합니다.

```
$ sudo yum update -y

$ sudo yum install -y docker
```

`-y` 옵션은 무언가 물어볼 때 yes를 자동으로 선택합니다.

## 3. docker 서비스 시작

```
$ sudo service docker start

$ sudo chkconfig docker on
```

위 명령어는 docker를 시작 하는 것이고, 아래 명령어는 EC2인스턴스가 재시작 되었을 때 docker service 를 자동 시작 되도록 설정하는 것입니다.

## redis 시작

redis image를 다운로드합니다.

```
$ sudo docker pull redis
```

redis를 시작합니다.

```
$ sudo docker run --name dev-redis -d -p 8000:6379 -v /host/data/directory:/data --restart unless-stopped redis redis-server --appendonly yes
```

`-p 8000:6379` : host 8000 포트로 접근하는 연결을 컨테이너 6379 포트로 연결해줍니다.

`-v /host/data/directory:/data` : 컨테이너 /data 디렉토리를 host /host/data/directory 디렉토리에 연결합니다. docker hub에서 다운로드한 redis image는 데이터베이스를 /data에 저장하도록 되어 있습니다. 컨테이너가 삭제되어도 redis 데이터는 EC2의 /host/data/directory 디렉토리에 있습니다. 이제 컨테이너가 없어져도 걱정없습니다.

`--restart unless-stopped` : 컨테이너 재시작 옵션입니다. 컨테이너가 시작하는 정책은 4가지가 있습니다.

플래그|설명
---|---
no|컨테이너가 자동으로 시작하지 않도록 설정합니다. (기본값)
on-failure|0이 아닌 종료 코드로 오류로 인해 컨테이너가 종료되면 컨테이너를 자동으로 다시시작 합니다.
unless-stopped|컨테이너가 명시적으로 중지되지 않았을때, docker 서비스가 중지 되거나 다시시작 되지 않는 한 컨테이너를 다시 시작합니다.
always|컨테이너가 stop 되어있다면 항상 재시작합니다.

`--appendonly yes` : redis 데이터 파일 저장 방법입니다. [(참고)](http://www.redisgate.com/redis/configuration/persistence.php)

## redis-cli만 설치하고 싶을 때

gcc를 설치합니다.

```
$ sudo yum -y gcc
```

redis-cli를 다운받아 make해줍니다.

```
$ wget http://download.redis.io/redis-stable.tar.gz && tar xvzf redis-stable.tar.gz && cd redis-stable && make
```

redis-cli를 /usr/bin 폴더에 복사해줍니다.

```
$ sudo cp src/redis-cli /usr/bin
```

## 컨테이너 shell에 접근하기

```
$ sudo docker exec -it dev-redis /bin/bash
```

dev-redis (name)대신 container id를 사용하셔도 됩니다.

```
$ sudo docker exec -it d539aec81e45 /bin/bash
```