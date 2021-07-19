package main

import (
	"testing"
)

func TestRemoveWhiteSpace(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{name: "empty", str: "", want: ""},
		{name: "Multi spaces", str: ` On  my   home world   `, want: `On my home world`},
		{name: "Multi spaces with two leading", str: `  On  my   home world   `, want: `On my home world`},
		{name: "Three empty spaces string", str: `    `, want: ``},
		{name: "One empty spaces string", str: ` `, want: ``},
		{name: "Two empty spaces string", str: `  `, want: ``},
		{name: "Internal spaces", str: `On  my   home world`, want: `On my home world`},
		{name: "Leading and trailng spaces", str: ` On my home world `, want: `On my home world`},
		{name: "Not Unicode", str: `  Привет   Мир!    `, want: `Привет Мир!`},
	}
	for _, test := range tests {
		// deccr := fmt.Sprintf("%v %q %q", test.str, test.str)
		got := RemoveWhiteSpace(test.str)
		if got != test.want {
			t.Errorf("Test %s got=%q, want=%q", test.name, got, test.want)
		}
	}
}

func TestRemoveWhiteSpaceFields(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{name: "empty", str: "", want: ""},
		{name: "Multi spaces", str: ` On  my   home world   `, want: `On my home world`},
		{name: "Multi spaces with two leading", str: `  On  my   home world   `, want: `On my home world`},
		{name: "Three empty spaces string", str: `    `, want: ``},
		{name: "One empty spaces string", str: ` `, want: ``},
		{name: "Two empty spaces string", str: `  `, want: ``},
		{name: "Internal spaces", str: `On  my   home world`, want: `On my home world`},
		{name: "Leading and trailng spaces", str: ` On my home world `, want: `On my home world`},
		{name: "Not Unicode", str: `  Привет   Мир!    `, want: `Привет Мир!`},
		{name: "Random", str: ` On  my   home world`, want: `On my home world`},
	}
	for _, test := range tests {
		// deccr := fmt.Sprintf("%v %q %q", test.str, test.str)
		got := StandardizeSpacesFields(test.str)
		if got != test.want {
			t.Errorf("Test %s got=%q, want=%q", test.name, got, test.want)
		}
	}
}

func BenchmarkSpaceFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StandardizeSpacesFields(" On  my   home world   ")
	}
}

func BenchmarkSpace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RemoveWhiteSpace(" On  my   home world   ")
	}
}
