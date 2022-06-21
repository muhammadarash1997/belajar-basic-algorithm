package main
 
import (
  "fmt"
  "strings"
)
 
func main() {

	// there are 3 times 'e' in 'cheese'
	fmt.Println(strings.Count("cheese", "e"))

	fmt.Println(strings.Count("five", ""))

	// there are 1 times 's' in 'cheese'
	fmt.Println(strings.Count("cheese", "s"))

	// there are 2 times 'wi' in 'JohnwickJohnwick'
	fmt.Println(strings.Count("JohnwickJohnwick", "wi"))
}