# File Input Plugin

## close_older

* value type: number 혹은 string duration
* 기본값: `"1 hour"`

이 옵션은 마지막으로 읽은 시간이 지정한 시간 이상 지나가면 파일을 닫는 옵션이다. 만약 파일이 begin tailed나 read 이냐에 따라 다른 결과를 미친다. 만약 tailing이면, 그리고 긴 시간동안 데이터가 들어오지 않는다면 파일은 closed될 수 있지만(다른 파일을 열 수 있음) 새 데이터가 감지되면 다시 열 수 있도록 대기 한다. 읽는 경우 파일은 마지막 바이트를 읽은 시점 부터 closed_older 초 후에 닫힌다. 

## delimiter

* value type: string
* 기본값: `"\n"`

압축된 파일을 읽을 때는 이 설정이 사용되지 않는다. 대신 표준 윈도우즈 또는 유닉스 줄 끝이 사용 됩니다.

## discover_interval

* value type: number
* 기본값: `15`

`path` 옵션에서 파일 이름 패턴을 확장하여 감시 할 새 파일을 찾는 빈도이다. 이 값은 `stat_interval`의 배수 이다. 예를 들어 `stat_interval`값이 `500ms`이면 `15`*`500ms`이며 7.5초 동안 실제로 새로운 콘텐츠를 읽는 데 걸리는 시간을 고려해야 하기 때문에 이것이 가장 좋은 경우이다.