package main

import(
	"bufio"
	"os"
	"strings"
	"fmt"
)

// Counts and gives address for labels
var globalCounter int = 0

// Counts and gives address for variables
var globalAddress int = 16

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Wrong command. Write: go run *.go <input_file> <output_file>")
		return
	}
	fileRead, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fileRead.Close()

	scanner := bufio.NewScanner(fileRead)
	scanner.Split(bufio.ScanLines)
	
	buf := ""
	for scanner.Scan() {
		if text, isEmpty := trimmer(scanner.Text()); !isEmpty{
			buf += text + "\n"
			globalCounter++
		}
	}
	sliceStr := strings.Fields(buf)
	buf = ""
	for _, text := range sliceStr {
		if strings.Contains(text, "@"){
			buf += translateAInstruct(text) + "\n"
			continue
		}
		buf += translateCInstruct(text) + "\n"
	}
	fileWrite, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(fileWrite)
	writer.WriteString(buf)
	writer.Flush()
	fmt.Println("File saved as " + os.Args[2])
}