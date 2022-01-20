package test

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("example.Add(int, int)  got %q, wanted %q", got, want)
	} else {
		t.Log("example.Add(int, int) Success")
	}
}

func TestIsFindName(t *testing.T) {
	got1 := IsFindName("Wayne",  []string{"Andy", "Jack", "Daniel", "Wayne"})
	want1 := true
	got2 := IsFindName("Arthur", []string{"Andy", "Jack", "Daniel", "Wayne"})
	want2 := false

	if got1 != want1 {
		t.Errorf("example.IsFindName(string, []string)  got %v, wanted %v", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("example.IsFindName(string, []string)  got %v, wanted %v", got2, want2)
	}
}