# GoRH
Simple RH stock market scraper written in Go


`go get github.com/anaskhan96/soup`

`gun test2.go | grep -Eo '\\"price\\",\\"[0-9]+\.[0-9]+' --color=yes | less` < -- Debug stuff

`gun test2.go | grep -m 1 -Eo '\\"price\\",\\"[0-9]+\.[0-9]+' | awk -F '"' '{print $(NF)}' | head -1`
