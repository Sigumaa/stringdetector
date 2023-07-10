package a

type A struct {
	Url string // want "Url is detected"
	Id  string // want "Id is detected"
	Utc string // want "Utc is detected"
}

func f() {
	var Url string // want "Url is detected"
	print(Url)     // want "Url is detected"

	var Id string // want "Id is detected"
	print(Id)     // want "Id is detected"

	var Utc string // want "Utc is detected"
	print(Utc)     // want "Utc is detected"
}
