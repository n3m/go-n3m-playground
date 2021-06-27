package dsheap

import (
	"testing"
)

func Test_MaxHeapInsert(t *testing.T) {
	tests := []int{25, 60, 48}

	mh := new(MaxHeap)
	for _, test := range tests {
		if err := mh.Insert(test); err != nil {
			t.Fatalf("[REAL] %+v", err.Error())
		}
		t.Logf("%+v", mh)
	}
	t.Fatalf("[FINAL] %+v", mh)

}

func Test_MaxHeapRemove(t *testing.T) {
	tests := []int{25, 60, 48, 47, 10, 56, 84, 96}

	mh := new(MaxHeap)
	for _, test := range tests {
		if err := mh.Insert(test); err != nil {
			t.Fatalf("[REAL] %+v", err.Error())
		}
		t.Logf("%+v", mh)

	}

	for i := 0; i < 3; i++ {
		topKey, err := mh.Top()
		if err != nil {
			t.Fatalf("[REAL] %+v", err.Error())
		}

		t.Logf("Top: %+v", topKey)
	}

	t.Fatalf("[FINAL] %+v", mh)
}
