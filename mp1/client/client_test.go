package main

import(
   "fmt"
   "testing"
)
 
func TestQueryTotalCount(t *testing.T) {
	very_frequent := QueryTotalCount("a")
	if very_frequent == 6353353 {
			fmt.Print("Success\n")
	} else {
		fmt.Print("Error!\n")
	}

	little_frequent := QueryTotalCount("ab")
	if little_frequent == 152281  {
			fmt.Print("Success\n")
	} else {
		fmt.Print("Error!\n")
	}

	not_frequent := QueryTotalCount("abc")
	if not_frequent == 1861   {
			fmt.Print("Success\n")
	} else {
		fmt.Print("Error!\n")
	}
}