##                                                **WEB SCRAPER**

[![CircleCI](https://circleci.com/gh/IITH-SBJoshi/concurrency-6.svg?style=svg&circle-token=14ee5065c2b3b4734bcbfaf926a2b4f48c2acab0)](https://circleci.com/gh/IITH-SBJoshi/concurrency-6)  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Web scraping is data scraping used for extracting data from websites. Web scraping may access the World Wide Web using a Hypertext Transfer Protocol or a web browser. While web scraping can be done manually by a software user, the term typically refers to automated processes using a bot or a web crawler. 

Our project works by taking an URL and depth as input and creating a text file which contains the list of URLs in the input URL upto the given depth.
The search for the URLs is based on BFS.
The output can be resembled with a tree model with input URL as root and its children are the list of URLs present in the link.
The above process continues until the required depth is obtained.


### The execution flow is :

At the beginning of the program we read the input from the file input.txt using [InputData.go](https://github.com/IITH-SBJoshi/concurrency-6/blob/master/functions/InputData.go) program.

Then we make use GetURLs.go program created in Functions directory to search for URLs present in the given input URL page.
In the [GetURLs.go](https://github.com/IITH-SBJoshi/concurrency-6/blob/master/functions/GetURLs.go) program we import strings and goquery packages.
We create GetURLs() function which takes an URL and two channels as arguments.
Out of two channels one channel is used to store all the fetched URLs concurrently from each goroutine.
And the other channel is used  as a notification channel which publishes a done message after a goroutine fetches all URLs.

Then after fetching an URL we make use of GetInfo.go program to acquire title and metadata in the webpage associated with the URL.
In the [GetInfo.go](https://github.com/IITH-SBJoshi/concurrency-6/blob/master/functions/GetInfo.go) program we import strings and goquery packages.
We create GetInfo() function which takes an URL and two channels as arguments.

At the end of the program we generate an output file named output.txt containing all the URLs traversed by our program and certain information about the links.


The [goquery](https://github.com/PuerkitoBio/goquery/blob/master/LICENSE)  package was imported from github.


##### User Manual :

$ cd go/src/github.com
 
$ git clone https://github.com/IITH-SBJoshi/concurrency-6

$ cd

$ mv /$HOME/go/src/github.com/concurrency-6/main/input.txt /$HOME/go/bin/

$ cd go/bin

$ go build github.com/concurrency-6/functions

$ go install github.com/concurrency-6/main

$ ./main input.txt

###### For documenting go program :

$ godoc cmd/github.com/concurrency-6/functions 

#### Developer's Manual :

-> To detect and block malicious webpages.

-> Can be used to extract large amount of data from a specific webpage.


#### Contributors :

-> RAVI VENKATA KRISHNA         :     CS17BTECH11032

-> PRAMOD REDDY					            :     CS17BTECH11028

-> NIKHIL KUMAR REDDY           :     CS17BTECH11008

-> ASHWANTH KUMAR				           :     CS17BTECH11017
