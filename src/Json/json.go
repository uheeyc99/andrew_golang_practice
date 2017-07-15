package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)


type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float32
	Sales []Sale
}

type Sale struct {
	Where string
	Howmany int
}

func enc()  {
	aa:=Sale{
		Where:"广州",
		Howmany:5000,

	}
	book1 := Book{
		Title:"Aiden最棒",
		Authors:[]string{"Jessica","Aiden"},
		Publisher:"aiden.cn",
		IsPublished:true,
		Price:99.99,
		Sales:[]Sale{aa,aa,aa},
	}

	book1.Sales[1].Where = "福州"
	book1.Sales[2].Where = "厦门"

	b,e:=json.Marshal(book1)
	if e !=nil {
		fmt.Println(e)
	}

	f,e:=os.Create("uploads/jsonTest.txt")
	if e !=nil {
		fmt.Println(e)
	}
	defer f.Close()
	f.Write(b)

	fmt.Println(string(b))
	fmt.Println(book1)
}

func dec()  {

	f,e:=os.Open("uploads/jsonTest.txt")
	if e !=nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer f.Close()

	reader:=bufio.NewReader(f)
	line,_,_:=reader.ReadLine()

	//解析到结构体里
	var book2 Book
	json.Unmarshal(line,&book2)  //取到结构体中
	fmt.Println("打印结构体:")
	fmt.Println(book2)

	//解析到interface
	var i interface{}
	json.Unmarshal(line,&i)
	m :=i.(map[string]interface{})  //取到interface
	fmt.Println("打印interface:")
	fmt.Println(i)
	fmt.Println(m)

	//for k,v:=range m{
	//	switch v2:=v.(type) {
	//	case string:
	//		fmt.Println(k,"is string:",v2)
	//	case int:
	//		fmt.Println(k,"is int:",v2)
	//	case bool:
	//		fmt.Println(k,"is bool:",v2)
	//	case float32:
	//		fmt.Println(k,"is float32:",v2)
	//
	//	case []interface{}:
	//		fmt.Println(k,"is array:",v2)
	//	default:
	//		fmt.Println(k,"is another type not handle yet")
	//
	//	}
	//
	//}


}


func test(){
	type Server struct {
		Name string
		IP string
	}
	type ServerList struct {
		ServerList []Server
	}

	var s ServerList

	s.ServerList = append(s.ServerList,Server{
		Name:"上海VPN",
		IP:"202.106.88.90",
	})
	s.ServerList = append(s.ServerList,Server{
		Name:"杭州VPN",
		IP:"202.106.88.91",
	})
	fmt.Println(s)
	b,e:=json.Marshal(s)
	//{"ServerList":[{"Name":"上海VPN","IP":"202.106.88.90"},{"Name":"杭州VPN","IP":"202.106.88.91"}]}

	//b,e:=json.Marshal(s.ServerList)
	//[{"Name":"上海VPN","IP":"202.106.88.90"},{"Name":"杭州VPN","IP":"202.106.88.91"}]
	if e!=nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))

	var i  interface{}
	json.Unmarshal(b,&i)
	fmt.Println(i)
	fmt.Println(i.(map[string]interface{}))




}


func main()  {

	dec()
	fmt.Println("########################")
	test()

}
