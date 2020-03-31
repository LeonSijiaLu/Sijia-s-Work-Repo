package main

import(
	"fmt"
)

func CaesarCipherEncryptor(str string, key int) string {
	for _,value := range str{
		asciiNum := int(value)
		if(67 <= asciiNum <= 90){
			
		}
	}
	return str
}

func main(){
	fmt.Println(CaesarCipherEncryptor("ABCXYZabcxyz", 3)) 
	//A-Z 65-90
	//a-z 97-122
}