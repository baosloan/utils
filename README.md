# utils
Go开发工具包


## timeutil
时间工具包
```go
package main

import (
	"fmt"
	"github.com/baosloan/utils/timeutil"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(timeutil.GetDateTime(now)) //2024-01-30 17:07:48
	fmt.Println(timeutil.GetDateOnly(now)) //2024-01-30
	fmt.Println(timeutil.GetTimeOnly(now)) //17:07:48
	timeStr := "2024-10-01 08:00:00"
	t1, _ := timeutil.ParseInLocationShanghai(timeStr)
	t2, _ := timeutil.ParseInLocation(timeStr, timeutil.LocationShanghai)
	fmt.Println(t1) //2024-10-01 08:00:00 +0800 CST
	fmt.Println(t2) //2024-10-01 08:00:00 +0800 CST
}
```

## stringutil

字符串工具包
```go
package main

import (
	"fmt"
	"github.com/baosloan/utils/stringutil"
	"math"
)

func main() {
	fmt.Println(math.Round(float64(5) / 2))
	//求交集
	var s1 = []string{"jack", "rose", "jane"}
	var s2 = []string{"michael", "jack", "rose"}
	fmt.Println(stringutil.Intersect(s1, s2)) //[jack rose]
	//字符串切片翻转(不改变原有切片)
	var s3 = []string{"111", "222", "333", "444", "555"}
	var s4 = stringutil.Reverse(s3)
	fmt.Println(s3) //[111 222 333 444 555]
	fmt.Println(s4) //[555 444 333 222 111]
}
```

## intutil

整型工具包

```go
package main

import (
	"fmt"
	"github.com/baosloan/utils/intutil"
)

func main() {
	//求交集
	var s1 = []int{111, 222, 333}
	var s2 = []int{111, 222, 444}
	fmt.Println(intutil.Intersect(s1, s2)) //[111 222]
	//整型切片翻转(不改变原有切片)
	var s3 = []int{111, 222, 333, 444, 555}
	var s4 = intutil.Reverse(s3)
	fmt.Println(s3) //[111 222 333 444 555]
	fmt.Println(s4) //[555 444 333 222 111]
}
```

## mathutil

数学工具包

提供了一些精确计算的求和、求差、求积、小数四舍五入、向上/向下去整
```go
package main

import (
	"fmt"
	"github.com/baosloan/utils/mathutil"
)

func main() {
	fmt.Println(mathutil.Round(1.23))                   // int 1
	fmt.Println(mathutil.Round(1.56))                   // int 2
	fmt.Println(mathutil.RoundUp(1.1234567, 1))         //1.2
	fmt.Println(mathutil.RoundUp(1.1234567, 5))         //1.12346
	fmt.Println(mathutil.RoundDown(1.1234567, 1))       //1.1
	fmt.Println(mathutil.RoundDown(1.1234567, 5))       //1.12345
	fmt.Println(mathutil.Sum(1.123456789, 1.123456789)) //2.246913578
	fmt.Println(mathutil.Sub(1, 0.33))                  //0.67
	fmt.Println(mathutil.Mul(1.12, 1.23))               //1.3776
}
```

## idcard

身份证工具包

前置校验中国公民身份证号格式是否合法

很多公司有录入用户身份证号的需求，但是用户输入的身份证号是否合法还需要对接权威部门，调用的接口是需要费用的，对于一些不符合身份证号格式的输入我们可以自己先校验一下，可以节省一些不必要的调用接口费用。
> 注意：仅仅是校验身份证号格式是否合法，而不是校验身份证号是否合法，身份证号是否合法还得调用与权威部门对接
```go
package main

import (
	"fmt"
	"github.com/baosloan/utils/idcard"
)

func main() {
	var id = "33071919610920021X"
	if idcard.IsValid(id) {
		fmt.Println("身份证格式正确")
	} else {
		fmt.Println("身份证格式不正确")
	}
}
```
下面是身份证编码各个部分的含义：
```shell
//=============新的18位身份证号码各位的含义:=======================
//1-2位省、自治区、直辖市代码；11-65
//3-4位地级市、盟、自治州代码；
//5-6位县、县级市、区代码；
//7-14位出生年月日，比如19670401代表1967年4月1日；
//15-17位为顺序号，其中17位男为单数，女为双数；
//18位为校验码，0-9和X，由公式随机产生。
//举例：
//130503 19670401 0012这个身份证号的含义: 13为河北，05为邢台，03为桥西区，出生日期为1967年4月1日，顺序号为001，2为验证码
//===========15位身份证号码各位的含义:=======================
//1-2位省、自治区、直辖市代码；
//3-4位地级市、盟、自治州代码；
//5-6位县、县级市、区代码；
//7-12位出生年月日,比如670401代表1967年4月1日,这是和18位号码的第一个区别；
//13-15位为顺序号，其中15位男为单数，女为双数；
//与18位身份证号的第二个区别：没有最后一位的验证码。
//举例：
//130503 670401 001的含义; 13为河北，05为邢台，03为桥西区，出生日期为1967年4月1日，顺序号为001。

```