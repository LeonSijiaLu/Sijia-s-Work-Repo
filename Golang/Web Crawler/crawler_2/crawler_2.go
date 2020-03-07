package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, end int
	fmt.Println("Please enter start page: ")
	fmt.Scan(&start)
	fmt.Println("Please enter end page: ")
	fmt.Scan(&end)
	if end < start || end <= 0 || start <= 0 {
		fmt.Println("Re-enter start or end page")
	} else {
		crawler(start, end)
	}
}

func httpGet(url string) (result string) { // 1. 通过URL, 爬取数据
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result = result + string(buf[:n])
	}
	return result
}

func remove_previous_html() {
	cur_dir, _ := os.Getwd()          // Get current directory
	dirRead, _ := os.Open(cur_dir)    // Open current directory
	dirFiles, _ := dirRead.Readdir(0) // Read all files in current directory
	for index := range dirFiles {
		file_name := dirFiles[index].Name()
		if strings.HasPrefix(file_name, "Page") {
			os.Remove(file_name)
		}
	}
}

func crawler_thread(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa(i*50)
	result := httpGet(url)

	// Create an HTML file
	// 把返回的HTML数据存入HTML
	file, _ := os.Create("Page" + strconv.Itoa(i) + ".html")
	file.WriteString(result)
	file.Close()
	page <- i
}

func crawler(start int, end int) {
	remove_previous_html()
	page := make(chan int, end-start)

	for i := start; i <= end; i++ {
		go crawler_thread(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Println("Successfully fetched page ", <-page, " data")
	}
}
