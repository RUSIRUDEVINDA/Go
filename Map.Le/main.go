package main

import ("fmt"
	"sort"
)

func main(){
	shows:=make(map[string]string)
	shows["PK"]="Peaky Blinders"
	shows["BB"]="Breaking Bad"
	shows["Got"]="Game of thrones"

	fmt.Println(shows)

	peakybliners:=shows["PK"]
	fmt.Println(peakybliners)
	
	delete(shows,"BB")
	fmt.Println(shows)

	shows["SPN"]="Supernatural"
	fmt.Println(shows)

	for k,v :=range shows{
		fmt.Printf("%v: %v\n", k,v)	
	}

	Shows:=make([]string, len(shows))
	i:=0
	for s := range shows{
		Shows[i]=s
		i++
	}

	sort.Strings(Shows)
	fmt.Println("\nSorted")

	for i:=range Shows{
		fmt.Println(shows[Shows[i]])
	}
}