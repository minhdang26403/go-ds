package hashtable

import (
	"testing"
	"crypto/sha256"
)

func HashStringKey(key string) int {
	sum := 0
	for _, v := range sha256.Sum256([]byte(key)) {
		sum += int(v)
	}
	return sum
}

func HashIntKey(key int) int {
	return key
}

func TestHashTableStringKey(t *testing.T) {
	ht := NewHashTable(HashStringKey)
	input := []string{"yang", "yejin", "yoona", "joy", "md", "cs", "abc", "data", "code", "go"}
	for _, v := range input {
		ht.Insert(v, v)
	}

	if got, _ := ht.Get("yejin"); got != "yejin" {
		t.Errorf("Got %v expected %v", got, "yejin")
	}

	ht.Delete("abc")
	if got, err := ht.Get("abc"); got != "" && err == nil {
		t.Errorf("Got %v expected %v", got, "")
	}

	ht.Insert("yejin", "cute")
	if got, _ := ht.Get("yejin"); got != "cute" {
		t.Errorf("Got %v expected %v", got, "cute")
	}
}

func TestHashTableIntegerKey1(t *testing.T) {
	ht := NewHashTable(HashIntKey)
	for i := 0; i < 10; i++ {
		ht.Insert(i + 1, i + 1)
	}

	ht.Delete(2)
	ht.Delete(6)
	ht.Delete(4)

	if got := ht.Contains(1); got != true {
		t.Errorf("Got %v expected %v", got, true)
	}

	if got := ht.Contains(2); got != false {
		t.Errorf("Got %v expected %v", got, false)
	}
	ht.Insert(5, 10)
	if got, _ := ht.Get(5); got != 10 {
		t.Errorf("Got %v expected %v", got, 10)
	}
}

func TestHashTableIntegerKey2(t *testing.T) {
	ht := NewHashTable(HashIntKey)

	for i := 0; i < 10; i++ {
		ht.Insert(i + 1, i + 1)
	}
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			ht.Delete(i)
		}
	}

	if got := ht.Contains(5); got != true {
		t.Errorf("Got %v expected %v", got, true)
	}
	if got := ht.Contains(8); got != false {
		t.Errorf("Got %v expected %v", got, false)
	}
	if got, _ := ht.Get(7); got != 7 {
		t.Errorf("Got %v expected %v", got, 7)
	}
	// Get value of non-exist key
	if got, err := ht.Get(4); got != 0 && err == nil {
		t.Errorf("Got %v expected %v", got, 0)
	}
}