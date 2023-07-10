# stringdetector


```go
package main

func main() {
    var Url string
    print(Url)
}
```

```bash
$ go vet -vettool=`which stringdetector` -stringdetector.flag="Url" a
.\a.go:4:6: Url is detected
.\a.go:5:8: Url is detected
```
