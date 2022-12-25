package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string    //`json:"username"`
	Nickname string    //`json:"nickname"`
	Sex      uint8     //`json:"sex"`
	Birthday time.Time //`json:"birthday"`
}

// 问题产生原因：由于是在json包中使用序列化，如果结构体内元素首字母没有大写则无法访问包，无法引用数据，导致最后实际的输出结果为“{}”
// 解决方法1：将结构体首字母修改为大写
// 解决方法2：将结构体加上json的tag标签
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
