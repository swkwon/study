# permission error

```
2016/07/01 09:14:00 [crit] 1938#0: *4 connect() to 127.0.0.1:8080 failed 
(13: Permission denied) while connecting to upstream, client: 10.0.2.2, server: 
localhost, request: "GET / HTTP/1.1", upstream: "http://127.0.0.1:8080/", 
host: "localhost:20000"
```

proxy_pass 설정 시 위와 같은 permission error가 날 수 있는데 이는 nginx가 네트워크를 액세스 하지 못해 발생하는 문제 입니다. 아래와 같이 하면 해결 됩니다.

```bash
$ setsebool -P httpd_can_network_connect 1
```

`-p` 옵션을 사용하면 리눅스 설정파일에 반영되어 리붓후에도 설정값이 사라지지 않습니다.