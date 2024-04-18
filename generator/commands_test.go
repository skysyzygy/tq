package main

import (
	"fmt"
	"testing"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/tq"
)

func Test_Describe(t *testing.T) {
	tq := tq.New(nil, false, false)
	tq.Login(auth.New("", "", "", "", nil))
	short, long := Describe(tq.Put.ConstituentsUpdate)
	fmt.Println(short, long)
}

func Test_MarshalJSON(t *testing.T) {

}
