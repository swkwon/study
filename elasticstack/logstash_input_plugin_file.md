# File Input Plugin

## File Input Options

### close_older

* value type: number 혹은 string duration
* 기본값: `"1 hour"`

이 옵션은 마지막으로 읽은 시간이 지정한 시간 이상 지나가면 파일을 닫는 옵션이다. 만약 파일이 begin tailed나 read 이냐에 따라 다른 결과를 미친다. 만약 tailing이면, 그리고 긴 시간동안 데이터가 들어오지 않는다면 파일은 closed될 수 있지만(다른 파일을 열 수 있음) 새 데이터가 감지되면 다시 열 수 있도록 대기 한다. 읽는 경우 파일은 마지막 바이트를 읽은 시점 부터 closed_older 초 후에 닫힌다. 

### delimiter

* value type: string
* 기본값: `"\n"`

압축된 파일을 읽을 때는 이 설정이 사용되지 않는다. 대신 표준 윈도우즈 또는 유닉스 줄 끝이 사용 됩니다.

### discover_interval

* value type: number
* 기본값: `15`

`path` 옵션에서 파일 이름 패턴을 확장하여 감시 할 새 파일을 찾는 빈도이다. 이 값은 `stat_interval`의 배수 이다. 예를 들어 `stat_interval`값이 `500ms`이면 `15`*`500ms`이며 7.5초 동안 실제로 새로운 콘텐츠를 읽는 데 걸리는 시간을 고려해야 하기 때문에 이것이 가장 좋은 경우이다.

### exclude

* value type: array
* 기본값 없음

filename 패턴 매칭으로 파일을 읽을 때 특정 파일은 제외 시킬 수 있다. 만약 아래와 같이 설정했다면,

```
path => "/var/log/*"
```

/var/log/ 디렉토리의 모든 파일을 읽으려고 하지만, 아래와 같이 exclude를 사용하면,

```
exclude => "*.gz"
```

`/var/log/` 아래 모든 gz 파일은 읽지 않는다.

### file_chunk_count

* value type: number
* 기본값: `4611686018427387903`

`file_chunk_size`와 함께 사용하면 이 옵션은 다음 활성 파일로 이동하기 전에 각 파일에서 읽을 청크 수를 설정한다. 예를 들어 `file_chunk_count`가 32이고 `file_chunk_size`가 32KB는 각 활성 파일에서 다음 1MB를 처리합니다. 기본값이 매우 크기 때문에 다음 활성 파일로 이동하기 전에 EOF를 효과적으로 읽을 수 있다.

### file_chunk_size

* value type: number
* 기본값: `32768` (32KB)

파일 내용은 블록 또는 청크로 디스크에서 읽혀지며 청크에서 라인이 추출된다. 이 설정을 왜 언제 변경하는지 보려면 file_chunk_count를 참조 한다.

### file_completed_action

* value: `delete`, `log`, `log_and_delete`
* 기본값: `delete`

`read` 모드일 때 파일 읽기가 완료 될 때 액션을 수행한다. 만약 `delete`로 명시 되어있을 때 파일은 지워진다. `log`일 경우 `file_completed_log_path`에 설정된 path에 기록된다. `log_and_delete`일 경우 두가지가 모두 수행된다.

### file_completed_log_path

* value type: string
* 기본값: 없음

`file_completed_action`이 `log` 또는 `log_and_delete` 일 경우에만 지정한다. 지정 경로에 저장되는 파일은 추가만 되어 매우 커질 수 있다. 파일 관리에 대한 책임은 사용자에게 있다.

### file_sort_by

* value: `last_modified`, `path`
* 기본값: `last_modified`

파일을 정렬하는 기준이다. 기본값은 마지막으로 수정된 날짜이다. `path`일 경우 알파벳 순으로 정렬된다.

### file_sort_direction

* value: `asc`, `desc`
* 기본값: `asc`

파일 정렬 방법이다. `asc`는 오름차순, `desc` 내림차순이다. `last_modified` + `asc`일 경우 오래된 데이터 부터 읽게 된다. 

### ignore_older

* value type: number or string_duration
* 기본값: 없음

지정한 기간 전에 수정된 파일은 무시된다. number만 입력하였을 경우 초단위 계산을 한다. 설정하지 않으면 무시되는 파일은 없다.

