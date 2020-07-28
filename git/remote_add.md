# git remote 저장소와 local 저장소 연결하기

코딩하다가 git remote 저장소(github...)에 올리고 싶을 때 먼저 local 저장소를 만듭니다.

```bash
$ git init
$ git add .
$ git commit -m "first commit"
```

그리고 remote 저장소를 만든 후 아래와 같이 커맨드를 입력합니다.

```bash
$ git remote add origin https://github.com/.../repository_name.git
$ git push -u origin master
```