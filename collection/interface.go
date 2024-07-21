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

type IFilterFunc func(value any, key string, collection *ICollection) bool

type IFindFunc func(value any, key string, collection *ICollection) bool

type IReduceFunc func(accumulator any, value any, key string, collection *ICollection) any

type IMapFunc func(value any, key string, index int, collection *ICollection) any

type ISomeFunc func(value any, key string, collection *ICollection) bool

type ISweepFunc func(value any, key string, collection *ICollection) bool

type IEveryFunc func(value any, key string, collection *ICollection) bool

type IExecuteFunc func(collection *ICollection, size int)

type IEachFunc func(value any, key string, collection *ICollection)

type ISortFunc func(a any, b any) bool