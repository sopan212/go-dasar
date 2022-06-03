package main

import "fmt"

func main() {
	names := []string{
		"Sopan",
		"Mukti",
		"Aliyansyah",
	}

	for i, val := range names {
		fmt.Println(i, "=", val)
	}

	//for using map
	Books := map[string]string{
		"Naruto":   "Naruto",
		"One Pice": "lufi",
		"Gazebo":   "nobo",
	}

	for key, val := range Books {
		fmt.Println(key, "-->", val)
	}
}
