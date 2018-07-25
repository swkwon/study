# YAML

## 1. 개요

YAML은 사람이 읽고 쓰기 쉬운 것을 목표로한 데이터 포맷입니다. YAML의 파일은 선택적으로 `---` 으로 시작하고, `...`으로 끝납니다.

## 2. 장점

1. 사람이 읽기 쉽습니다.
2. JSON과 비교해서 주석이 있습니다.
3. 문자열에 쌍따옴표를 사용하지 않아도 됩니다.

## 3. 단점

1. 개행, 공백으로 블록을 인식하기 때문에 공백을 잘 작성해주어야 합니다. 그렇지 않으면 파싱이 불가능합니다.

## 4. Tutorial

캐릭터의 정보를 표현하는 예제를 작성해보겠습니다.

```YAML
---
character:
  name: John
  male: true
  skills:
    - healing
    - power up
    - trap
  items:
    -
      name: healing potion
      point: 100
    -
      name: mana potion
      point: 120
    -
      name: axe
      damage: 127
      options:
        -
          name: add critical
          value: +5%
        -
          name: add str
          value: +3
    -
      name: shield
      block_rate: +10%
      options:
        -
          name: add hp
          value: +10%
  description: >
    He was soldier.
    He was father.
...
```

기본적으로 key/value의 형식을 하고 있습니다. `character`라는 key에 들여쓰기한 모든 내용들이 `value`가 되는 것입니다. `name`과 `male`항목은 `character` 의 child node라고 생각할 수 있습니다. 콜론과 뒤에 오는 값은 꼭 공백이 있어야 합니다. `skills`도 `character`의 child node 입니다. 여기서는 `character`가 보유하고 있는 스킬 리스트 입니다. 리스트를 표현하기 위해서는 `skills`의 child node가 되어야 하며, `-`으로 시작 합니다. `-`과 skill의 이름은 꼭 공백이 있어야 합니다. 참고로 여기에 있는 공백은 YAML문서를 표현하는데 있어 꼭 필요한 것들이니 주의 해주셔야 합니다.

다음 `items`는 `character`가 보유하고 있는 item 리스트 입니다. 아이템의 속성값들이 있습니다. 그리고 item의 속성값 중에서 리스트 데이터를 갖는 `options`도 있습니다. 이런식으로 key/value 데이터를 리스트 형식으로 갖고 있을 수 있습니다.

description의 값은 긴 문자열인데 이 경우 한줄에 쓰기 어렵습니다. 그래서 `>`를 사용하는데 값의 new line은 space로 대체 됩니다. 만약 `>` 대신 `|`으로 작성할 경우 new line이 보존 됩니다.

## 5. 주의점

위에서도 잠깐 설명했지만 공백과 개행이 중요합니다. 그리고 key는 공백없이 작성하는 것이 좋습니다.

## 6. 소감

게임 패킷 데이터로 YAML을 쓰기에는 불필요한 공백이나 개행이 많이 들어가 부적합 합니다. 웹앱 개발 시 JSON 포맷을 많이 쓰는데 YAML으로 대체하기는 힘들것 같습니다. 다만 JSON보다 읽기 쉽고 작성하기 쉬워 설정 파일용으로 훌륭해 보입니다. 여기서 YAML의 모든 기능을 설명하지는 않았습니다. 축약이나 타입 명시 같은 좋은 기능들이 있습니다.