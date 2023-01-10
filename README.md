# TheDeveloper10/rem

Package `TheDeveloper10/errorland` makes it easier to manage lists of errors.

___

* [Install](#install)
* [Examples](#examples)

___

## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/TheDeveloper10/errorland@v1.0.0
```

## Examples
Let's try it
```go
func main() {
	// initialize an empty land to store errors
	errs1 := errorland.NewLand()
	
	// initialize a land with 1 error: "my error"
	errs2 := errorland.NewLandFromError(errors.New("my error"))
	
	// initialize a lnad with 1 error: "my next error"
	errs3 := errorland.NewLandFromString("my next error")
	
	// merge land 2 in land 1
	errs1.Merge(errs2)
	
	// merge land 3 in land 1
	errs1.Merge(errs3)
	// now land 1 has: "my error", "my next error"
	
	// append a new error to land 1
	errs1.AppendString("err 3")
    errs1.AppendError("err 4")
	// now land 1 has: "my error", "my next error", "err 3", "err 4"
	
	// remove an error from index
	errs1.RemoveIndex(1)
	// now land 1 has: "my error", "err 3", "err 4"
	
	log.Printf("Total error count: ", errs1.GetCount())
	// Total error count: 3
}
```
