package main

import (
	"fmt"
	"golang/examples/codepractice"
)

/*func main() {
	fmt.Println("Hello World... Let's GO")
	//WrapError()
	//codepractice.OpenFile("sample.txt")
	//codepractice.ReadFile("main.go")
	//codepractice.ReadFileLineByLine("main.go")
	codepractice.MapDemo("4")
	codepractice.MapIteration()

	var kantoPokemon = make(map[string]int, 151)
	fmt.Println(len(kantoPokemon))
}*/
/*
func main() {
	codepractice.DisplayStructVal()
	codepractice.MapOfPersonsDemo()
	codepractice.UnfurlingSlice()
	codepractice.ReceiverExample()
	codepractice.TestInterface()
	codepractice.ClosureDemo()
	codepractice.Demo()
	codepractice.StructsInterfaceDemo()
	Println(codepractice.Factorial(10))
}
*/

func init() {
	fmt.Println("This will run once before the main method.")
}

func main() {
	/*	std_lib_practice.JsonMarshal()
		fmt.Println()
		std_lib_practice.JsonUnMarshal()*/
	//data := std_lib_practice.ReadJsonData()

	//std_lib_practice.ExampleSortCustomType()
	//codepractice.ChannelDemo()
	//codepractice.UnidirChannelDemo()
	//codepractice.DemoChannelRange()
	//codepractice.SelectChannelDemo()
	//codepractice.FanInDemo()
	//codepractice.FainInDemo2()
	codepractice.CustomeErrorDemo(10)
	codepractice.CustomeErrorDemo(-10)
}
