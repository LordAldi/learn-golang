package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		// handle the error here
// 		return
// 	}
// 	defer file.Close()
// 	// get the file size
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return
// 	}
// 	// read the file
// 	bs := make([]byte, stat.Size())
// 	_, err = file.Read(bs)
// 	if err != nil {
// 		return
// 	}
// 	str := string(bs)
// 	fmt.Println(str)
// }

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// bs, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	return
	// }
	// str := string(bs)
	// fmt.Println(str)
	// file, err := os.Create("test1.txt")
	// if err != nil {
	// 	// handle the error here
	// 	return
	// }
	// defer file.Close()
	// file.WriteString("test")

	// dir, err := os.Open(".")
	// if err != nil {
	// 	return
	// }
	// defer dir.Close()
	// fileInfos, err := dir.Readdir(-1)
	// if err != nil {
	// 	return
	// }
	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name())
	// }
	err := errors.New("error message")

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	fmt.Println(err)

}
