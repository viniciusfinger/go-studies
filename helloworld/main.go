package main

func sum(x int, y int) int {
	return x + y
}

// return more than one attribute
func sumWithConfirmation(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}

func main() {
	//Declarating vars. Go is strongly typed
	var message string
	message = "Hello world!"
	//message = 1 cause errors because type is different

	//Declarating and instantiating var, use ':' only in first atribuition, like Pascal language.
	message2 := "Hello world again!"

	//message3 := "Garbage" <- This line prevents compilation.
	//Go is restrictive and does not allow compiling with unused variables

	//the next atribuitions don't need ':'
	message2 = "Hello new world!"

	println(message)
	println(message2)

	println(sum(1, 2))

	result, xGreatherThan10 := sumWithConfirmation(10, 20)
	println(result, xGreatherThan10)
}
