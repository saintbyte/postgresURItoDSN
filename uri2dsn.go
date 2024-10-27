package postgresURItoDSN

import (
	"errors"
	"net/url"
	"strings"
)

func mapToString(m map[string]string) string {
	var pairs []string
	for key, value := range m {
		pairs = append(pairs, key+"="+value)
	}
	return strings.Join(pairs, " ")
}

func UriToDSN(URI string) (string, error) {
	//postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	//"user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	if len(URI) < 14 {
		return "", errors.New("wrong uri'")
	}
	if URI[0:13] != "postgresql://" {
		return "", errors.New("wrong protocol, support only 'postgresql://'")
	}
	UriObj, err := url.Parse(URI)
	if err != nil {
		return "", err
	}

	var dsnMap = map[string]string{}
	if UriObj.User.Username() != "" {
		dsnMap["user"] = UriObj.User.Username()
	}
	password, has := UriObj.User.Password()
	if has {
		dsnMap["password"] = password
	}
	path := UriObj.Path[1:]
	dsnMap["dbname"] = path
	if UriObj.Hostname() == "" {
		return "", errors.New("Empty host")
	}
	dsnMap["host"] = UriObj.Hostname()
	port := UriObj.Port()
	if port != "" {
		dsnMap["port"] = port
	}
	qs := UriObj.Query()
	for k, _ := range UriObj.Query() {
		dsnMap[k] = qs.Get(k)
	}
	result := mapToString(dsnMap)
	return result, nil
}
