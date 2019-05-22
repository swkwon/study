# GPC에서 VM 인스턴스 생성 후 putty로 접속하는 방법

## puttygen으로 key 생성

1. puttygen으로 key를 생성 합니다.
1. Type은 RSA로 합니다.
1. `Generate` 버튼을 누르면 key가 생성됩니다.
1. key comment에는 user (계정)을 입력합니다.
1. key passphrase는 원하는 문자열을 입력합니다. (옵션)
1. `save private key` 버튼을 눌러 .ppk 파일을 저장합니다.
1. public key는 GCP에 등록합니다.

## GCP에 public key 등록하기

1. VM인스턴스 리스트가 나오는 대쉬보드에서 VM인스턴스 이름을 클릭합니다.
1. `수정` 버튼을 클릭 합니다.
1. 맨 아래에 `SSH 키`에 public key를 입력할 수 있는 textbox가 있습니다.
1. `저장`합니다.

이제 putty에서 ppk를 입력해주면 ssh 접근이 가능합니다.