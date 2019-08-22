package functions

import (
	"testing"
)
const filename="input.txt"

func TestInputData(t *testing.T) {
	act_url,act_depth := InputData(filename)
	var exp_depth int64
	exp_url,exp_depth := "https://www.amazon.in",1
	if act_url!=exp_url {
		t.Error("Instead of this url",exp_url,"we got this url",act_url)
	}
	if act_depth!=exp_depth {
		t.Error("Instead of this depth",exp_depth,"we got this depth",act_depth)
	}
}
