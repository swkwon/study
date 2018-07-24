# MessagePack

## 1. 개요

MessagePack은 효율적인 바이너리 직렬화 포맷입니다. JSON과 같이 여러 언어에서 데이터를 교환할 수 있습니다. 그러나 JSON보다 경량화 되어 있고, 더 빠릅니다.

---

## 2. 직렬화된 구조의 예

![MSGPACK](images/msgpack.jpg)

---

## 3. release

MessagePack은 언어별로 github page가 분산 되어 있습니다. [www.msgpack.org](www.msgpack.org) 에서 언어 메뉴를 클릭하면 API 예제 상단에 라이브러리 링크가 있습니다.

---

## 4. 장점

1. schema free 구조라 유연한 개발이 가능합니다.
2. JSON 보다는 빠릅니다.
3. C++의 경우 Header only library라 라이브러리 빌드가 필요 없습니다.
4. 같은 언어라도 여러가지 라이브러리가 있습니다.
5. 사용하기가 편합니다. (protobuf나 flatbuffers 보다 더)
6. 약 50여개의 프로그래밍 언어의 라이브러리가 있습니다.(몇몇언어 에서는 JSON과 호환이 쉽다.)

---

## 5. 단점

1. schema free 구조라 msg 구조에 대해 문서화가 중요합니다.
2. protobuf 보다는 느립니다.(그러나 메모리 효율은 좋습니다.)

---

## 6. Tutorial

C++에서 직렬화 하고 Go에서 역직렬화 하는 코드를 작성해보겠습니다.

```Cpp
class Option {
public:
  Option() {}
  Option(const char* ab)
    : Ability(ab) {}

private:
  std::string Ability;
public:
  MSGPACK_DEFINE_MAP(Ability);
};

class Item {
public:
  Item() {}
  Item(const int64_t idx, const char* n, const Option& op)
    : Index(idx), Name(n), ItemOption(op) {}
private:
  int64_t Index;
  std::string Name;
  Option ItemOption;
public:
  MSGPACK_DEFINE_MAP(Index, Name, ItemOption);
};

class Inventory {
public:
  Inventory() {}
  void AddItem(const int64_t i, const Item& it) {
    Inv[i] = it;
  }
private:
  std::map<int64_t, Item> Inv;
public:
  MSGPACK_DEFINE_MAP(Inv);
};
```

Inventory 클래스는 복잡한 구조로 이루어져 있습니다. Inventory class에서 msgpack 라이브러리가 직렬화할 멤버를 알 수있도록 `MSGPACK_DEFINE_MAP(inv);` 라고 표시해 줍니다.
Item class도 마찬가지, Option class 도 마찬가지 입니다. 직렬화 방법은 아래와 같습니다.

```cpp
  Option damage("damage");
  Option critical("critical");
  
  Item damage_item(1, "damage_potion", damage);
  Item critical_item(2, "critical_potion", critical);
  
  Inventory inv;
  inv.AddItem(1, damage_item);
  inv.AddItem(2, critical_item);

  std::stringstream ss;
  msgpack::pack(ss, inv);

  msgpack::object_handle oh = msgpack::unpack(ss.str().data(), ss.str().size());
  msgpack::object obj = oh.get();
  std::cout << obj << std::endl;

  Inventory deserialize = obj.as<Inventory>();
```

Inventory 데이터를 만들고, std::stringstream 객체에 inv의 직렬화 데이터를 담습니다. `msgpack::unpack` 함수는 역직렬화 함수이며 object_handle 객체를 반환 합니다. `get()` method로 object 객체를 얻을 수 있으며, object 객체는 `as` method를 통해서 본래의 객체 `Inventory` 객체를 가져올 수 있습니다.

이 바이너리를 다른 언어에서 어떻게 역직렬화 하는지 확인 해보겠습니다.

```go
type Inventory struct {
  Inv map[int64]GameItem
}

type GameItem struct {
  Index      int64
  Name       string
  ItemOption Option
}

type Option struct {
  Ability string
}
```

스키마가 없다보니 어플마다 구조가 다르면 에러가 발생 할 수 있습니다. 예를 들어 struct GameItem에서 ItemOption이 빠져도 상관 없는데 중간 변수인 Name이 빠지면 에러가 발생 합니다.

```go
  // 직렬화 데이터
  bin := []byte{...}
  var inv Inventory
  e := msgpack.Unmarshal(bin, &inv)
  if e != nil {
    fmt.Println(e.Error())
  } else {
    fmt.Printf("%v", inv)
  }
```

`msgpack.Unmarshal`으로 바이너리 데이터를 Inventory 객체로 변환할 수 있습니다. 참고로 직렬화 하는 방법은 아래와 같습니다.

```go
  var testInv Inventory
  testInv.Inv = make(map[int64]GameItem)

  testInv.Inv[1] = GameItem{
    Index: 1,
    Name:  "damage_potion",
    ItemOption: Option{
      Ability: "damage",
    },
  }

  testInv.Inv[2] = GameItem{
    Index: 2,
    Name:  "critical_potion",
    ItemOption: Option{
      Ability: "critical",
    },
  }

  serialize, _ := msgpack.Marshal(&testInv)
```

---

## 7. 소감

1. 더 심오하게 사용하기 위해서는 study가 필요하지만 사용하기 편합니다.
2. 라이브러리 빌드가 필요 없습니다.