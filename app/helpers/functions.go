package helpers

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 获取md5明文字串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	x := h.Sum(nil)
	y := fmt.Sprintf("%x", x)
	return y
}

func Test() {
	fmt.Print(555)
}

// 处理union id
func DealUnionId(union_id string) string {

	if strings.Contains(union_id, "weixin_") || strings.Contains(union_id, "wechat_") {
		return union_id[7:]
	} else if strings.Contains(union_id, "opensdk_") {
		return union_id[8:]
	}
	return union_id
}

// 获取订单ID
func GetOrderId(prefix string) string {
	return prefix + time.Now().Format("060102150405") + strconv.Itoa(RandInt(1000, 9999))
}

// 生成订单号
func generateOrderNo() (orderNo string) {
	yearCode := "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	yearCodeSlice := []rune(yearCode)

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()

	y := yearCode[(year - 2015):(year - 2014)]
	m := strings.ToUpper(fmt.Sprintf("%x", int(month)))
	d := fmt.Sprintf("%02d", day)

	unixTime := strconv.FormatInt(now.Unix(), 10)
	a := unixTime[(len(unixTime) - 5):]

	microTime := strconv.FormatInt(now.UnixNano(), 10)
	b := microTime[12:18]

	tmp := make([]rune, 2)
	for i := range tmp {
		tmp[i] = append(yearCodeSlice, yearCodeSlice...)[rand.Intn(len(yearCodeSlice))]
	}
	c := string(tmp)

	orderNo = y + m + d + a + b + c

	return
}

// 取随机数
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// 获取当前执行文件路径
func Getwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

// 获取storage目录
func StoragePath() string {
	path := Getwd()
	return path + "/storage/" //将\替换成/
}

// 判断值是否在数组中，字符串操作
func InArrayStr(value string, arr []string) bool {
	if len(arr) <= 0 {
		return false
	}

	for _, v := range arr {
		if value == v {
			return true
		}
	}

	return false
}

// 判断值是否在数组中，整数操作
func InArrayInt(value int, arr []int) bool {
	if len(arr) <= 0 {
		return false
	}

	for _, v := range arr {
		if value == v {
			return true
		}
	}

	return false
}

// 校验传递的int参数
func CheckParamInt(key string, ret *uint) bool {
	if _, ok := RequestParam[key]; !ok {
		return false
	}

	i, err := strconv.Atoi(RequestParam[key].(string))
	*ret = uint(i)
	if err != nil || *ret <= 0 {
		return false
	}

	return true
}

// 校验传递的float32参数
func CheckParamFloat32(key string, ret *float32) bool {
	if _, ok := RequestParam[key]; !ok {
		return false
	}

	i, err := strconv.ParseFloat(RequestParam[key].(string), 32)
	*ret = float32(i)
	if err != nil || *ret <= 0 {
		return false
	}

	return true
}

// map排序
func SortParam(p map[string]interface{}) string {

	var tmp []string
	for k, v := range p {
		switch v.(type) {
		case string:
			tmp = append(tmp, k+"="+v.(string))
			break
		case int:
			tmp = append(tmp, k+"="+strconv.Itoa(v.(int)))
			break
		default:
			break
		}

	}
	sort.Strings(tmp)

	return strings.Join(tmp, "&")
}

// HMAC_SHA256加密
func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
