package day12

import (
	"encoding/json"
	"strconv"
)

func searchNums(iface interface{}, skipRed bool) int {
	var sum int
	switch v := iface.(type) {
	case float64:
		sum += int(v)
	case []interface{}:
		for _, i := range v {
			sum += searchNums(i, skipRed)
		}
	case map[string]interface{}:
		var red bool
		if skipRed {
			for _, value := range v {
				if s, ok := value.(string); ok {
					if s == "red" {
						red = true
					}
				}
			}
		}
		if !red {
			for _, value := range v {
				sum += searchNums(value, skipRed)
			}
		}
	}
	return sum
}

func Part1(input string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(searchNums(data, false)), nil
}

func Part2(input string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(searchNums(data, true)), nil
}
