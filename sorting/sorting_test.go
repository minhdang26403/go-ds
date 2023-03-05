package sorting

import (
	_ "fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var inputSize = 10000

func randomList(listSize int) ([]int){
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, listSize)
	for i := range list {
		list[i] = random.Intn(1000)
	}
	return list
}

func generateInputList(numList, listSize int) ([][]int) {
	lists := make([][]int, numList)
	for i := 0; i < numList; i++ {
		list := randomList(listSize)
		lists[i] = list
	}
	return lists
}

func TestBubbleSort(t *testing.T) {
	array1 := randomList(inputSize)
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	BubbleSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	lists := generateInputList(b.N, inputSize)
	b.StartTimer()
	for _, list := range lists {
		BubbleSort(list)
	}
}

func TestSelectionSort(t *testing.T) {
	array1 := randomList(inputSize)
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	SelectionSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	lists := generateInputList(b.N, inputSize)
	b.StartTimer()
	for _, list := range lists {
		SelectionSort(list)
	}
}

func TestMergeSort(t *testing.T) {
	array1 := randomList(inputSize)
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	MergeSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.StopTimer()
	lists := generateInputList(b.N, inputSize)
	b.StartTimer()
	for _, list := range lists {
		MergeSort(list)
	}
}

func TestQuickSort(t *testing.T) {
	array1 := randomList(inputSize)
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	QuickSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	lists := generateInputList(b.N, inputSize)
	b.StartTimer()
	for _, list := range lists {
		MergeSort(list)
	}
}