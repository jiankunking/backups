package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name         string
	Age          int
	BlogArticles map[string]interface{}
}

type BlogArticle struct {
	Detail string
	Author string
	Urls   map[string]string
}

type PersonCorrect struct {
	Name         string
	Age          int
	BlogArticles map[string]BlogArticle
}

func main() {

	blogArticle1 := new(BlogArticle)
	blogArticle1.Author = "jiankunking"
	blogArticle1.Detail = "csdn blog"

	url1 := make(map[string]string)
	url1["1"] = "http://blog.csdn.net/jiankunking/article/details/52143504"
	url1["2"] = "http://blog.csdn.net/jiankunking/article/details/52673302"
	url1["3"] = "http://blog.csdn.net/jiankunking/article/details/45479105"
	blogArticle1.Urls = url1

	blogArticle2 := new(BlogArticle)
	blogArticle2.Author = "JIANKUNKING"
	blogArticle2.Detail = "CSDN BLOG"

	url2 := make(map[string]string)
	url2["1"] = "http://blog.csdn.net/jiankunking/article/details/52684722"
	url2["2"] = "http://blog.csdn.net/jiankunking/article/details/78808978"
	blogArticle2.Urls = url2

	blogArticles := make(map[string]interface{})
	blogArticles["one"] = blogArticle1
	blogArticles["two"] = blogArticle2

	person := new(Person)
	person.Age = 12
	person.Name = "jiankunking"
	person.BlogArticles = blogArticles

	PrintJson(person)

	// unmarshal to struct
	var str = `{
  "Name": "jiankunking",
  "Age": 12,
  "BlogArticles": {
    "one": {
      "Detail": "csdn blog",
      "Author": "jiankunking",
      "Urls": {
        "1": "http://blog.csdn.net/jiankunking/article/details/52143504",
        "2": "http://blog.csdn.net/jiankunking/article/details/52673302",
        "3": "http://blog.csdn.net/jiankunking/article/details/45479105"
      }
    },
    "two": {
      "Detail": "CSDN BLOG",
      "Author": "JIANKUNKING",
      "Urls": {
        "1": "http://blog.csdn.net/jiankunking/article/details/52684722",
        "2": "http://blog.csdn.net/jiankunking/article/details/78808978"
      }
    }
  }
}`
	json.Unmarshal([]byte(str), &person)
	fmt.Println(typeof(person.BlogArticles["one"]))
	// 错误做法
	// blogArticle3 := person.BlogArticles["one"].(BlogArticle)
	// PrintJson(blogArticle3)

	//正确做法
	var personCorrect PersonCorrect
	json.Unmarshal([]byte(str), &personCorrect)
	fmt.Println(typeof(personCorrect.BlogArticles["one"]))
	// PrintJson(personCorrect.BlogArticles["one"])
}
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func PrintJson(obj interface{}) {
	jsons, errs := json.Marshal(obj)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons))
}

//输出结果如下：
// {"Name":"jiankunking","Age":12,"BlogArticles":{"one":{"Detail":"csdn blog","Author":"jiankunking","Urls":{"1":"http://blog.csdn.net/jiankunking/article/details/52143504","2":"http://blog.csdn.net/jiankunking/article/details/52673302","3":"http://blog.csdn.net/jiankunking/article/details/45479105"}},"two":{"Detail":"CSDN BLOG","Author":"JIANKUNKING","Urls":{"1":"http://blog.csdn.net/jiankunking/article/details/52684722","2":"http://blog.csdn.net/jiankunking/article/details/78808978"}}}}
// map[string]interface {}
// main.BlogArticle
