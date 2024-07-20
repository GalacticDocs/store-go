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
> Returns the value associated with the provided key, if non could be found, it returns nil.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to fetch. |
>
> Returns **{any}**

## [.Filter(fn)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L220)
> Returns a collection that has been filtered by said fn function.
> | Param | Type | Description |
> | --- | --- | --- |
> | fn | [IFilterFunc](https://github.com/GalacticDocs/store-go/blob/main/collection/interface.go#L16) | The function to execute. |
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md#top)}**