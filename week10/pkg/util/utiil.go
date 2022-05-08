package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"parent-api-go/global"
	"parent-api-go/pkg/context"
	"path/filepath"
	"strings"
	"time"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func CurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// fmt.Println(fmt.Sprintf("%+v", req))

func GenRandValue() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return SafeBase64Replace(base64.StdEncoding.EncodeToString(b)), nil
}

func SafeBase64Replace(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimRight(str, "="), "/", "-"), "+", "_")
}

func UniqueInt(is []int) []int {
	i := []int{}
	unq := map[int]bool{}
	for _, v := range is {
		if _, ok := unq[v]; ok {
		} else {
			unq[v] = true
			i = append(i, v)
		}
	}
	return i
}

func UniqueUint32(is []uint32) []uint32 {
	i := []uint32{}
	unq := map[uint32]bool{}
	for _, v := range is {
		if _, ok := unq[v]; ok {
		} else {
			unq[v] = true
			i = append(i, v)
		}
	}
	return i
}

func toNormalType(i interface{}) {

}

func ArrayMapStringInterface(i []interface{}) []map[string]interface{} {
	msi := []map[string]interface{}{}
	for _, v := range i {
		switch ii := v.(type) {
		case map[interface{}]interface{}:
			mi := map[string]interface{}{}
			for k, vv := range ii {
				mi[fmt.Sprintf("%s", k)] = vv
			}
			if len(mi) > 0 {
				msi = append(msi, mi)
			}

		case map[string]interface{}:
			msi = append(msi, ii)
		}
	}
	return msi
}

func InArrayInts(i int, ii []int) bool {
	for _, m := range ii {
		if m == i {
			return true
		}
	}

	return false
}
func InArrayStrings(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func TimeLocalFmt(t string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
}

// 获取当前日期(yyyy-mm-dd H:i:s)
func NowDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 格式化time类型为(yyyy-mm-dd H:i:s)
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 时间戳（毫秒）转 日期(yyyy-mm-dd H:i:s)
func UnixToDate(ms int64) string {
	tm := time.Unix(ms, 0)
	return FormatDate(tm)
}

func GetAppContext(c *gin.Context) (*context.AppContext, error) {
	if ac, ok := c.Get(global.AppContext); !ok {
		logrus.Error("not found app context")
		return nil, errors.New("not found app context")
	} else if act, ok := ac.(*context.AppContext); !ok {
		return nil, errors.New("type error")
	} else {
		return act, nil
	}
}

func IsAndroid(ua, _aKey string) bool {
	if strings.Contains(ua, "JingyupeiyouAndroidApp") {
		return true
	}

	if strings.Contains(strings.ToLower(_aKey), "android") {
		return true
	}
	return false
}

func IsIOS(ua, _aKey string) bool {
	if strings.Contains(ua, "jingyupeiyouIOSApp") {
		return true
	}

	key := strings.ToLower(_aKey)
	iosOld := "com.rouchi.jingyupeiyou.parents"

	if strings.Contains(key, "ios") || strings.Contains(key, iosOld) {
		return true
	}

	return false
}

func GetMobileOS(ua, _aKey string) uint8 {
	mobileOS := 0
	if IsAndroid(ua, _aKey) {
		mobileOS = 2
	} else if IsIOS(ua, _aKey) {
		mobileOS = 1
	}

	return uint8(mobileOS)
}

/**
 * @Description: string类型切片分割 类似php array_chunk()
 * @param s
 * @param size
 * @return ss
 */
func SliceChunkForString(s []string, size int) (ss [][]string) {
	j := 0
	sLen := len(s)
	if size <= 0 {
		size = 50
	}

	for i := 0; i <= sLen-1; i += size {
		j = i + size
		if j > sLen {
			j = sLen
		}
		ss = append(ss, s[i:j])
	}
	return
}

//格式护数值    1,234,567,898.55
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}
