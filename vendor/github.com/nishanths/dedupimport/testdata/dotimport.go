package pkg

import (
	"expvar"
	_ "expvar"
	. "testing"
)

import (
	"testing"
)

// TODO: maybe consider removing this import since it's repeated? 
// we don't at the moment, to keep the logic around special handling 
// related to _ and . simple.
import _ "expvar"