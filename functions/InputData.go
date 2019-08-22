package functions

import (
    "strconv"
    "os"
    "bufio"
    "fmt"
)

// InputData takes a File name as the argument and returns the url and depth.
func InputData(InputFileName string) (url string,depth int64) {

    InputFile, ERROR1 := os.Open(InputFileName)
    if ERROR1 != nil {
        fmt.Println("Error opening input file.",ERROR1)
    }

    Iterator := bufio.NewScanner(InputFile)
    
    var InputData [2]string
    var i int = 0

    for Iterator.Scan() {
        InputData[i] = Iterator.Text()
        i++
    }

    url = InputData[0]
    depth,_ = strconv.ParseInt(InputData[1],10,64)

    return
}