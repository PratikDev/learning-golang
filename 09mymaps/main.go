package main

import "fmt"

func main() {
	languages := map[string]string{
		"js": "javascript",
		"rb": "ruby-n-rails",
		"py": "pythong",
	}
	delete(languages, "rb")
	fmt.Println("Langues are:", languages)

	// loop over
	for key, value := range languages {
		fmt.Println(key, ":", value)
	}
}
