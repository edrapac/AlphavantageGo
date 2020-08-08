package main
import (
	"fmt"
	 // "net/http" soup module will handle making the actual requests
	// "io/ioutil"
	"bufio"
	"os"
	"strings"
	"github.com/anaskhan96/soup"
)

func main() {
	base_url := "https://robinhood.com/stocks/" // base landing page for RH's stock information
	fmt.Println("Enter a Stock's ticker symbol here. Try AAPL")
	reader := bufio.NewReader(os.Stdin)
	fragment, _ := reader.ReadString('\n')
	base_url += strings.ToUpper(strings.TrimSpace(fragment)) // kind of an ugly way to do it but it gets the job done!
	

	response, err := soup.Get(base_url)

	if err != nil {
		panic(err)
	}

	// defer response.Body.Close()
	doc := soup.HTMLParse(response)

	if err != nil {
		panic(err)
	}
	price := doc.Find("span","class","up")
	fmt.Println(price)
}