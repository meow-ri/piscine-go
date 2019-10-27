package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	lengthOfArg := 0
	for i := range os.Args {
		lengthOfArg = i
	}
	if lengthOfArg < 1 {
		fmt.Println("File name missing")
	} else if lengthOfArg > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, er := os.Open(os.Args[1])
		if er != nil {
			fmt.Println(er.Error())
		} else {
			text, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				defer file.Close()
				fmt.Println(string(text))
			}
		}
	}
}
