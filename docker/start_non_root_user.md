# Docker container 에서 root가 아닌 유저로 실행하고 싶을 때

기본적으로 container의 명령어는 root 유저로 실행하게 됩니다. 

그런데 가끔 non root user로 무언가를 실행해야하는 경우가 발생합니다.

참여한 프로젝트의 바이너리가 root user일 경우 실행되 안되도록 설계가 되어 있어

user를 생성해야 합니다.

그런데 많은 Dockerfile 샘플에서는 `USER` 명령(instruction)을 사용하면 된다고 합니다.

그러나 user add에 대한 설명이 하나도 없었습니다.