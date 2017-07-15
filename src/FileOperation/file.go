package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"io"
)


/*
	列出目录下文件和目录
*/
func list_dir()  {
	dirInfoArr,err:=ioutil.ReadDir("./")
	if err !=nil {
		fmt.Println(err)
	}
	for i,fileInfo :=range dirInfoArr{
		filename:=fileInfo.Name()
		fmt.Println(i,"name:"+ filename,fileInfo.IsDir())

	}
}

func read_file(filename string){
	file,err:=os.Open(filename)
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	br:=bufio.NewReader(file)
	var totalsize int = 0
	for{
		line,isPrefix,err:=br.ReadLine()
		if err!=nil {

			fmt.Println("到文件末尾了",err)
			break
		}

		if isPrefix{
			fmt.Println("" +
				"a too long line,seems unexpected.")
		}
		//fmt.Println(string(line),len(line))
		fmt.Println(len(line))
		totalsize = totalsize +len(line)
	}
	fmt.Println(totalsize , "bytes")
	return
}

func write_file(filename string)  {
	f,err:=os.Create(filename)
	if err !=nil {
		fmt.Println(err)
	}
	defer f.Close()
	s1 :=[]byte("Hello")
	f.Write(s1)
	s2:="Morning"
	f.WriteString(s2)


}


func copy_file(filename string){
	fs,err:=os.Open(filename)
	if err !=nil {
		return
	}
	defer fs.Close()
	fd,err:=os.Create("./uploads/123.rmvb")
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer fd.Close()
	io.Copy(fd,fs)

}


func main()  {
	//copy_file("uploads/dqdg01.rmvb")
	//read_file("uploads/dqdg01.rmvb")
	write_file("uploads/dqdg01.txt")

}
