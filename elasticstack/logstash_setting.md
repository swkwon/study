# docker에서 logstash setting 파일 연결

[여기](./elk_on_docker.md)에서 filter 파일 연결까지 해보았다.

filter 파일을 수정하면 자동으로 reload를 할 수 있다. 그러기 위해서는 logstash.yml옵션을 주어야 한다.

```yml
# logstash.yml logstash 설정 파일
http.host: "0.0.0.0"
xpack.monitoring.elasticsearch.hosts:["http://elasticsearch:9200"]
config.reload.automatic: true
config.reload.interval: 3s
```
`config.reload.automatic`의 값을 `true`로 하면 자동으로 filter 파일을 리로딩 한다. `config.reload.interval`의 값이 `3s`로 되어 있는데 이는 3초마다 파일이 변경되었는지 검사한다.

docker로 logstash를 실행할 때 `logstash.yml`파일을 -v 옵션으로 연결해주면 된다.

```
$ docker run --rm -it -p 9900:9900 \
-v c:\logstash_pipeline_conf:/usr/share/logstash/pipeline \
-v c:\logstash_conf\logstash.yml:/usr/share/logstash/config/logstash.yml
--net elknet --name logstash logstash:7.6.0
```

이렇게 되면 logstash가 실행될 때 logstash.yml 파일을 읽게 된다. pipeline 볼륨설정은 pipeline conf 파일을 자동으로 읽게 된다.