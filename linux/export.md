# 리눅스 환경변수 설정

## 임시(현재 세션용)

```bash
$ export MY_HOME=/home/myspace
```

## 영구적 적용(모든 세션용)

bashrc를 편집합니다.

```bash
$ sudo vi /etc/bashrc
```

bashrc마지막 줄에 삽입합니다.

```bash
export MY_HOME=/home/myspace
```

저장한 다음 아래 커맨드를 실행합니다.

```bash
$ source /etc/bashrc
```

## 확인하는 법

```bash
$ env
```