// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reflect

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// TODO: This is first order germ projection. It can be extended to infinite order.
type Index struct {
	gv Circuit // Gate-Valve dictionary
	shadow Circuit
}

func (x *Index) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	x.gv, x.shadow = New(), New()
	return nil
}

func (x *Index) CognizeFlowFrame(eye *be.Eye, v interface{}) {

	// place involved gates in dictionary
	f := v.(Circuit)
	ag := x.remember(f, 0)
	bg := x.remember(f, 1)

	// place flow links in shadow
	ak, bk := x.shadow.Degree(ag), x.shadow.Degree(bg)
	x.shadow.Link(Vector{ag, ak}, Vector{bg, bk})
}

//	….Gate.Valve.[Index]
func (x *Index) remember(frame Circuit, i int) (index int) {

	half := frame.CircuitAt(i)
	valve := half.NameAt("Valve")

	// remember gate in Gate-Valve dictionary
	var g Circuit
	switch t := half.At("Value").(type) {
	case Address:
		g = x.gv.Place(t, New()).(Circuit)
	case int:
		g = x.gv.Refine("int")
	case string:
		g = x.gv.Refine("string")
	case float64:
		g = x.gv.Refine("float64")
	case complex128:
		g = x.gv.Refine("complex128")
	case Circuit:
		g = x.gv.Refine("meaningless")
	default:
		g = x.gv.Refine("unknown")
	}

	// and valve
	var ok bool
	if index, ok = g.IntOptionAt(valve); !ok {
		index = x.shadow.Len()
		g.Gate[valve] = index
	}

	// place gate in index
	x.shadow.Include(valve, "GateValve")

	return index
}

func (x *Index) CognizeControl(eye *be.Eye, v interface{}) {
	gv, shadow := x.gv, x.shadow
	x.gv, x.shadow = New(), New()
	eye.Show(DefaultValve, New().Grow("GateValve", gv).Grow("Shadow", shadow))
}

func (x *Index) Cognize(eye *be.Eye, v interface{}) {}