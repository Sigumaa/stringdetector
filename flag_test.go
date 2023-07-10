package stringdetector

func ExportSet(s string) func() {
	o := flag
	flag = s
	return func() {
		flag = o
	}
}
