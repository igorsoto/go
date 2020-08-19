package main

import "fmt"

// map is reference type
func set(m map[string]string, k string, v string) {
	m[k] = v
}

func print(m map[string]string) {
	for k, v := range m {
		fmt.Println("k", k, "v", v)
	}
}

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	print(colors)

	set(colors, "white", "#ffffff")

	print(colors)
}
