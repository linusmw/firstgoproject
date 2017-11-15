package main

import "fmt"
func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter text: ")
	//text, _ := reader.ReadString('\n')

	zahl := 25200
	i := 0
	for ; i < zahl; {
		// furzkopf fred
		//    for {
		i = i + 1
		if (zahl % i) == 0 {
			fmt.Println(i)

		}
		//fmt.Println(i, "10§ zu bezahlen:   du duumme nuss")
	}
}