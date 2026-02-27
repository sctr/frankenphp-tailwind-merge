package twmerge

import "testing"

func TestTwJoin_Basic(t *testing.T) {
	result := TwJoin("px-2", "py-1")
	if result != "px-2 py-1" {
		t.Errorf("expected 'px-2 py-1', got %q", result)
	}
}

func TestTwJoin_Empty(t *testing.T) {
	result := TwJoin()
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestTwJoin_SkipsEmpty(t *testing.T) {
	result := TwJoin("px-2", "", "py-1", "")
	if result != "px-2 py-1" {
		t.Errorf("expected 'px-2 py-1', got %q", result)
	}
}

func TestTwJoin_SingleValue(t *testing.T) {
	result := TwJoin("px-2")
	if result != "px-2" {
		t.Errorf("expected 'px-2', got %q", result)
	}
}

func TestTwJoin_AllEmpty(t *testing.T) {
	result := TwJoin("", "", "")
	if result != "" {
		t.Errorf("expected empty, got %q", result)
	}
}
