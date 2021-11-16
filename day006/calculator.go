package main

import "fmt"

func main() {
	fnMap := map[string]func(int, int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
	}
	fmt.Println(fnMap)

	type exercise struct {
		fn     string
		value1 int
		value2 int
	}

	todos := []exercise{
		{"+", 4, 2},
		{"pow", 2, 10},
		{"-", 4, 2},
	}

	for _, todo := range todos {
		fmt.Println(todo)
		fn, ok := fnMap[todo.fn]
		if !ok {
			fmt.Println("function not found: ", todo.fn)
			continue
		}

		value1 := todo.value1
		value2 := todo.value2
		result := fn(value1, value2)
		fmt.Printf("calculating %d %s %d: %d\n", value1, todo.fn, value2, result)
	}
}
