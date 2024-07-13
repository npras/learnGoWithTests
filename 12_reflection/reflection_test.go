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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it did not", haystack, needle)
	}
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
		{
			"arrays",
			[2]Profile{
				Profile{77, "ArrCityX"},
				Profile{88, "ArrCityY"},
			},
			[]string{"ArrCityX", "ArrCityY"},
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

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("channels", func(t *testing.T) {
		ch := make(chan Profile)
		go func() {
			ch <- Profile{1, "Baby1"}
			ch <- Profile{2, "Baby2"}
			close(ch)
		}()
		var got []string
		want := []string{"Baby1", "Baby2"}
		walk(ch, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("function", func(t *testing.T) {
		aFn := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}
		var got []string
		want := []string{"Berlin", "Katowice"}
		walk(aFn, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
