package codes

import (
	"testing"
)

func TestRandomDecimalDigits(t *testing.T) {
	for i := 0; i < 20; i++ {
		s, err := NewRandomDecimalString(i)
		t.Logf("%v: %v, err=%v", i, s, err)
		if err != nil {
			t.Errorf("Expected err == nil but got %v", err)
		}
		if len(s) != i {
			t.Errorf("Expected string to be length %v but got %v", i, len(s))
		}
	}

	s1, err := NewRandomDecimalString(8)
	if err != nil {
		t.Error("Expected nil but got", err)
	}
	if len(s1) != 8 {
		t.Error("Expected len 8 but got", len(s1))
	}
	s2, err := NewRandomDecimalString(8)
	if err != nil {
		t.Error("Expected nil but got", err)
	}
	if len(s2) != 8 {
		t.Error("Expected len 8 but got", len(s1))
	}
	if s1 == s2 {
		t.Errorf("Expected s1 != s2 but got s1=%v and s2=%v", s1, s2)
	}
}
