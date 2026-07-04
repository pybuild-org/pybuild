package standalone

import "fmt"

func Build() {
	for _, target := range Targets {
		fmt.Println(target)
	}
}
