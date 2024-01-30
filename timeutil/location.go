package timeutil

import (
	"errors"
	"time"
)

const (
	LocationCairo        = "Africa/Cairo"                   //开罗
	LocationCasablanca   = "Africa/Casablanca"              //卡萨布兰卡
	LocationJohannesburg = "Africa/Johannesburg"            //约翰内斯堡
	LocationArgentina    = "America/Argentina/Buenos_Aires" //布宜诺斯艾利斯
	LocationChicago      = "America/Chicago"                //芝加哥
	LocationDenver       = "America/Denver"                 //丹佛
	LocationLosAngeles   = "America/Los_Angeles"            //洛杉矶
	LocationNewYork      = "America/New_York"               //纽约
	LocationBangkok      = "Asia/Bangkok"                   //曼谷
	LocationBeirut       = "Asia/Beirut"                    //贝鲁特
	LocationDubai        = "Asia/Dubai"                     //迪拜
	LocationHongKong     = "Asia/Hong_Kong"                 //香港
	LocationJerusalem    = "Asia/Jerusalem"                 //耶路撒冷
	LocationKolkata      = "Asia/Kolkata"                   //孟买、加尔各答
	LocationShanghai     = "Asia/Shanghai"                  //上海时间
	LocationTokyo        = "Asia/Tokyo"                     //上海
	LocationAmsterdam    = "Europe/Amsterdam"               //阿姆斯特丹
	LocationAthens       = "Europe/Athens"                  //雅典
	LocationBerlin       = "Europe/Berlin"                  //柏林
	LocationLisbon       = "Europe/Lisbon"                  //里斯本
	LocationLondon       = "Europe/London"                  //伦敦
	LocationMoscow       = "Europe/Moscow"                  // 莫斯科
	LocationAuckland     = "Pacific/Auckland"               //奥克兰
	LocationHonolulu     = "Pacific/Honolulu"               //檀香山
)

// ParseInLocation 将时间字符串转换为对应时区的time.Time
func ParseInLocation(timeStr string, loc string) (t time.Time, err error) {
	if len(timeStr) != dateOnlyLength && len(timeStr) != DateTimeLength {
		return t, errors.New("timeStr format is invalid")
	}
	//上海时区
	location, err := time.LoadLocation(loc)
	if err != nil {
		return
	}
	//解析时间字符串并指定时区
	if len(timeStr) == dateOnlyLength {
		t, err = time.ParseInLocation(time.DateOnly, timeStr, location)
	} else if len(timeStr) == DateTimeLength {
		t, err = time.ParseInLocation(time.DateTime, timeStr, location)
	}
	return
}

// ParseInLocationShanghai 将日期字符串转换为上海时区time.Time
func ParseInLocationShanghai(timeStr string) (time.Time, error) {
	return ParseInLocation(timeStr, LocationShanghai)
}

// ParseInLocationNewYork 将日期字符串转换为纽约时区time.Time
func ParseInLocationNewYork(timeStr string) (time.Time, error) {
	return ParseInLocation(timeStr, LocationNewYork)
}
