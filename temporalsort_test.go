package temporalsort_test

import (
	"reflect"
	"testing"

	"github.com/willbarkoff/temporalsort"
)

func TestSort(t *testing.T) {
	arr1 := []int{5, 8, 2, 9}
	arr1Sorted := []int{2, 5, 8, 9}
	new := temporalsort.Sort(arr1)

	if reflect.DeepEqual(new, arr1Sorted) {
		t.Errorf("Got %v but expected %v\n", new, arr1)
	}
}
