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