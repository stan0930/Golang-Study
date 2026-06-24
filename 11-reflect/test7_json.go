//结构体标签的应用  1.json编解码 2.orm映射关系
package main

import (
	"encoding/json"
	"fmt"
)



type Movie struct {//结构体标签
	Title string 	`json:"title"`
	Year int	 	`json:"year"`
	Price int 	 	`json:"rmb"`
	Actor []string  `json:"actors"`
}

func main()  {
	movie := Movie{"喜剧之王",2000,10,[]string{"xingye","zhangbozhi"}}


	//编码的过程 结构体---》json
	jsonStr,err := json.Marshal(movie)//json库里的Marshal，返回两个字段
	if err != nil{ //说明导出失败
		fmt.Println("json marshal error",err)
		return
	}

	fmt.Printf("jsonStr = %s\n",jsonStr)

	//解码的过程 jsonstr---》结构体
	//jsonStr =  {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}
	myMovie := Movie{}
	json.Unmarshal(jsonStr,&myMovie)
	if err != nil {
		fmt.Println("json unmarshal error",err)
		return
	}

	fmt.Println("%v\n",myMovie)
}