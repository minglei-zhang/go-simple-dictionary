package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type result struct {
	ResArray []string `xml:"string"`
}

func main() {
	fmt.Print("\n---------------  英汉词典词典 ------------------------\n")
	fmt.Print(`请输入查询的字段：`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		searchRun(scanner.Text())
	}
}

func searchRun(scanner string) {
	// 捕获错误
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s", r)
		}
	}()

	result := result{}

	// 请求api接口查询翻译结果
	fmt.Print("查询中...\n")
	res, httpErr := http.Get("http://ws.webxml.com.cn//WebServices/TranslatorWebService.asmx/getEnCnTwoWayTranslator?Word=" + scanner)
	if httpErr != nil {
		panic("\n----------------【错误】请重新查询！----------------\n\n请输入查询字段：")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("\n----------------【错误】请重新查询！----------------\n\n请输入查询字段：")
	}
	defer res.Body.Close()

	// 解析xml
	errs := xml.Unmarshal([]byte(string(body)), &result)
	if errs != nil {
		panic("\n----------------【错误】请重新查询！----------------\n\n请输入查询字段：")
	}

	fmt.Print("\n查询结果如下：\n")

	// 遍历结果输出
	for _, value := range result.ResArray {
		fmt.Print(value, "\n")
	}

	fmt.Print("\n\n请输入查询字段：")
}
