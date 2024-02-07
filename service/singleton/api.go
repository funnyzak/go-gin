package singleton

import "sync"

var (
	ApiLock sync.RWMutex
)
