package main

import "fmt"

func main() {
	fuel := 0
	for {
		var mass int
		_, err := fmt.Scanf("%d", &mass)
		if err != nil {
			break
		}
		module_fuel := mass/3 - 2
		additional := module_fuel/3 - 2
		for additional > 0 {
			module_fuel += additional
			additional = additional/3 - 2
		}
		fuel += module_fuel
	}
	fmt.Println(fuel)
}
