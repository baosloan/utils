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