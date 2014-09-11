// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	. "github.com/gocircuit/escher/circuit"
)

// Genus_ is a name type for a genus structure.
type Genus_ struct{}

func (Genus_) String() string {
	return "*Genus"
}

//
type FacultyGenus struct {
	Acid map[string]string // acid to directory
	Walk []Name // walk within hierarchy
}

func NewFacultyGenus() *FacultyGenus {
	return &FacultyGenus{
		Acid: make(map[string]string),
	}
}

type CircuitGenus struct {
	Dir, File string
}