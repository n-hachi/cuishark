package utils

import "testing"

func TestCutStringTail(t *testing.T) {
	str := "SampleTest"
	cmp := "Sample"

	if result := CutStringTail(str, 6); result != cmp {
		t.Errorf("str should be %s, but actually %s\n", cmp, result)
	}

	cmp = str
	if result := CutStringTail(str, len(str)); result != cmp {
		t.Errorf("str should be %s, but actually %s\n", cmp, result)
	}

	cmp = str
	if result := CutStringTail(str, len(str)+1); result != cmp {
		t.Errorf("str should be %s, but actually %s\n", cmp, result)
	}
}
