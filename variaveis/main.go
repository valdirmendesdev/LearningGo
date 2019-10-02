package main

func main() {
	explicitDeclaration(1, 2)
	implicitDeclaration()
}

func explicitDeclaration(num1 int32, num2 int32) {
	//Simple declaration
	var number1 int32
	var number2 int32
	number1 = num1
	number2 = num2

	//Multiple declaration
	var name, description string
	name = "Valdir"
	description = "Learning GO"

	//Simple declaration with defined value
	var year int16 = 2019

	//Multiple declaration with defined values
	var lastYear, actualYear int32 = 2018, 2019

}

func implicitDeclaration() {

	//Simple implicit declaration
	implicitName := "Valdir"

	//Multiple implicit declaration with defined values
	lastYear, actualYear := 2018, 2019

}
