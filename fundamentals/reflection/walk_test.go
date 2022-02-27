package walk

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

	cases := []struct {
		Name          string
		Input         interface{}
		ExepctedCalls []string
	}{
		{
			"Struct with one string",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string",
			struct {
				Name     string
				Location string
			}{"Bob", "SXB"},
			[]string{"Bob", "SXB"},
		},
		{
			"Struct with non string value",
			struct {
				Name string
				Age  int
			}{"Bob", 23},
			[]string{"Bob"},
		},
		{
			"Nested fields",
			Person{
				"Bob",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Bob", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"John",
				Profile{
					33,
					"London",
				},
			},
			[]string{"John", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{25, "Strasbourg"},
			},
			[]string{"London", "Strasbourg"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{25, "Strasbourg"},
			},
			[]string{"London", "Strasbourg"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExepctedCalls) {
				t.Errorf("got %v, want %v", got, test.ExepctedCalls)
			}

		})
	}

	// handling map differently because order is not guaranteed, therefore hardcoding expected test will fail for no reason
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{33, "Bruxelles"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Bruxelles"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{22, "Berlin"}, Profile{33, "Bruxelles"}
		}

		var got []string
		want := []string{"Berlin", "Bruxelles"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contains %q but it didn't", haystack, needle)
	}
}
