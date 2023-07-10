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
.\main.go:4:2: Url is detected
.\main.go:5:2: Id is detected
.\main.go:6:2: Utc is detected
.\main.go:10:6: Url is detected
.\main.go:11:8: Url is detected
.\main.go:13:6: Id is detected
.\main.go:14:8: Id is detected
.\main.go:16:6: Utc is detected
.\main.go:17:8: Utc is detected
```
