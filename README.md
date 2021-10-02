# merge_sort
[![CodeQL](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml)
[![CodeQL](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml)

Golang program for implementation of Merge Sort with gorutin

use:
```go
type Int int

func (v Int) Compare(value interface{}) bool{
	return v < value.(Int)
}

func main(){
  var text []cre.Compare
	text = []cre.Compare{Int(8), Int(1), Int(5), Int(7), Int(8), Int(10)}
	value, err := MergeSort(text)
	if err != nil{
		// chechk error
	}
	// you operation
	return
}
```
