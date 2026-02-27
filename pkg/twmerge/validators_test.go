package twmerge

import "testing"

func TestIsFraction(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1/2", true},
		{"3/4", true},
		{"1.5/2", true},
		{"100/100", true},
		{"1", false},
		{"px", false},
		{"", false},
		{"1/", false},
		{"/2", false},
	}
	for _, tt := range tests {
		if got := IsFraction(tt.input); got != tt.want {
			t.Errorf("IsFraction(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"1.5", true},
		{"0", true},
		{"-1", true},
		{"-1.5", true},
		{"1e5", true},
		{"px", false},
		{"", false},
		{"abc", false},
	}
	for _, tt := range tests {
		if got := IsNumber(tt.input); got != tt.want {
			t.Errorf("IsNumber(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsInteger(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"0", true},
		{"-1", true},
		{"100", true},
		{"1.5", false},
		{"px", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsInteger(tt.input); got != tt.want {
			t.Errorf("IsInteger(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsPercent(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"50%", true},
		{"100%", true},
		{"0%", true},
		{"3.5%", true},
		{"50", false},
		{"%", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsPercent(tt.input); got != tt.want {
			t.Errorf("IsPercent(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsTshirtSize(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"sm", true},
		{"md", true},
		{"lg", true},
		{"xl", true},
		{"xs", true},
		{"2xl", true},
		{"2.5xl", true},
		{"10xl", true},
		{"abc", false},
		{"", false},
		{"xxl", false},
	}
	for _, tt := range tests {
		if got := IsTshirtSize(tt.input); got != tt.want {
			t.Errorf("IsTshirtSize(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryValue(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[#B91C1C]", true},
		{"[10px]", true},
		{"[length:10px]", true},
		{"[color:red]", true},
		{"red-500", false},
		{"(--var)", false},
		{"", false},
		{"[]", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryValue(tt.input); got != tt.want {
			t.Errorf("IsArbitraryValue(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryLength(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[10px]", true},
		{"[length:10px]", true},
		{"[2rem]", true},
		{"[calc(100%-4px)]", true},
		{"[#B91C1C]", false},
		{"[color:red]", false},
		{"10px", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryLength(tt.input); got != tt.want {
			t.Errorf("IsArbitraryLength(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[1.5]", true},
		{"[number:2]", true},
		{"[0]", true},
		{"[px]", false},
		{"1.5", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryNumber(tt.input); got != tt.want {
			t.Errorf("IsArbitraryNumber(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryPosition(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[position:absolute]", true},
		{"[percentage:50%]", true},
		{"[absolute]", false},
		{"absolute", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryPosition(tt.input); got != tt.want {
			t.Errorf("IsArbitraryPosition(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryImage(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[url(image.png)]", true},
		{"[image:url(image.png)]", true},
		{"[linear-gradient(red,blue)]", true},
		{"[#B91C1C]", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryImage(tt.input); got != tt.want {
			t.Errorf("IsArbitraryImage(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryShadow(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[0_0_10px_red]", true},
		{"[shadow:0_0_10px_red]", true},
		{"[inset_0_0_10px_red]", true},
		{"[#B91C1C]", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryShadow(tt.input); got != tt.want {
			t.Errorf("IsArbitraryShadow(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryVariable(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"(--my-var)", true},
		{"(length:--var)", true},
		{"[--my-var]", false},
		{"--my-var", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryVariable(tt.input); got != tt.want {
			t.Errorf("IsArbitraryVariable(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryVariableLength(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"(length:--my-var)", true},
		{"(--my-var)", false},
		{"(number:--var)", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryVariableLength(tt.input); got != tt.want {
			t.Errorf("IsArbitraryVariableLength(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryVariableShadow(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"(shadow:--my-var)", true},
		{"(--my-var)", true}, // shouldMatchNoLabel=true
		{"(number:--var)", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryVariableShadow(tt.input); got != tt.want {
			t.Errorf("IsArbitraryVariableShadow(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsArbitraryVariableWeight(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"(weight:--my-var)", true},
		{"(number:--my-var)", true},
		{"(--my-var)", true}, // shouldMatchNoLabel=true
		{"(length:--var)", false},
	}
	for _, tt := range tests {
		if got := IsArbitraryVariableWeight(tt.input); got != tt.want {
			t.Errorf("IsArbitraryVariableWeight(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsAny(t *testing.T) {
	if !IsAny("anything") {
		t.Error("IsAny should always return true")
	}
	if !IsAny("") {
		t.Error("IsAny should return true for empty string")
	}
}

func TestIsAnyNonArbitrary(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"red-500", true},
		{"block", true},
		{"[10px]", false},
		{"(--var)", false},
	}
	for _, tt := range tests {
		if got := IsAnyNonArbitrary(tt.input); got != tt.want {
			t.Errorf("IsAnyNonArbitrary(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
