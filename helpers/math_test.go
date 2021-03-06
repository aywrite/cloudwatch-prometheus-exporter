package helpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var array []*float64

func init() {
	array = f64Ptrs(2, 1, 4, 3)
}

func f64Ptrs(values ...float64) []*float64 {
	fp := make([]*float64, len(values))
	for i := range fp {
		fp[i] = &values[i]
	}
	return fp
}

func tPtrs(times ...time.Time) []*time.Time {
	tp := make([]*time.Time, len(times))
	for i := range tp {
		tp[i] = &times[i]
	}
	return tp
}

func TestMax(t *testing.T) {
	got, err := Max(array)
	if err != nil {
		t.Errorf("Got err %s", err)
	}
	assert.Equal(t, 4.0, got)
}

func TestMin(t *testing.T) {
	got, err := Min(array)
	if err != nil {
		t.Errorf("Got err %s", err)
	}
	assert.Equal(t, 1.0, got)
}

func TestSum(t *testing.T) {
	got, err := Sum(array)
	if err != nil {
		t.Errorf("Got err %s", err)
	}
	assert.Equal(t, 10.0, got)
}

func TestAverage(t *testing.T) {
	got, err := Average(array)
	if err != nil {
		t.Errorf("Got err %s", err)
	}
	if got != 2.5 {
		t.Errorf("Average(%v) = %f; want 2.5", array, got)
	}
}

var newValuesTests = []struct {
	values    []*float64
	times     []*time.Time
	threshold time.Time
	expected  []*float64
}{
	{f64Ptrs(), tPtrs(), time.Now(), f64Ptrs()},                             // Empty input should give empty output
	{f64Ptrs(1), tPtrs(time.Now().Add(-time.Hour)), time.Now(), f64Ptrs()},  // threshold > time should give no output
	{f64Ptrs(1), tPtrs(time.Now()), time.Now().Add(-time.Hour), f64Ptrs(1)}, // threshold < time should not filter result
}

func TestNewValues(t *testing.T) {
	for _, v := range newValuesTests {
		got := NewValues(v.values, v.times, v.threshold)
		assert.Equal(t, v.expected, got)
	}
}
