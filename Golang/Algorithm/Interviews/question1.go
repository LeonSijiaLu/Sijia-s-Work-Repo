// https://blog.csdn.net/cyan1956/article/details/104804609

package main
import (
	"fmt"
	"strconv"
	"strings"
    "bufio"
    "os"
)

func main(){
	inputReader := bufio.NewReader(os.Stdin)
	in, _ := inputReader.ReadString('\n')
	in = strings.TrimSpace(in)
	num, _ := strconv.ParseFloat(in, 64)

	var words []string
	
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		words = append(words, scanner.Text())
		num = num - 1
		if num == 0 {
			break
		}
	}
	
	flag := 0
	tmp := ""

	for _, word := range words{
		for i, letter := range word{
			if i == 0{
				tmp = string(letter)
			}else{
				switch flag{ // helllo
				case 0:
					tmp = tmp + string(letter)
					if tmp[i] == tmp[i - 1]{
						flag = 1
					}
				case 1:
					if string(word[i]) != string(word[i-1]){ // AABBCC
						tmp = tmp + string(letter)
						flag = 2
					}
				case 2:
					if string(word[i]) != string(word[i-1]){
						tmp = tmp + string(letter)
						flag = 0
					}
				}
			}
		}
		fmt.Println(tmp)
	}
}