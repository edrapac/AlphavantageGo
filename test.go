package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bufio"
	"os"
	"strings"
)

func main() {
	base_url := "https://robinhood.com/stocks/" // base landing page for RH's stock information
	fmt.Println("Enter a Stock's ticker symbol here. Try AAPL")
	reader := bufio.NewReader(os.Stdin)
	fragment, _ := reader.ReadString('\n')
	base_url += strings.ToUpper(strings.TrimSpace(fragment)) // kind of an ugly way to do it but it gets the job done!
	

	response, err := http.Get(base_url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(bytes))
}