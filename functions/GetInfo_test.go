package functions

import (
	"testing"
)

//const outfileInfoname="output_info.txt"
const url1="https://google.com"

func TestGetInfo(t *testing.T) {
	ch := make(chan string)
	flag := make(chan bool)
	/*exp_Info := make(chan[string]bool)
	Outputfile, err := os.Open(outfileInfoname)
	if err!=nil {
		t.Error("File containing expected output is not opening due to the following error",err)
	}
	Iterator := bufio.NewScanner(Outputfile)
	for Iterator.Scan() {
		temp_info = Iterator.Text()
		exp_Info[temp_info]=true
	}*/
	act_Info := make(map[string] bool)
	go GetInfo(url1,ch,flag)
	for c:=0;c<1; {
		select {
		case info:= <-ch:
			act_Info[info]=true
		case <-flag:
			c++
		}
	}
	/*for info,_:= range exp_Info {
		if !act_Info {
			t.Error("This Information is expected to be present\n",info)
		}
	}
	for info,_ := range act_Info {
		if !exp_Info {
			t.Error("This Information is not expected to be present\n",info)
		}
	}*/
}
