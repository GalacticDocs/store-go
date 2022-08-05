package maps

import (
	"github.com/iVitaliya/logger-go"
)

func StringMap() *IMapReturnable {
	var store = make(map[string]string)

	return &IMapReturnable{
		New: func()
	}
}