# Protocol Buffers

## 1. 개요

Protocol Buffers (이하 protobuf)는 여러언어를 지원하는 직렬화 라이브러리 입니다. protobuf는 IDL을 작성하여 원하는 언어로 컴파일 하여 사용합니다. 이점은 FlatBuffers와 비슷 합니다.

## 2. 장점

1. 여러 언어의 라이브러리가 있습니다.
2. 빠르고 메모리 효율이 좋습니다.
3. 사용하기 편합니다.

## 3. release

[https://github.com/google/protobuf/releases/](https://github.com/google/protobuf/releases/)에서 릴리즈 버전을 확인 할 수 있습니다.

## 4. Tutorial

먼저 IDL 파일을 만들어 주어야 합니다. protobuf는 `.proto` 파일에 메시지 명세서를 작성합니다.

```protobuf
syntax = "proto3";

package tutorial;

message Person
{
  string Name = 1;
  int32  Age  = 2;
}
```

syntax는 `proto2` 아니면 `proto3` 입니다. 특별한 이유가 없으면 `proto3` 문법으로 개발하는 기능이나 성능면에서 좋습니다. 두번째 줄에는 `package tutorial` 이 있습니다. 이 메시지의 package(또는 namespace)를 지정해주는 것입니다. `Person`이라는 message는 `Name`과 `Age`가 있습니다. proto파일을 모두 작성하였으면 `protoc`로 C++ code를 생성해보겠습니다.

```shell
shell> protoc --cpp_out=. address.proto
```

`address.proto`는 컴파일할 proto파일 입니다. `--cpp_out`은 cpp 파일이 출력될 폴더위치 입니다. 현재 폴더로 하였습니다. 컴파일에 성공하면 address 이름을 따와서 파일명이 `address.pb.h`와 `address.pb.cc`인 파일이 생성 됩니다. 이 파일을 사용할 프로젝트에 추가해 줍니다. 다음은 직렬화하는 cpp 코드 입니다.

```cpp
#include "address.pb.h"
#include <iostream>
int main()
{
  tutorial::Person p;
  p.set_name("william");
  p.set_age(10);
  std::string ser = p.SerializeAsString();

  return 0;
}
```

`tutorial::Person p;`로 객체를 생성합니다. `set_name`과 `set_age`로 이름과 나이를 설정합니다. `SerializeAsString()`으로 Person 객체를 serializing합니다. cpp에서는 직렬화된 byte array를 std::string에다가 담아서 줍니다. 아래는 역직렬화 하는 cpp 코드 입니다.

```cpp
  std::string ser{...}; // 직렬화되어 있는 객체
  tutorial::Person deserialize;
  deserialize.ParseFromString(ser);
  std::cout << deserialize.name() << " " << deserialize.age() << std::endl;
```

역직렬화 하려는 객체 `deserialize`를 선언해줍니다. 그리고 `ParseFromString()` 메소드를 이용하여 parsing 해줍니다. 제대로 parsing이 되면 `name`과 `age`함수 호출로 값을 읽을 수 있습니다.

### 4-1 Go에서 사용하기

Go에서 사용하려면 추가적으로 protoc-gen-go와 protobuf lib가 필요합니다. `go get`으로 설치해 줍니다.

```shell
shell> go get -u github.com/golang/protobuf/protoc-gen-go
shell> go get -u github.com/golang/protobuf/proto
```

protoc로 go 파일을 생성하는 것은 아래와 같습니다.

```shell
shell> protoc --go_out=. address.proto
```

cpp 파일을 만들때는 `--cpp_out`을 사용했지만 go 파일은 `--go_out`으로 해줍니다. 다른 언어도 비슷 합니다.

```go
package main

import (
  "ProtobufGo/tutorial"
  "fmt"

  "github.com/golang/protobuf/proto"
)

func main() {
  person := tutorial.Person{
    Name: "william",
    Age:  10,
  }
  b, e := proto.Marshal(&person)
  if e != nil {
    fmt.Println(e.Error())
  } else {
    fmt.Printf("%q\n", b)
  }

  var des tutorial.Person
  proto.Unmarshal(b, &des)

  fmt.Printf("%s %d\n", des.GetName(), des.GetAge())
}
```

위와 같이 객체를 생성해서 Marshal로 직렬화 하고 Unmarshal로 역직렬화 합니다.

## 5. 단점

1. IDL 파일을 작성해야 합니다.
2. proto 파일이 바뀔때마다 빌드해주어야 합니다.
3. C++의 경우 pb.cc 파일에 descriptor string이 자동생성 되는데 이 string length가 2^16이상이 되면 컴파일 에러가 납니다.(protobuf 2.6, x86 기준)

## 6. 소감

두차례 실무에 적용해 보았는데 proto 파일이 변경될 떄마다 재빌드 하는 것을 싫어하는 개발자도 있었습니다. 의외로 사용법에 대해서 어려워 하는 사람도 있었습니다. 개인적으로는 어려움 없이 사용했습니다.