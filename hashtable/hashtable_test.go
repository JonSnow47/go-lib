//
// Revision History:
//     Initial: 2019-02-18 09:57    Jon Snow

package hashtable

import (
	"testing"
)

type robot struct {
	Id   int
	Name string
}

func TestHashtable(t *testing.T) {
	table := NewHashTable(100)
	if table == nil {
		t.Errorf("NewHashtable failed")
	}

	table.Set(1, &robot{1, "robot_1"})
	table.Set("1", &robot{1, "robot_2"})
	element := table.Get(1)
	if element == nil {
		t.Errorf("Hashtable.Get() failed")
	}

	if element.(*robot).Name != "robot_1" {
		t.Errorf("data error")
	}

	element = table.Get("1")
	if element == nil {
		t.Errorf("Hashtable.Get() failed")
	}

	if element.(*robot).Name != "robot_2" {
		t.Errorf("data error")
	}

	table.Del(1)
	element = table.Get(1)
	if element != nil {
		t.Errorf("Hashtable.Del() failed")
	}
}

var table *Hashtable

func BenchmarkHashtable_Set(b *testing.B) {
	table = NewHashTable(10000)
	robot := &robot{1, "robot"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		table.Set(i, robot)
	}
}

func BenchmarkHashtable_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table.Get(i)
	}
}

func BenchmarkHashtable_Del(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table.Del(i)
	}
}
