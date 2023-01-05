package main

import (
	"fmt"
	"learn/examples"
)

func main() {
	fmt.Println("Start learn...")

	fmt.Println("=== 1.Variables BEGIN ===")
	examples.RunVariables()
	fmt.Println("=== 1.Variables END ===")

	fmt.Println("=== 2.If BEGIN ===")
	examples.RunIf(10, 15)
	fmt.Println("=== 2.If END ===")

	fmt.Println("=== 3.Switch BEGIN ===")
	examples.RunSwitch(15)
	fmt.Println("=== 3.Switch END ===")

	fmt.Println("=== 4.For BEGIN ===")
	examples.RunFor(5, 7)
	fmt.Println("=== 4.For END ===")

	defer runDefer()

	fmt.Println("=== 6.Array BEGIN ===")
	examples.RunArray()
	fmt.Println("=== 6.Array END ===")

	fmt.Println("=== 7.Slice BEGIN ===")
	examples.RunSlice()
	fmt.Println("=== 7.Slice END ===")

	fmt.Println("=== 8.Function BEGIN ===")
	examples.RunFunction()
	fmt.Println("=== 8.Function END ===")

	fmt.Println("=== 9.ErrorAndPanic BEGIN ===")
	// examples.RunErrorAndPanic()
	fmt.Println("=== 9.ErrorAndPanic END ===")

	fmt.Println("=== 10.Pointers BEGIN ===")
	examples.RunPointers()
	fmt.Println("=== 10.Pointers END ===")

	fmt.Println("=== 11.WebServer BEGIN ===")
	// examples.RunWebServer()
	fmt.Println("=== 11.WebServer END ===")

	fmt.Println("=== 12.Struct BEGIN ===")
	examples.RunStruct()
	fmt.Println("=== 12.Struct END ===")

	fmt.Println("=== 13.Interface BEGIN ===")
	examples.RunInterface()
	fmt.Println("=== 13.Interface END ===")
}

func runDefer() {
	fmt.Println("=== 5.Defer BEGIN ===")
	fmt.Println("Эта функция сработала в конце, хотя ее вызвали в середине")
	fmt.Println("=== 5.Defer END ===")
}
