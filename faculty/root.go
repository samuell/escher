// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	"sync"

	. "github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

var lk sync.Mutex
var root = NewIndex()

func Root() Index {
	lk.Lock()
	defer lk.Unlock()
	return root
}

func Register(v Materializer, addr ...Name) {
	lk.Lock()
	defer lk.Unlock()
	root.Memorize(v, addr...)
}
