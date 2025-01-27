package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	tests := []struct {
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
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"london", "england"},
			[]string{"london", "england"},
		},
		{

			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 18},
			[]string{"Chris"},
		},
		{

			"struct with nested string fields",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  18,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{

			"struct with pointers to values",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  18,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{18, "Cincinnati"},
				{24, "San Juan"},
			},
			[]string{"Cincinnati", "San Juan"},
		},
		{
			"arrays",
			[2]Profile{
				{18, "Cincinnati"},
				{24, "San Juan"},
			},
			[]string{"Cincinnati", "San Juan"},
		},
		{
			"maps",
			map[string]string{
				"foo":  "bar",
				"fizz": "buzz",
			},
			[]string{"bar", "buzz"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			want := test.ExpectedCalls
			var got []string

			Walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got '%q' want '%q'", got, want)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"dog": "bark",
		}

		var got []string
		Walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "bark")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, val := range haystack {
		if val == needle {
			contains = true
		}
	}

	if contains != true {
		t.Errorf("expected '%q' to contain '%q' but it didn't", haystack, needle)
	}

}
