package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/projectdiscovery/gologger"
	"os"
	"regexp"
)

func main() {
	inputRegex := flag.String("r", "", "Your regex [e.g: ^\\w+]")
	groupNumber := flag.Int("g", 0, "Group number")
	help := flag.Bool("h", false, "Help, show usage parameters")
	flag.Parse()

	if *inputRegex == "" || *help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stdin.Stat()

	myRegex, err := regexp.Compile(*inputRegex)
	checkError(err)

	input := readInput()
	for _, v := range input {
		for _, value := range myRegex.FindAllStringSubmatch(v, -1) {
			if len(value) > *groupNumber {
				fmt.Println(value[*groupNumber])
			}
		}
	}
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var result []string
	for scanner.Scan() {
		txt := scanner.Text()
		result = append(result, txt)
	}
	return result
}

func checkError(e error) bool {
	if e != nil {
		gologger.Error().Msg(e.Error())
		return true
	}
	return false
}
