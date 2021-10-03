# Merge sort on go with gorutin
[![CodeQL](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml)
[![CodeQL](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dark0ghost/merge_sort/actions/workflows/codeql-analysis.yml)
[![gobenchdata publish](https://github.com/dark0ghost/merge_sort/actions/workflows/go_bench.yml/badge.svg?branch=main)](https://github.com/dark0ghost/merge_sort/actions/workflows/go_bench.yml)

## About

***Golang program for implementation of Merge Sort with gorutin***

## Use
```go
type Int int

func (v Int) Compare(value interface{}) bool{
	return v < value.(Int)
}

func toCompareArray(array []int) (buffer []cre.Compare){
        buffer = []cre.Compare{}
	for _, i := range array {
		buffer = append(buffer, Int(i))
	}
	return
}

func main(){
        var test []cre.Compare
	test := toCompareArray([]int{1,3,9,6,34,21,12})
	value, err := MergeSort(test)
	if err != nil{
		// chechk error
	}
	// you operation
	return
}
```
