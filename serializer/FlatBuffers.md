# 개요
FlatBuffers는 플랫폼 종속성 없이 사용가능한 직렬화 라이브러리 입니다. 

메모리 제약이 큰 모바일 하드웨어나 성능요구 사항이 가장 높은 게임같은 응용프로그램에 중점을 두고 있습니다.

# why
FlatBuffers를 쓰는 이유는?
1. runtime에 serialize, deserialize가 필요 없습니다.
2. 메모리 효율성이 높고, 빠른 속도를 보장합니다.
3. 크로스 플랫폼, 종속성 없이 사용 가능 합니다.
4. C, C++, Go, Java, C#, JavaScript, TypeScript, PHP, Python, Dart gRPC등에서 사용할 수 있습니다.
5. 필요하다면 schema-less 하게 사용할 수 있습니다.

# release
[여기](https://github.com/google/flatbuffers/releases)서 FlatBuffers release 버전을 다운 받을 수 있습니다.

# Tutorial
1. FlatBuffers `schema` 파일을 작성합니다. (.fbs)
2. FlatBuffers의 컴파일러 flatc.exe 를 사용합니다.
3. 스키마를 준수하는 JSON파일을 FlatBuffers 바이너리 파일로 구문분석 합니다.
4. 필요한 언어에 맞게 생성된 파일을 사용합니다.

## Schema 예제
```flatbuffers
// Example IDL file for our monster's schema.
namespace MyGame.Sample;
enum Color:byte { Red = 0, Green, Blue = 2 }
union Equipment { Weapon } // Optionally add more tables.
struct Vec3 {
  x:float;
  y:float;
  z:float;
}
table Monster {
  pos:Vec3; // Struct.
  mana:short = 150;
  hp:short = 100;
  name:string;
  friendly:bool = false (deprecated);
  inventory:[ubyte];  // Vector of scalars.
  color:Color = Blue; // Enum.
  weapons:[Weapon];   // Vector of tables.
  equipped:Equipment; // Union.
  path:[Vec3];        // Vector of structs.
}
table Weapon {
  name:string;
  damage:short;
}
root_type Monster;
```
스키마는 C언어와 유사합니다. 스키마의 시작은 `namespace`를 선언하는 것입니다. 

그 다음에 `enum` 정의가 있습니다. 예제에서는 enum의 타입이 byte이며 이름은 Color 입니다.
Red는 0, Blue는 2 라고 명시적으로 작성되었습니다. Green은 암시적으로 Red보다 하나 큰 1이 됩니다.

다음에 오는 enum은 `union` 입니다. 이 예제에서는 하나의 table만 사용하기 때문에 유용하지는 않습니다. 여러개의 table을 만들면 Equipment에 추가할 수 있습니다.

그 다음은 `struct Vec3` 입니다. `struct Vec3`는 3차원 벡터값을 표현합니다. 여기서 table을 사용하지 않고 struct를 사용한 이유는 변경되지 않는 데이터 구조를 표현하기에는 struct가 적합하기 때문입니다. 왜냐면 메모리를 더 적게 사용하고 더 빠른 검색이 가능하기 때문입니다.

`Monster` table은 이 예제에서 중요한 object입니다. `mana:short = 150;`은 mana라는 변수는 type이 short이며 default값으로 150이 됩니다. 명시되지 않은 모든 필드는 0이거나 null 이 됩니다. 
`friendly:bool = false (deprecated);`를 보면 friendly 필드가 deprecated가 된 것을 볼 수 있습니다. 하위호환을 위해 필드를 삭제하는 것은 위험합니다. 

`Weapon` table은 sub-table입니다. 이 예제에서는 `Monster` table과 `Equipment` enum 에서 사용합니다. 

마지막에는 `root_type`이 있습니다. `root_type`은 직렬화 되는 최상위 table이 됩니다. 

스키마에서는 short과 float을 사용했는데 short대신 int16, float대신 float32를 사용할 수 있습니다.

## Type
* 8 bit : byte(int8), ubyte(uint8), bool
* 16 bit: short(int16), ushort(uint16)
* 32 bit: int(int32), uint(uint32), float(float32)
* 64 bit: long(int64), ulong(uint64), double(float64)

## 스키마 컴파일
flatc를 이용해 fbs를 컴파일 합니다.
```bash
$ flatc --go .\monster.fbs
```

현재 폴더에 monster.fbs를 컴파일하고 생성된 파일도 현재폴더 밑에 생성됩니다. 

## 소스에서 사용하기
생성된 fbs 소스와 FlatBuffers 라이브러리를 import합니다.
```go
import (
    flatbuffers "github.com/google/flatbuffers/go"
    sample "MyGame/Sample"
)
```
buffer size가 1024 Byte인 buffer를 생성합니다.

```go
builder := flatbufers.NewBuilder(1024)
```

Weapon을 생성해보겠습니다. Sword와 Axe 두개를 생성합니다.

```go
weaponOne := builder.CreateString("Sword")
weaponTwo := builder.CreateString("Axe")
// Sword라는 첫번째 무기 생성
sample.WeaponStart(builder)
sample.Weapon.AddName(builder, weaponOne)
sample.Weapon.AddDamage(builder, 3)
sword := sample.WeaponEnd(builder)
// Axe라는 두번째 무기 생성
sample.WeaponStart(builder)
sample.Weapon.AddName(builder, weaponTwo)
sample.Weapon.AddDamage(builder, 5)
axe := sample.WeaponEnd(builder)
```

이제 Monster를 생성해보겠습니다.
```go
// 몬스터 이름 Orc 직렬화
name := builder.CreateString("Orc")
// builder에 size 10인 vector를 만듭니다.
// for loop에서는 0 - 9까지 숫자를 입력합니다.
sample.MonsterStartInventoryVector(builder, 10)
for i := 9; i >= 0; i-- {
        builder.PrependByte(byte(i))
}
inv := builder.EndVector(10)
```

weapons vector에 weapon을 추가 합니다.
```go
sample.MonsterStartWeaponsVector(builder, 2)
builder.PrependUOffsetT(axe)
builder.PrependUOffsetT(sword)
weapons := builder.EndVector(2)
```

이제 직렬화되는 몬스터 오브젝트를 만들 수 있습니다. 
```go
sample.MonsterStart(builder)
sample.MonsterAddPos(builder, sample.CreateVec3(builder, 1.0, 2.0, 3.0))
sample.MonsterAddHp(builder, 300)
sample.MonsterAddName(builder, name)
sample.MonsterAddInventory(builder, inv)
sample.MonsterAddColor(builder, sample.ColorRed)
sample.MonsterAddWeapons(builder, weapons)
sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
sample.MonsterAddEquipped(builder, axe)
sample.MonsterAddPath(builder, path)
orc := sample.MonsterEnd(builder)
```

buffer가 완성되면 finish 메소드를 호출해서 완료 해주어야 합니다.
```go
// Call `Finish()` to instruct the builder that this monster is complete.
builder.Finish(orc)
```

직렬화된 byte 배열을 얻기 위해서는 FinishedBytes() 메소드를 호출합니다.
```go
// 이 함수는 Finish() 후에 호출해야 합니다.
buf := builder.FinishedBytes() // Of type `byte[]`.
```

## Byte데이터 읽기
```go
var buf []byte = /* 직렬화된 byte 배열 */
// Get an accessor to the root object inside the buffer.
monster := sample.GetRootAsMonster(buf, 0)
```

아래와 같이 accessor를 통해 buffer 내의 값을 읽을 수 있습니다.
```go
hp := monster.Hp()
mana := monster.Mana()
name := string(monster.Name()) // Note: `monster.Name()` 은 byte[] 을 return 합니다.
pos := monster.Pos(nil)
x := pos.X()
y := pos.Y()
z := pos.Z()
// Note: Whenever you access a new object, like in `Pos()`, a new temporary
// accessor object gets created. If your code is very performance sensitive,
// you can pass in a pointer to an existing `Vec3` instead of `nil`. This
// allows you to reuse it across many calls to reduce the amount of object
// allocation/garbage collection.

// We need a `flatbuffers.Table` to capture the output of the
// `monster.Equipped()` function.
unionTable := new(flatbuffers.Table)
if monster.Equipped(unionTable) {
        unionType := monster.EquippedType()
        if unionType == sample.EquipmentWeapon {
                // Create a `sample.Weapon` object that can be initialized with the contents
                // of the `flatbuffers.Table` (`unionTable`), which was populated by
                // `monster.Equipped()`.
                unionWeapon = new(sample.Weapon)
                unionWeapon.Init(unionTable.Bytes, unionTable.Pos)
                weaponName = unionWeapon.Name()
                weaponDamage = unionWeapon.Damage()
        }
}
```

# 개인적인 소감
학습비용이 커보인다. Table 구조가 복잡해질 때를 염두해 두어야 할 것 같다.