### max_open_files

* value type: number
* 기본값: 없음

한번에 읽을 수 있는 파일의 갯수를 설정한다. 이 숫자보다 많은 파일을 읽고 싶을 경우 close_older를 사용한다. 내부적으로는 기본값이 4095이다.

### mode

* value: `tail`, `read`
* 기본값: `tail`

파일 읽는 방식을 선택할 수 있다. `tail`은 몇 개의 파일을 추적하거나
많은 내용-완전한 파일을 읽다. `read` 모드는 이제 gzip파일을 처리 할 수 있도록 지원한다. `read`모드 일 경우 아래 옵션이 무시된다.

1. `start_position` (항상 파일의 처음 부터 읽는다)
2. `close_older` (EOF를 만나면 자동으로 파일을 닫는다)

만약 `read`모드일 경우 아래 옵션은 주의를 해야 한다.

1. `ignore_older`
2. `file_completed_action` 
3. `file_completed_log_path`

### path

* 이 옵션은 `필수` 설정이다.
* value type: array
* 기본값: 없음

읽을 파일이 있는 위치를 설정한다. 파일이름은 패턴을 사용할 수 있다. ex)`/var/log/*.log`. 만약 `/var/log/**/*.log`와 같이 설정하면 `/var/log` 하위의 모든 log파일을 찾는다. 이 값은 반드시 절대경로로 입력해야 한다.

### sincedb_clean_after

* value type: number or string_duration
* 기본값: 2 weeks
* 숫자만 입력할 경우 days로 계산한다. 0.5일 경우 12시간이다. 

sincedb는 마지막으로 활성화된 timestamp를 기록한다. 추적하던 파일의 변경사항이 감지 되지 않으면 설정한 일시 뒤에 추적기록이 만료된다.

### sincedb_path

* value type: string
* 기본값: 없음

sincedb 파일이 저장되는 위치이다. 기본 위치는 `<path.data>/plugins/inputs/file` 이다. 이 값은 반드시 파일 path 이어야 한다. 디렉토리 path이면 안된다.

### sincedb_write_interval

*  value type: number or string_duration
*  기본값: `"15 seconds"`

sincedb에 저장하는 빈도시간을 지정한다.

### start_position

* value: `beginning`, `end`
* 기본값: `end`

logstash가 파일을 읽을 때 어디서 부터 시작하는지 지정한다. 기본적으로 파일 추적은 live stream과 같다. 그러므로 파일의 끝부터 읽기 시작한다. 만약에 import하고 싶은 오래된 데이터가 있다면 이 옵션을 `beginning`으로 지정하면 된다. 이 옵션을 변경하면 처음 만나는 파일에 대해서 적용된다. 이미 읽고 있는 파일은 변경된 옵션이 적용되지 않는다. 

### stat_interval

* value type: number or string_duration
* 기본값: `"1 second"`

파일이 수정되었는지 확인하는 시간 간격 값이 커지만 시스템 호출 횟수가 감소되지만 로그 라인을 감지하는 시간이 길어진다.

## string duration format

### Weeks

`w`, `week`, `weeks`를 사용한다. e.g. `"2 w"`, `"1 week"`, `"4 weeks"`

### Days

`d`, `day`, `days`를 사용한다. e.g. `"2 d"`, `"1 day"`, `"2.5 days"`

### Hours

`h`, `hour`, `hours`를 사용한다. e.g. `"4 h"`, `"1 hour"`, `"0.5 hours"`

### Minutes

`m`, `min`, `minute`, `minutes`를 사용한다. e.g. `"45 m"`, `"35 min"`, `"1 minute"`, `"6 minutes"`

### Seconds

`s`, `sec`, `second`, `seconds`를 사용한다. e.g. `"45 s"`, `"15 sec"`, `"1 second"`, `"2.5 seconds"`

### Milliseconds

`ms`, `msec`, `msecs`를 사용한다. e.g. `"500 ms"`, `"750 msec"`, `"50 msecs"`

### Microseconds

`us`, `usec`, `usecs`를 사용한다. e.g. `"600 us"`, `"800 usec"`, `"900 usecs"`

* [출처: 엘라스틱 사이트](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-file.html#plugins-inputs-file-stat_interval)