package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"math/rand"
	"time"
	"strings"
)
func fragment() string {
	base_url := "https://robinhood.com/stocks/"
	fragment := ""
	for i := 0; i < 4; i++ {
		rand.Seed((time.Now().UnixNano()))
    	min := 65
		max := 90
		newascii := rand.Intn(max - min)+min
		fragment += string(newascii)
	}
	new_fragment := strings.ToUpper(fragment)
	base_url += new_fragment // kind of an ugly way to do it but it gets the job done!
	return base_url
}
func body(url string) string {
	response, err := http.Get(url)
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fullcorp := string(bytes)
	return fullcorp
}

func main() {
	 // base landing page for RH's stock information
	fmt.Println("Randomly selecting a stock")
	new_fragment := ""
	fullcorpus := ""
	for {
		new_url := fragment()
		new_fragment = new_url 
		fullcorpus = body(new_url)
		// fmt.Println(new_url) lol
		if len(fullcorpus) > 89280 {
			break
		}
	}
	
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