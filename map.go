package store

import (
	"github.com/iVitaliya/logger-go"
)

func Map() *IMapReturnable {
	return &IMapReturnable{
		New: func(return_type string, name string) *IMap {
			var store = make(map[string]interface{})
			var _map = Map().New(return_type, name)

			return &IMap{
				find: func(key string) *structure {
					_, ok := store[key]
					if ok {
						res := store[key]
						return &res
					}

					logger.Error("Specified key couldn't be found in the map storage.")
					return nil
				},

				set: func(key string, value any) *IMap {
					if key == "" {
						logger.Error("Can't set a empty key in a map storage.")
						return Map().New(structureble)
					}

					if _map.find(key) == nil {
						store[key] = value
						return Map().New(structureble)
					}
				},
			}
		},
	}
}
