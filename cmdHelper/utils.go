package cmdHelper

import (
	"bufio"
	"fmt"
	"os"
)

func ReverseSlice(s []interface{}) []interface{} {
	reversed := make([]interface{}, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = s[j], s[i]
	}
	return reversed
}

func UpdateMap(key string, value string, currentMap map[string][]interface{}) map[string][]interface{} {
	if val, ok := currentMap[key]; ok {
		currentMap[key] = append(val, value)
	} else {
		currentMap[key] = []interface{}{value}
	}
	return currentMap
}

func ReadJson() (string, error) {
	fmt.Println("Please Input the setting you want to set (type `save` to save setting): ")

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	jsonString := ""
	for _, line := range lines {
		jsonString += line
	}

	return jsonString, nil
}
