package pageable

import (
	"fmt"
	"log"
	"testing"
)

func TestPageQuery(t *testing.T) {
	// WIP
}

func TestSetDefaultRPP(t *testing.T) {
	//set valid rpp
	if SetDefaultRPP(25) != nil {
		t.Fatal("set valid rpp but occur an error")
	}
	//set invalid rpp
	if SetDefaultRPP(-1) == nil {
		t.Fatal("set invalid rpp but occur no error")
	}
}

func TestSetRecovery(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(fmt.Sprint(err))
		}
	}()
	SetRecovery(func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	})
}
