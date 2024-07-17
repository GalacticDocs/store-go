# Collection

* [Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md)

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

## [.Clone()](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L33)
> Clones the Collection and returns it.
>
> Returns **{[Collection](https://github.com/GalacticDocs/store-go/blob/main/Docs/Collection.md)}**

## [.Delete(key)](https://github.com/GalacticDocs/store-go/blob/main/collection/collection.go#L38)
> Deletes a key from the Collection.
> | Param | Type | Description |
> | --- | --- | --- |
> | key | string | The key to delete from the Collection. |
>
> Type **{bool}**
