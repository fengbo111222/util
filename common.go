/*
@Author : FengB
@Time : 2022/1/25 11:16
@Remark :
*/
package util

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"login/commonMap"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 时间转换 设置时区 东巴区
var cstZone = time.FixedZone("CST", 8*3600)
func init() {
	time.Local = cstZone
}
func GetTimeStringByDate(date string) string {
	location, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", date, cstZone)
ErrInfo("GetTimeStringByDate",err)
	return location.Format("2006-01-02 15:04:05")
}
func GetUUId() string {

	uid := uuid.NewV1()
	return strings.Replace(uid.String(), "-", "", -1)

}
// 生成token
func GenerateUserToken() (*commonMap.Token, error) {
	tokenByte, err := GenerateUserTokenStruct()
	if err != nil {
		return nil, err
	}
	tokenInfo := commonMap.Token{}
	err = json.Unmarshal(tokenByte, &tokenInfo)
	if err != nil {
		return nil, err
	}
	tokenStr, err := GenerateToken(tokenByte)
	if err != nil {
		return nil, err
	}
	tokenInfo.Token = tokenStr
	return &tokenInfo, nil
}
func GetNowExpiretime() string {
	ExpireDay, _ := strconv.Atoi(os.Getenv("Token_Expire_Day"))
	expireTime := GetNowTimeStamp() + ExpireDay*3600*24
	timeStr, _ := GetDateFormat(int64(expireTime), "")
	return timeStr
}

// 获取当前时间戳到秒
func GetNowTimeStamp() int {
	t := time.Now().In(cstZone)
	nowTime := strconv.FormatInt(t.UTC().UnixNano(), 10)
	nowTime = nowTime[:10]
	timeStamp, _ := strconv.Atoi(nowTime)
	return timeStamp
}

func GenerateUserTokenStruct() (tokenByte []byte, err error) {
	// 获取有效时间

	tokenStruct := commonMap.Token{
		Token:      GetUUId(),
		Expiretime: GetNowExpiretime(),
	}

	tokenByte, err = json.Marshal(tokenStruct)
	if err != nil {
		return tokenByte, err
	}
	return tokenByte, nil
}

// 生成token
func GenerateToken(tokenByte []byte) (token string, err error) {
	token, err = Encrypt(tokenByte)
	if err != nil {
		return token, err
	}
	return token, nil
}

func GetDateFormat(timeStamp int64, formatString string) (date string, err error) {
	secondTimeStamp := strconv.FormatInt(timeStamp, 10)
	i, err := strconv.ParseInt(secondTimeStamp[:10], 10, 64)
	if err != nil {
		return date, err
	}
	t := time.Unix(i, 0).In(cstZone)
	switch formatString {
	case "YYYYMMDD":
		return t.Format("20060102"), nil
	}
	return t.Format("2006-01-02 15:04:05"), nil
}
func ErrInfo(funcName string,err error)  {
	if err!=nil  {
		log.Panicln(fmt.Sprintf("%v err %v:",funcName,err))
	}
}
func ErrInfoPanic(funcName string,err error)  {
	if err!=nil  {
		panic(fmt.Sprintf("%v err %v:",funcName,err))
	}
}
//func main() {
//	call(map[string]interface{}{"a":a},"a",3)
//}
//func a(b int)  {
//	fmt.Println(b)
//}
func Call(funMap map[string]interface{},name string,parames...interface{})  {
	f:=reflect.ValueOf(funMap[name])
	if len(parames)!=f.Type().NumIn() {
		log.Panicln("参数不对")
	}
	in:=make([]reflect.Value,len(parames))
	for i, parame := range parames {
		in[i]=reflect.ValueOf(parame)
	}
	f.Call(in)
}