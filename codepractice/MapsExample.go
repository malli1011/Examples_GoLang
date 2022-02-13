package codepractice

import "fmt"

//var ranks map[string]string = make(map[string]string)
//with initial capacity
var ranks = make(map[string]string, 10)

/*var ranks = map[string]string{
	"1": "One",
	"2": "Two",
	"3": "Three",
	"4": "four",
}*/

func MapDemo(key string) {
	ranks["gold"] = "🥇"
	ranks["silver"] = "🥈"
	ranks["bronze"] = "🥉"

	fmt.Println(ranks)
	fmt.Println("Length of map: ", len(ranks))
	delete(ranks, "4")
	fmt.Println("Length of map: ", len(ranks))
	if val, ok := ranks[key]; ok {
		fmt.Printf("The ranks,%s, is %s ", key, val)
		return
	}
	fmt.Printf("The ranks,%s, is not in the map \n", key)
}

func MapIteration() {
	movieRatings := map[string]int{
		"The Matrix":          88,
		"Speed":               94,
		"The Matrix Reloaded": 73,
		"John Wick":           86,
	}

	for key := range movieRatings {
		fmt.Println("Key :", key)
		fmt.Println("Value :", movieRatings[key])
	}

	for key, val := range movieRatings {
		fmt.Println("Key: ", key, "Value: ", val)
	}
}
