package main

import "fmt"

func main() {
	//Input:
	//	[]string{"go", "java", "go", "python", "go", "java"}
	//Output:
	//	map[string]int{"go": 3, "java": 2, "python": 1}
	languages := []string{"go", "java", "go", "python", "go", "java"}
	addMap := map[string]int{}
	goL := 0
	javaL := 0
	pythonL := 0

	for i := 0; i < len(languages); i++ {
		if languages[i] == "go" {
			goL++
			addMap[languages[i]] = goL
		} else if languages[i] == "java" {
			javaL++
			addMap[languages[i]] = javaL
		} else if languages[i] == "python" {
			pythonL++
			addMap[languages[i]] = pythonL
		}

	}
	for k, v := range addMap {
		fmt.Println(k, v)
	}

	//fmt.Printf("%d", addMap["go"])
	//task 5 is done

}
