# Collection

* [Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)

**Extensions**
* [Cache](https://github.com/GalacticDocs/store-go/blob/main/Docs/Cache.md)
* [LiveCache](https://github.com/GalacticDocs/store-go/blob/main/Docs/LiveCache.md)

**Managers**
* [Manager](https://github.com/GalacticDocs/store-go/blob/main/Docs/Manager.md)
* [DataDock](https://github.com/GalacticDocs/store-go/blob/main/Docs/DataDock.md)

```golang
import "github.com/GalacticDocs/store-go/"

func main() {
    col := store.Collection()
}
```

# Values
## [.Size()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L)
> Amount of keys/values in the Collection.
>
> Returns **{int}**

# Methods
## [.Clear()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L24)
> Removes all the items from the Collection.
>
> Returns **{bool}**

## [.Clone()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L36)
> Clones the Collection and returns it.
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Delete(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L76)
> Deletes a key from the Collection.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to delete from the Collection. |
>
> Returns **{bool}**

## [.Difference(against)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L95)
> Returns a new Collection containing only the properties ones that don't exist in against.
> | Param | Type | Description |
> | --- | --- | --- |
> | against | [Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md) | The collection to compare against. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Each(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L119)
> Runs the given function over every key-value pair in the collection.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IEachFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L32) | The function to execute. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Every(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L143)
> Runs the given function over the entire collection on every element.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IEveryFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L28) | The function to execute.
>
> Returns **{bool}**

## [.Execute(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L170)
> Executes a function over the collection. (different from .Every() & .Each())
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IExecuteFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L30) | The function to execute.
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Exist(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L184)
> Returns whether or not the provided key exists in the collection.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to search for. |
>
> Returns **{bool}**

## [.Fetch(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L198)
> Searches for a value by the associated key. It will return nil if no key-value could be found.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to search with. |
>
> Returns **{any}**

## [.Filter(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L220)
> Filters the collection by the provided function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IFilterFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L16) | The function to execute. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Find(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#249)
> Searches for a value in the collection using the provided function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IFindFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L18) | The function to use for searching.
>
> Returns **{any}**

## [.First()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L267)
> Searches the first value in the collection.
>
> Returns **{any}**

## [.Get(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L293)
> Searches for a value by the associated key. It will return nil if no key-value could be found.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to search with. |
>
> Returns **{any}**

## [.GetIndex(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L305)
> Searches for the index by the provided key.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to search with. |
>
> Returns **{int}**

## [.Has(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L323)
> Checks if the collection has the provided key.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to search with. |
>
> Returns **{bool}**

## [.Implement(...collections)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L336)
> Implements the provided collections into one.
> | Param | Type | Description |
> | --- | --- | --- |
> | collections | ...[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top) | An array of collections to merge into one. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Intersect(secondary)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L355)
> Returns a collection containing the intersection of the current and given collection.
> | Param | Type | Description |
> | --- | --- | --- |
> | secondary | [Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/collection.md#top) | The collection to intersect. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/collection.md#top)}**

## [.Iterator()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L364)
> Iterates over the collection using the GO built-in map[string]any
>
> Returns **{map[string]any}**

## [.Last()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L377)
> Searches the last value in the collection.
>
> Returns **{any}**

## [.Map(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L407)
> Maps the elements in the collection and returns it as an array.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IMapFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L22) | The function to execute. |
>
> Returns **{[]any}**

## [.Merge(...collections)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L428)
> Merges the provided collections into one.
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Reduce(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L460)
> Applies a function against an accumulator and each value of the Collection (from left-to-right) to reduce it to a single value.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IReduceFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L20) | The function to execute. |
>
> Returns **{any}**

## [.Set(key, value)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L497)
> Sets a value for the specified key. Or adds the value with the provided key if none defined yet.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to set a value for. |
> | value | any | The value to use. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Some(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L529)
> Checks if at least one element in the Collection satisfies the given function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [ISomeFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L24) | The function to execute. |
>
> Returns **{bool}**

## [.Sort(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L552)
> Sorts the elements in the collection using the provided function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [ISortFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L34) | The function to execute. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**

## [.Sweep(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L596)
> Removes all elements in the collection that match the given function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [ISweepFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L26) | The function to execute. |
>
> Returns **{int}**

## [.ToArray()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L613)
> Returns all the key-value pairs in the collection as an array.
>
> Returns **{[][ICollectionAll](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L11)}**

## [.ToKeyArray()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L626)
> Returns all the keys in the collection as an string array.
>
> Returns **{[]string}**