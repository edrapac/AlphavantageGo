package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html"
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
	base_url += strings.ToUpper(strings.TrimSpace(fragment)) // kind of an ugly way to do it but it gets the job done!
	

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

	if len(hold.FindString(fullcorpus)) > 0 {
		percent_hold := strings.Replace(holdpercent.FindString(hold.FindString(fullcorpus)),"</span>","",-1)
		fmt.Println(percent_hold,"Analysts say to hold",fragment)
	}

	z := html.NewTokenizer(response.Body)
	previousStartTokenTest := z.Token()
	loopOver:
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
			// EOF
				break loopOver
			case tt == html.StartTagToken:
				previousStartTokenTest = z.Token()
			case tt == html.TextToken:
				if previousStartTokenTest.Data == "span" {
					continue
				}
				TxtContent := string(z.Text())
				if len(TxtContent) > 0 {
					if price.MatchString(TxtContent) {
						matched := price.FindString(TxtContent)
						final := len(matched)
						fmt.Println("Current trading price: ","$",strings.TrimSpace(matched[12:final]))
					}
					
				}
				
				defer response.Body.Close()
				break
			}

		}

	if err != nil {
		panic(err)
	}
	
}