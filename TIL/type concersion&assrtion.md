# Type conversion 과 Type assertion  
```
v1 := int(n)  // type conversion
v2, ok := n.(int) // type assertion
```  
v1도 정수이고 v2도 정수일 것이다(에러가 나지 않는다면)  

## type conversion  
type casting. Golang에서는 명시적으로 형변환을 해줘야 함. 자동으로 형변환을 해주는 일 따위는 없음.  
ex) int형을 int64형으로 바꾸고 싶다면 아래와 같이 명시적으로 코드 작성해야 함.  
```
var n int = 15
var v1 int64 = int64(n)
```  

## type assertion  
결론부터 말하자면, interface type의 value type을 확인하기 위함.  

assert: If someone **asserts** a fact or belief, they state it firmly.  
사실임을 주장하는 것. (conversion은 전환, 개조)  
type assertion은 형 변환과는 상관이 없음.  
> type assertion provides access to an interface value's underlying concrete value  
인터페이스가 가지고 있는 실제 값(concrete value)에 접근할 수 있게 함.  
Golang의 인터페이스는 임의의 타입 값을 가질 수 있기 때문에 concrete라는 표현을 쓴 듯하다.  

interface type의 x와 타입 T를 x.(T)로 표현했을 때,  
* x가 nil이 아니며,  
* x는 T 타입에 속한다는 점을  

확인하기(assert)하는 것으로 이러한 표현을 Type assertion 이라고 부른다.  
* x가 nil이거나 x의 타입이 T가 아니라면 -> Runtime ERROR  
* x가 T타입인 경우 -> T 타입의 x를 리턴한다.  

```
package main

func main() {
  var a interface{} = 1
  
  i := a        // a와 i는 dynamic type이고 값은 1
  j := a.(int)  // j는 int 타입이고, 값은 1
  
  println(i)  // 포인터 주소 출력
  println(j)  // 1 출력
}
```  
* 여기서 value a를 담은 i를 출력했을 떄 포인터의 주소가 출력되는 부분을 주목!  
즉, interface value의 직접적인 값을 얻기 위해서는 **type assertion** 을 반드시 사용해야한다.  



```
v, ok := n.(int)
```  
어떤 인터페이스 값이 특정 타입인지 확인하기 위해서 ok 값을 받을 수도 있다.  
만약 n의 타입이 int가 아니라면 v는 int의 zero value인 0이 될 것이고, ok는 false가 될 것이다.  
(타입이 int가 맞다면 int 타입의 n을 리턴하고, ok는 true가 됨)  




## 출처  
https://2kindsofcs.tistory.com/13  
https://iamjjanga.tistory.com/47#recentComments  
