package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html"
	"regexp"
	//"bufio"
	//"os"
	//"strings"
)

func main() {
	base_url := "https://robinhood.com/stocks/AAPL" // base landing page for RH's stock information
	//fmt.Println("Enter a Stock's ticker symbol here. Try AAPL")
	//reader := bufio.NewReader(os.Stdin)
	//fragment, _ := reader.ReadString('\n')
	// base_url += strings.ToUpper(strings.TrimSpace(fragment)) // kind of an ugly way to do it but it gets the job done!
	

	response, err := http.Get(base_url)

	if err != nil {
		panic(err)
	}
	//r, _ := regexp.Compile(`</path></g></svg>[0-9]+%</h2><p>of [0-9]+ ratings`)
	r, _ := regexp.Compile(`Hold</div>.*[0-9]+%`)
	z := html.NewTokenizer(response.Body)
	previousStartTokenTest := z.Token()
	bytes, err := ioutil.ReadAll(response.Body)
	fullcorpus := (string(bytes))
	narrow := (r.FindString(fullcorpus))
	percent, _ := regexp.Compile(`[0-9]+%</span>`)
	fmt.Println(percent.FindString(narrow))
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
				if previousStartTokenTest.Data == "css-131wcsq" {
					fmt.Println(string(z.Text()),"you foudn me")
					continue
				}
				TxtContent := string(z.Text())
				if len(TxtContent) > 0 {
					if r.MatchString(TxtContent) {
						// matched := r.FindString(TxtContent)
						fmt.Println(TxtContent)
					}
					
				}
				
				defer response.Body.Close()
				break
			}
		}
	// bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	
	// fmt.Printf(string(bytes))
}