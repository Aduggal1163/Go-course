package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites,"Amazon Web Services")
	fmt.Println(websites)
	for k,v := range websites {
		fmt.Print("Key is: ",k)
		fmt.Print(" Value is: ",v)
		fmt.Println()
	}
	
}
