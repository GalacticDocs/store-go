package map_store

import (
	"sync"
)

type IMapAll struct {
	Key   string
	Value any
}

type IMap struct {
	store sync.Map
}
