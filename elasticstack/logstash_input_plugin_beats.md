# Beats Input Plugin

## Description

이 input plugin은 logstash가 Elastic Beats 프레임워크로 부터 이벤트를 받도록 활성화 한다.

아래 예가 5044 포트를 열어 Beats로 들어오는 것을 Elasticsearch의 index로 보내주고 있다.

```
input {
    beats {
        port => 5044
    }
}

output {
    elasticsearch {
        host => ["http://localhost:9200"]
        index => "%{[@metadata][beat]}-%{[@metadata][version]}"
    }
}
```

`%{[@metadata][beat]}` 는 `metadata` 중 `beat`를 `%{[@metadata][version]}` 는 `metadata` 중 `version`을 가져와서 elasticsearch의 index를 생성한다. 예를들면 `metricbeat-7.4.0`가 index가 된다.

```
Note

만약 ILM(Information Lifecycle Management)가 지정되지 않으면
logstash가 매일 `{[@metadata][beat]}-%{[@metadata][version]}-%{+YYYY.MM.dd}`
형식으로 @timestamp 정보를 이용하여 인덱스를 만든다.
```

```
Important

여러줄에 걸친 이벤트를 제공하는 경우 이벤트 데이터를 logstash에 보내기 전에 
Filebeat에서 사용가능한 구성 옵션을 사용하여 여러줄 이벤트를 처리해야 한다. 
multiline codec plugin을 이용하여 multilne event를 처리할 수 없다. 
그렇게 하면 logstash를 시작하지 못할 수 있다.
```

## 버전화된 Beats 인덱스들

향후 스키마 변경이 elasticsearch의 기존 인덱스 및 매핑에 미치는 영향을 최소화 하려면 버전이 지정된 인덱스에 쓰도록 elasticsearch 출력을 구성한다. 색인 설정에 지정하는 패턴은 아래와 같은 이름으로 제어한다.

```
index => "%{[@metadata][beat]}-%{[@metadata][version]}-%{+YYYY.MM.dd}"
```

`%{[@metadata][beat]}`

인덱스 이름의 첫번째 파트는 메타데이터 필드의 `beat`의 값이다. 예를들면 `filebeat`이다.

`%{[@metadata][version]}`

인덱스 이름의 두번째 파트는 Beat의 버전이다. 예를들면 `7.6.0`이다.

`%{+YYYY.MM.dd}`

인덱스 이름의 세번째 파트는 logstash의 `@timestamp` 필드 값이다.

이 설정의 결과는 아래와 같이 만들어 진다.

`filebeat-7.6.0-2020-02-19`.

## Beats Input Configuration Options

아래는 이 플러그인의 옵션들이다. 공통 옵션은 [여기](./common_option.md)를 참조 한다.

### add_hostname

* value type: boolean
* 기본값: `false`

6.0.0에서 deprecated 되었다. `hostname` 필드에서 beat가 제공한 값을 사용하여 `host`필드를 이벤트에 추가할지 여부를 결정하는 플래그 이다.

### cipher_suites

* value type: array
* 기본값: `java.lang.String[TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256]@459cfcca`

우선 순위별로 나열한 사용할 cipher suites 목록 이다.

### client_inactivity_timeout

* value type: number
* 기본값: `60`

비활동적인 클라이언트는 설정한 시간(초)후 연결을 끊어버린다.

### host

* value type: string
* 기본값: `"0.0.0.0"`

listen socket의 IP 주소이다.

### include_codec_tag

* value type: boolean
* 기본값: true

### port

* 이 옵션을 필수이다.
* value type: number
* 기본값: 없음

listen socket의 port이다.

### ssl

* value type: boolean
* 기본값: `false`

기본적으로 이벤트는 plain text로 전송된다. `ssl`을 true로 하면 암호화된 이벤트를 보낼 수 있다. `ssl_certificate`와 `ssl_key` 옵션을 같이 사용한다.

### ssl_certificate

* value type: path
* 기본값: 없음

SSL certificate를 사용할 때 쓴다.

### ssl_certificate_authorities

* value type: array
* 기본값: `[]`

권한들에 대해 클라이언트 인증서를 검증한다. 여러 파일 또는 경로를 정의 할 수 있다. 모든 인증서를 읽고 신뢰 저장소에 추가한다. 확인을 사용하려면 `ssl_verify_mode` 옵션의 값을 `peer` 또는 `forece_peer`로 사용한다.

### ssl_handshake_timeout

* value type: number
* 기본값: `10000`

ssl handshake timeout 값이다. 단위는 millisecond 이다.

### ssl_key

* value type: path
* 기본값: 없음

ssl key 이다. 이 키는 PKCS8 포맷을 필요로 한다. OpenSSL로 변경할 수 있다.

### ssl_key_passphrase

* value type: password
* 기본값: 없음

ssl key 암호이다.

### ssl_verify_mode

* value: `none`, `peer`, `force_peer`
* 기본값: `none`

기본적으로 서버는 어떤 클라이언트도 인증을 하지 않는다. 

`peer`는 클라이언트에게 제공된 자격을 물어본다. 클라이언트에게 자격이 있으면 인증된 것이다. 

`force_peer` 기본적으로 `peer`와 같다. 만약 클라이언트가 자격이 없으면 연결은 끊긴다.

이 옵션은 `ssl_certificate_authorities`와 함께 사용한다.

### ssl_peer_metadata

* value type: boolean
* 기본값: `false`

이벤트의 메타데이터에 클라이언트 인증서 정보를 저장할 수 있다. 이 옵션은 `ssl_verify_mode`가 `peer`이거나 `force_peer`일때 유효하다.

### tls_max_version

* value type: number
* 기본값: `1.2`

암호화 연결에 허용되는 최대 TLS 버전이다. 1.0은 TLS 1.0, 1.1은 TLS 1.1, 1.2은 TLS 1.2이다.

### tls_min_version

* value type: number
* 기본값: `1`

암호화 연결에 허용되는 최소 TLS 버전이다. 1.0은 TLS 1.0, 1.1은 TLS 1.1, 1.2은 TLS 1.2이다.

* [출처: 엘라스틱 사이트](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-beats.html)

