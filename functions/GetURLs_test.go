package functions

import (
	"testing"
)


//const outfilename="output_url.txt"
const url2="https://google.com"

func TestGetURLs(t *testing.T) {
	ch := make(chan string)
	flag := make(chan bool)
	/*exp_foundUrls := make(map[string]bool)
	Outputfile, err := os.Open(outfilename)
	if err!=nil {
		t.Error("File containing expected output is not present")
	}
	Iterator := bufio.NewScanner(Outputfile)
	for Iterator.Scan() {
		temp_url = Iterator.Text()
		if strings.Contains(temp_url,"http") {
			exp_foundUrls[temp_url]=true
		}
	}*/
	act_foundUrls := make(map[string]bool)
	go GetURLs(url2,ch,flag)
	for c:=0;c<1; {
		select {
		case link:=<-ch:
			act_foundUrls[link]=true
		case <-flag:
			c++
		}
	}
	/*for Url,_ := range exp_foundUrls {
		if !act_foundUrls[Url] {
			t.Error("The url",Url,"is expected to be found")
		}
	}
	for Url,_ := range act_foundUrls {
		if !exp_foundUrls[Url] {
			t.Error("The url",Url,"is not expected to be found")
		}
	}*/
}
