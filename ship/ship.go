package ship

import "fmt"

type Ship struct {
	Name string
	year int
}

/*
 * As this is a package, main() method is not needed neither executed
func main() {
	fmt.Println("This is package ship")
}
*/

func Stop() {
	fmt.Println("stop the engines")
}