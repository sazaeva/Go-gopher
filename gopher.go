package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("`^........^`")
	fmt.Println("............")
	fmt.Println("..(o)..(o)..")
	fmt.Println(" ..   &   ..")
	fmt.Println("..   UU   ..")
	fmt.Println("  ........  ")
	fmt.Println(" ......... ")
	fmt.Println("8.........8")
	fmt.Println(" .........")
	fmt.Println(" .........")
	fmt.Println(" .........")
	fmt.Println("o.........o")
	t := time.Date(2022, time.February, 02, 16, 20, 0, 0, time.UTC)
	fmt.Printf("Gopher is created  at %s\n", t.Local())
}
