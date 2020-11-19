# 컨테이너간 통신

* 컨테이너의 네트워크는 브릿지로 생성된다. 
* 컨테이너 시작 시 IP는 자동할당된다. 
* 다른 컨테이너의 IP를 알 수 없으니 hostname으로 접근한다.
* 접속할 컨테이너는 같은 네트워크에 속해야 한다.

# server

`server.go`를 보면 `:9000` 으로 listen한다.

# client

`client.go`를 보면 접속할 host정보가 `server:9000` 이다.

# docker-compose.yaml

```yaml
version: "3.7"

services: 
  server:
    build:
      context: ./
      dockerfile: Dockerfile.server
    image: server:latest
    expose:
      - 9000
    networks:
      - mynet
  client:
    build:
      context: ./
      dockerfile: Dockerfile.client
    image: client:latest
    networks: 
      - mynet
    depends_on: 
      - server
networks: 
  mynet:
```

server는 9000포트를 노출 시키고 `mynet`이라는 network에 연결한다.

client도 `mynet` network에 연결한다. `depends_on`을 설정한 이유는 client가 server보다 먼저 실행되면 연결 오류가 되기 때문에 server가 정상적으로 실행이 되었을 때 client를 실행시키기 위함이다.

# 서버와 클라이언트 시작

```
$ docker-compose up -d --build
```


# 서버의 로그

```
$ docker-compose logs -f server
server_1  | 2020/11/19 08:35:09 hello server. I am client.
server_1  | 2020/11/19 08:35:10 hello server. I am client.
server_1  | 2020/11/19 08:35:11 hello server. I am client.
server_1  | 2020/11/19 08:35:12 hello server. I am client.
server_1  | 2020/11/19 08:35:13 hello server. I am client.
```

위와 같이 클라이언트로 부터 메시지를 잘 받는 것을 볼 수 있다.