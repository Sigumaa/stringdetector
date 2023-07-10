# stringdetector

## Install

```bash
go install github.com/Sigumaa/stringdetector/cmd/stringdetector@latest
```


## Example

```go
package main

type A struct {
	Url string
	Id  string
	Utc string
}

func main() {
	var Url string
	print(Url)

	var Id string
	print(Id)

	var Utc string
	print(Utc)
}
```

```bash
$ go vet -vettool=`which stringdetector` -stringdetector.flag="Url,Id,Utc" main
.\a.go:4:2: Url is detected
.\a.go:5:2: Id is detected
.\a.go:6:2: Utc is detected
.\a.go:10:6: Url is detected
.\a.go:11:8: Url is detected
.\a.go:13:6: Id is detected
.\a.go:14:8: Id is detected
.\a.go:16:6: Utc is detected
.\a.go:17:8: Utc is detected
```
