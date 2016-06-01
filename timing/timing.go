package timing

import (
	"time"

	"github.com/bcokert/terragen/log"
)

// Track tracks the elapsed time from the given start time.
// It should be called in a defer withing the function we are tracking
func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Info("%s took %s", name, elapsed)
}
