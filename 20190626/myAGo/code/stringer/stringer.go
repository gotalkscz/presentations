package painkiller

import (
	"fmt"
)

// START OMIT
type Pill int

//go:generate stringer -type=Pill
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

var _ fmt.Stringer = Pill(nil)

// END OMIT
