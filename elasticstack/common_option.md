# 모든 input plugin에서 사용할 수 있는 옵션

## Common options

아래 옵션은 모든 input plugin에 적용할 수 있는 옵션이다. 

### add_field

* value type: hash
* 기본값: `{}`

이벤트에 필드 추가

### codec

* value type: codec
* 기본값: `"plain"`

codec은 input data를 위해 사용한다. 입력 codec은 logstash 파이프라인에 별도의 필터가 필요 없이 입력 될 수 있도록 데이터를 디코딩할 수 있는 편리한 방법이다.

### enable_metric

* value type: boolean
* 기본값: true

메트릭 수집에 대한 여부를 지정한다.

### id

* value type: string
* 기본값: 없음

플러그인 구성에 고유한 ID를 추가한다. 만약 ID를 지정하지 않으면 logstash 자체적으로 하나 생성한다. 반드시 ID를 추가해주는 것을 추천한다. 동일한 유형의 플러그인이 2개 이상 있을 경우 유용하다. logstash의 모니터링 api를 통해 모니터링 할 때 이 ID가 도움을 준다. 

### tags

* value type: array
* 기본값: 없음

이벤트에 임의의 태그를 추가 추가한다.

### type

* value type: string
* 기본값: 없음

이 입력으로 처리되는 모든 이벤트에 type 필드를 추가한다. type은 필터 활성화 하는데 주로 사용된다. type은 이벤트 자체의 일부로 저장되므로, 그래서 키바나에서 검색하는데 사용할 수 있다.

* [출처: 엘라스틱 사이트](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-file.html#plugins-inputs-file-stat_interval)