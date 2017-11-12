package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println("Papa ist der größte")

	i := 0
	for ; i < 1000; {
		// furzkopf fred
		//    for {
		i = i + 1
		if i >= 990 {
			fmt.Println(i,text)

		}
		//fmt.Println(i, "10§ zu bezahlen:   du duumme nuss")
	}
	fmt.Println("wir sind die stärksten")
	fmt.Println("Papa ist der größte")
	fmt.Println("Papa ist der größte")
	fmt.Println("Papa ist der größte")
}
