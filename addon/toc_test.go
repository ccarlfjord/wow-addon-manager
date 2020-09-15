package addon

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	file := `
## Interface: 11302

## Title: Azeroth Auto Pilot <Classic> 1-60
## Notes: AAP Classic
## Author: Zyrrael
## Version: 0.13
## SavedVariables: AAPC1, AAPC2, AAPC3, AAPCQuestNames, AAPC_HoldZoneVar
`

	reader := strings.NewReader(file)

	toc := tocReader(reader)
	fmt.Printf("%#v\n", toc)
}
