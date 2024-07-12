package main

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				Job  string
			}{"Joe", "employed"},
			[]string{"Joe", "employed"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Jill", 42},
			[]string{"Jill"},
		},
		{
			"nested fields",
			Person{
				"Jack",
				Profile{100, "Antarctica"},
			},
			[]string{"Jack", "Antarctica"},
		},
		{
			"pointers to things",
			&Person{
				"Jack Pointer",
				Profile{10, "Africa"},
			},
			[]string{"Jack Pointer", "Africa"},
		},
		{
			"slices",
			[]Profile{
				Profile{11, "CityA"},
				Profile{12, "CityB"},
			},
			[]string{"CityA", "CityB"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			fn := func(input string) {
				got = append(got, input)
			}
			walk(test.Input, fn)
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got: %v, want: %v", got, test.ExpectedCalls)
			}
		})
	}
}
