// https://blog.csdn.net/Cyan1956/article/details/104804206

// Sliding Window Question

package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"bufio"
)

func main(){
	inputReader := bufio.NewReader(os.Stdin)
	line1, _ := inputReader.ReadString('\n')
	line2, _ := inputReader.ReadString('\n')

	num_buildings, _ := strconv.Atoi(strings.Fields(line1)[0])
	max_distance, _ := strconv.Atoi(strings.Fields(line1)[1])
	locations := strings.Fields(line2)

	var pos []int
	var cases int

	for i, j := 0,0; i < num_buildings; i ++{
		loc, _ := strconv.Atoi(locations[i])
		pos = append(pos, loc)
		for i >=3 && pos[i] - pos[j] >= max_distance{
			j ++
		}
		cases = cases + (i - j) * (i - j - 1) / 2
	}
	fmt.Println(cases%99997867)
}