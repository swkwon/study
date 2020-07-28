# cmd에서 text검색

리눅스의 grep과 같은 기능을 하는 find가 있다.

```powershell
C:\> dir | find "Wind"
2019-06-27  오전 09:27    <DIR>          Windows
```

텍스트파일 에서 검색할 때는 아래와 같이 합니다.

```powershell
C:\> type log.txt | find "{search-string}"
```