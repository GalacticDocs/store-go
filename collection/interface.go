package collection_store

import (
	"sync"
)

type ICollection struct {
	store sync.Map
}

type ICollectionAll struct {
	Key   string
	Value any
}

type IFilterFunc func(key string, value any, collection *ICollection) bool

type IFindFunc func(key string, value any, collection *ICollection) bool

type IReduceTransformer func(accumulator any, value any, key string, collection *ICollection) any

type IMapTransformer func(value any, key string, index int, collection *ICollection) any

type ISomeFunc func(key string, value any, collection *ICollection) bool

type ISweepFunc func(key string, value any, collection *ICollection) bool

type IEveryFunc func(value any, key string, collection *ICollection) bool