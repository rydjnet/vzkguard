package housemd

import (
	"testing"
	housemd "vzkguard/houseMD"
)

func TestCheckLatinCharacters(t *testing.T) {
	type tstMSG struct {
		ex   string
		want bool
	}
	testCasesMsg := []tstMSG{
		{"Да, понос водой прям до 10 раз (извините за подробности)", false},
		{"Привет", false},
		{"Ghbdtn", false},
		{"привет!", false},
		{"привеt", false},
		{"приvet", true},
		{"делf", false},
		{"LK452", false},
	}
	for _, ca := range testCasesMsg {
		got := housemd.CheckLatinCharacters(ca.ex)
		if got != ca.want {
			t.Errorf("got %t want %t", got, ca.want)
		}
	}

}
