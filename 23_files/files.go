package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("File name:",fileInfo.Name())
	// fmt.Println("Is folder:",fileInfo.IsDir())
	// fmt.Println("File size:",fileInfo.Size())
	// fmt.Println("File permission:",fileInfo.Mode())
	// fmt.Println("File modified at:",fileInfo.ModTime())

	// ------------read file------------
	// ------> Method 1
	// f, err := os.Open("example.txt")
	// if err!=nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// buf := make([]byte,12)
	// d,err:=f.Read(buf)
	// if err!=nil {
	// 	panic(err)
	// }

	// for i:=0;i<len(buf);i++ {
	// 	fmt.Println("data",d,string(buf[i]))
	// }

	// ------> Method 2
	// data,err := os.ReadFile("example.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// ----------read folders----------
	// dir,err:=os.Open("../")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo,err:=dir.ReadDir(-1)
	// for _,fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// ----------create a file----------
	// f,err:=os.Create("example2.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// // f.WriteString("hi go")
	// // f.WriteString("nice language")

	// bytes := []byte("Hello Golang")
	// f.Write(bytes)
	// f.Write([]byte("Its a very nice language..."))

	// ***read and write to another file (streaming fashion) ***//
	// sourceFile, err := os.Open("example.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("Example2.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e:=writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("written to new file successfully")

	// ----------delete file----------
	// sourceFile, err := os.Open("Example2.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	err:=os.Remove("Example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted successfully")
}