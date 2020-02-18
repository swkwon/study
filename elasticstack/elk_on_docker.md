# docker를 이용해 ELK 구성

## elastic 제품군을 설치하기 전에

먼저 elastic 제품이 사용할 network를 구성해준다. 나는 elknet이라고 네트워크 이름을 지었다.

```
$ docker network create elknet
```

docker network를 구성하면 docker위에서 돌아가는 다른 서비스(키바나와 같은..)를 연결할 때 편리하다. 

## elasticsearch 설치하기

```
$ docker pull elasticsearch:7.6.0
```

elastic 제품들은 버전 tag를 정확히 적어주어야 한다. 

```
$ docker run -d --name elasticsearch --net elknet -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.0
```

위와 같이 docker container를 실행하고 웹브라우저에서 접근되는지 확인해보자.

```
http://localhost:9200
```

출처
* [엘라스틱서치 도커 허브](https://hub.docker.com/_/elasticsearch)

Production mode
* [참고](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/docker.html)

## kinaba 설치하기

```
$ docker pull kinaba:7.6.0
```

아래 실행은 기본 설정으로 `http://localhost:9200 #elasticsearch 기본 위치`에 자동으로 연결한다.

```
$ docker run -d --name kibana --net elknet -p 5601:5601 kibana:7.6.0
```

kibana연결이 잘되는지 웹브라우저로 확인한다.

```
http://localhost:5601
```

출처
* [키바나 도커 허브](https://hub.docker.com/_/kibana)

Production mode
* [참고](https://www.elastic.co/guide/en/kibana/current/docker.html)

## logstash 설치하기

```
$ docker pull logstash:7.6.0
```

기본설정으로 실행해보자.

```
$ docker run --rm -d --net elknet --name logstash logstash:7.6.0
```

logstash는 pipeline 구성을 해주어야 실행이 되지만, 위와 같이 그냥 실행할 경우 기본 설정으로 실행된다. 기본 pipeline은 아래와 같다.

```
input {
    beats {
        port => 5044
    }
}

output {
    stdout {
        codec => rubydebug
    }
}
```

나는 아래와 같이 host 머신에 `test.conf`파일을 만들어서 logstash container 실행 시 이 파일을 읽도록 하려고 한다.

```
input {
    tcp {
        port => 9900
    }
}

output {
    elasticsearch {
        hosts => ["elasticsearch"]
        codec => rubydebug
    }
    stdout {
        codec => rubydebug
    }
}
```

tcp port 9900으로 부터 메시지를 받아서 elasticsearch와 stdout으로 rubydebug 형식으로 출력 한다. elasticsearch.hosts의 값이 `"elasticsearch"`인데 container 이름이 elasticsearch를 찾아서 연결한다. docker network를 구성한 이유이기도 하다.

conf파일을 연결하여 재실행해보자. 위와 같이이 설정을 파일에 저장한다.

```
$ docker run --rm -it -v c:\logstash:/usr/share/logstash/pipeline -p 9900:9900 --network elknet --name logstash logstash:7.6.0
```

-it 옵션으로 stdin, stdout을 실행할 수 있다.

netcat같은 도구를 이용하여 9900 port로 메시지를 보내보자.

```
echo "hello world" | nc localhost 9900
```

kibana에서 조회 하니 아래와 같은 document가 있다.

```
{
  "_index": "logstash-2020.02.17-000001",
  "_type": "_doc",
  "_id": "mcGsUXABnMkHFyjaAgVS",
  "_version": 1,
  "_score": null,
  "_source": {
    "@version": "1",
    "host": "gateway",
    "port": 43866,
    "message": "hello world\r",
    "@timestamp": "2020-02-17T05:43:12.082Z"
  },
  "fields": {
    "@timestamp": [
      "2020-02-17T05:43:12.082Z"
    ]
  },
  "sort": [
    1581918192082
  ]
}
```

