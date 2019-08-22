//Package main tells the Go compiler that the package should compile 
//as an executable program instead of a shared library.
package main

import (
	
	// Package os provides a platform-independent interface to 
	// Operating system functionality.
	"os"

	// Package fmt implements formatted I/O with functions analogous 
	// To C's printf and scanf. 
	"fmt"

	// Package goquery implements features similar to jQuery, including 
	// The chainable syntax, to manipulate and query an HTML document.
	"github.com/concurrency-6/functions"
)

func main() {

	InputFileName := os.Args[1]
	
// url stores the URL given in the input file.
// depth stores the distance upto which we have to search for the URLs.	
url,depth := functions.InputData(InputFileName)

	// URLS stores all the URLs that have to be parsed.
	var URLS []string

	// Output stores the information of every URL that has been fetched.
	var Output []string

	// URLSch is a channel that is passed as the argument for the getURLs function.
	// This stores the URLs published by the goroutine GetURLs.
	URLSch := make(chan string)

	// flagCh1 is a channel that is used to check if all the URLs have been fetched
	// From a goroutine GetURLs.
	flagCh1 := make(chan bool)

	// DataCh is a channel that is passed as the argument for the GetInfo function.
	// This stores the Information published by the goroutine GetInfo.
	DataCh := make(chan string)

// flagCh2 is a channel that is used to check if all the info has been fetched
// From a goroutine GetInfo.
flagCh2 := make(chan bool)

	var i int64 = 0
	URLS = append(URLS,url)

	for ; i < depth ; i++ {

		// length stores the length of the URLS slice at this instant.
		// Which can be used later when the URLS is sliced.
		length := len(URLS)

		for _,url := range URLS {
			go functions.GetURLs(url,URLSch,flagCh1)
		}

		// The following for loops makes sure the URLs fetched from each
		// URL in the URLS slice is appended back to the URLS slice.
		for j := 0 ; j < length ; {
			select {
				case urls := <- URLSch :
					URLS = append(URLS,urls)
				case <- flagCh1 :
					j++
			}
		}

		//Slice the URLS so that we don't have to go through the parsed URLs again. 
		URLS = URLS[length:]
		
		for _,url := range URLS {
			go functions.GetInfo(url,DataCh,flagCh2)
		}

		// The following for loops makes sure the Information fetched from each
		// URL in the URLS slice is appended to the Output slice.
		for k := 0 ; k < len(URLS) ; {
			select {
				case info := <- DataCh :
					Output = append(Output,info)
				case <- flagCh2 :
					k++
			}
		}
	}

	close(URLSch)
	close(DataCh)
	close(flagCh1)
	close(flagCh2)

	// fout is the file opbject that is used to output the information in the Output slice.
	fout, err := os.Create("Output.txt")
    if err != nil {
        fmt.Println("Error creating file",err)
        return
    }

	for _,out := range Output {
		l, err := fout.WriteString(out + "\n\n")
    	if err != nil {
        	fmt.Println(err)
        	fout.Close()
        	return
    	}
    	l = l + 1
	}
	fout.Close()
}
