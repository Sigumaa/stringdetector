# stringdetector


```go
package main

func main() {
    var Url string
    print(Url)
}
```

```bash
$ go vet -vettool=`which stringdetector` -stringdetector.flag="Url" main
.\main.go:4:6: Url is detected
.\main.go:5:8: Url is detected
```
