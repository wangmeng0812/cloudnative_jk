package main

import "fmt"

func main() {
	a := []string{"I","am","stupid","and","weak"}
	fmt.Println(a)
	for i := range a {
		switch a[i]{
		case "stupid":
			a[i]="smart"
		case "weak":
			a[i]="strong"
		}
	}
	fmt.Println(a)
       
}
