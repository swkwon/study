# crontab 사용하기

## 기본 명령어

```bash
$ crontab -e // 편집
$ crontab -l // 리스트 보기
$ crontab -r // 삭제
```

## crontab 예약 설정

```
*             *             *             *             *
minute(0-59)  hour(0-23)    day(1-31)     month(1-12)   dayOfWeek(0-6:sun-sat)
```

## 예제

```
# 매 분 test.sh 실행
* * * * * /home/myspace/test.sh

# 매주 금요일 오전 5시 45분에 test.sh 실행
45 5 * * 5 /home/myspace/test.sh

# 매일 매시간 0분, 30분에 test.sh 실행
0,30 * * * * /home/myspace/test.sh

# 매일 1시 0분 부터 30분까지 매 분 마다 test.sh 를 실행
0-30 1 * * * /home/myspace/test.sh

# 매 10분 마다 test.sh를 실행
*/10 * * * * /home/myspace/test.sh
```

## 로깅

```
# 실행할 때 마다 로그파일 overwrite
* * * * * /home/myspace/test.sh > test.sh.log 2>&1

# 실행할 때 마다 로그파일 append
* * * * * /home/myspace/test.sh >> test.sh.log 2>&1

# 로그 안남기기
* * * * * /home/myspace/test.sh >> /dev/null 2>&1
```

## 백업

```
# shell에서 직접 커맨드
$ crontab -l > /home/myspace/backup/crontab.bak

# crontab에서 자동화 매일 0시 0분에 백업
0 0 * * * crontab -l > /home/myspace/backup/crontab.bak
```

## 주의사항

crontab은 유저 환경변수를 읽지 못합니다. 그러므로 script 파일을 만들어서 사용하는 환경변수를 직접 export 명령으로 설정해주어야 합니다. 만약 소스코드에서 DB설정을 환경변수로 부터 읽는 다고 합시다.

```go
// my_app이란 프로그램에서 아래와 같이 환경변수를 가져오는 경우가 있다고 가정해 봅시다.
// 아래코드는 Golang 코드이며, Getenv함수는 환경변수를 가져오는 함수 입니다.
database := os.Getenv("DATABASE")
user := os.Getenv("USER")
password := os.Getenv("PASSWORD")
```

```
# 쉘스크립트에서 export로 환경변수를 작성해주고, 프로그램을 실행합니다.
export DATABASE=my_database
export USER=root
export PASSWORD=1234

/home/myhome/bin/my_app
```