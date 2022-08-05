package map

// Used to structure all the initiations of sets, maps, collections, etc...
type structure struct{}

// Structure used for the map implementation.
type IMap struct {
	set  func(key string, value any) *IMap
	get  func(key string) *structure
	find func(key string) *structure
}
type IMapReturnable struct {
	// Creates a new Map with a specified name.
	// 
	// Return Type: string, int, float32, float64, struct
	New func(return_type string, name string) *IMap
}
