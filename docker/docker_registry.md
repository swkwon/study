# Docker 개인 저장소 구축

## registry server 만들기

Docker로 개인 저장소(registry)도 Docker를 이용하여 서비스를 만들 수 있습니다.

먼저 registry image를 받습니다.

```
$ sudo docker pull registry:latest
```

registry 이미지를 컨테이너로 실행 합니다.

```
$ sudo docker run -d -p 5000:5000 --name hello-registry -v /tmp/registry:/tmp/registry registry
```

## 서비스 이용 시 insecure 설정

registry를 컨테이너로 실행했을 때 서비스를 이용하는 곳에서는 https로 접근하게 됩니다. 그러나 지금 서비스가 http로 실행되었기 때문에
서비스를 이용하는 곳은 insecure 설정을 해주어야 합니다. [(공식가이드 링크)](https://docs.docker.com/registry/insecure/)

리눅스의 경우 `/etc/docker/daemon.json` 에 아래내용을 추가합니다.

```json
{
  "insecure-registries" : ["myregistrydomain.com:5000"]
}
```

만약 `daemon.json` 파일이 없다면 생성해주면 됩니다.

https 도메인으로 registry 서비스를 연결해주었다면 insecure 설정은 필요 없습니다.

## AWS S3에 이미지 올리기

registry 서버에 image를 올리지 않고 AWS S3에 image를 올릴 수도 있습니다.

```
$ sudo docker run -d -p 5000:5000 --name s3-registry \
    -e SETTINGS_FLAVOR=s3 \
    -e AWS_BUCKET=example \
    -e STORAGE_PATH=/registry \
    -e AWS_KEY=AKIABCDEFGHIJKLMNOPQ \
    -e AWS_SECRET=sF4321ABCDEFGHIJKLMNOPqrstuvwxyz21345Afc \
    registry
```

* SETTING_FLAVOR: 이미지 저장방법 입니다. S3를 사용합니다.
* AWS_BUCKET: S3의 버킷이름입니다.
* STORAGE_PATH: 이미지 저장 경로 입니다.
* AWS_KEY: 액세스 KEY 입니다.
* AWS_SECRET: 시크릿 KEY 입니다.