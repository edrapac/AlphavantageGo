package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"bufio"
	"os"
	"strings"
)

func main() {
	base_url := "https://robinhood.com/stocks/" // base landing page for RH's stock information
	fmt.Println("Enter a Stock's ticker symbol here. Try AAPL")
	reader := bufio.NewReader(os.Stdin)
	fragment, _ := reader.ReadString('\n')
	new_fragment := strings.ToUpper(strings.TrimSpace(fragment))
	base_url += new_fragment // kind of an ugly way to do it but it gets the job done!
	

	response, err := http.Get(base_url)

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fullcorpus := (string(bytes))
	price, _ := regexp.Compile(`\\"price\\",\\"[0-9]+.[0-9]+`)
	hold, _ := regexp.Compile(`Hold</div>.*[0-9]+%`)
	holdpercent, _ := regexp.Compile(`[0-9]+%</span>`)
	if len(price.FindString(fullcorpus)) > 0 {
		matched_price := strings.Replace(price.FindString(fullcorpus),`\"price\",\"`,"",-1)
		
		fmt.Println("\n",new_fragment," is currently trading at",matched_price)
	}
	if len(hold.FindString(fullcorpus)) > 0 {
		percent_hold := strings.Replace(holdpercent.FindString(hold.FindString(fullcorpus)),"</span>","",-1)
		fmt.Println(percent_hold,"Analysts say to hold",fragment)
	}
	
}