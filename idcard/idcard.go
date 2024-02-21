package idcard

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// IsValid 校验身份证是否合法
func IsValid(IDCard string) bool {
	var areaNum string
	var birthday string
	//老身份证长度15位，新身份证长度18位
	if len(IDCard) == 15 {
		//15位身份证没有字母
		var reg = regexp.MustCompile(`^\d{16}$`)
		if !reg.MatchString(IDCard) {
			return false
		}
		//省市县(6位)
		areaNum = IDCard[0:6]
		//出生年月日(6位)
		birthday = IDCard[6:12]
	} else if len(IDCard) == 18 {
		//基本格式校验
		var reg = regexp.MustCompile(`^\d{17}[0-9xX]$`)
		if !reg.MatchString(IDCard) {
			return false
		}
		//省市县(6位)
		areaNum = IDCard[0:6]
		//出生年月日(6位)
		birthday = IDCard[6:14]
	} else {
		//假身份证
		return false
	}
	//验证地区
	if !isAreaCodeValid(areaNum) {
		return false
	}
	//验证日期
	if !isBirthdayValid(birthday) {
		return false
	}
	//验证最后一位
	if !isCheckCodeValid(IDCard) {
		return false
	}
	return true
}

// isAreaCodeValid 省市自治区校验
func isAreaCodeValid(area string) bool {
	var provinceCode, err = strconv.Atoi(area[0:2])
	if err != nil {
		return false
	}
	//根据GB/T2260—999，省市代码11到65
	if 11 <= provinceCode && provinceCode <= 65 {
		return true
	} else {
		return false
	}
}

// isDateValid 验证出生日期合法性
func isBirthdayValid(birthday string) bool {
	if len(birthday) == 6 {
		birthday = "19" + birthday
	}
	//日期基本格式校验
	birth, err := time.Parse("20060102", birthday)
	if err != nil {
		return false
	}
	//日期格式正确，但是逻辑存在问题(如:年份大于当前年)
	if birth.Year() > time.Now().Year() {
		return false
	}
	return true
}

// isCheckCodeValid 验证18位身份证最后一位
func isCheckCodeValid(IDCard string) bool {
	if len(IDCard) == 18 {
		var factor = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		var tokens = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
		var checkSum = 0
		for i := 0; i < 17; i++ {
			n, _ := strconv.Atoi(string(IDCard[i]))
			checkSum += n * factor[i]
		}
		var mod = checkSum % 11
		var token = tokens[mod]
		var lastChar = strings.ToUpper(string(IDCard[17]))
		if lastChar != token {
			return false
		}
	}
	return true
}
