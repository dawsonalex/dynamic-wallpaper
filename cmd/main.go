package main

import (
	"fmt"
	"github.com/reujab/wallpaper"
)

func main() {
	path, err := wallpaper.Get()
	if err != nil {
		fmt.Errorf("error getting wallpaper: %v", err)
	}

	fmt.Println(path)
}
