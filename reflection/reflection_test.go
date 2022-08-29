package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name     string
	Nickname string
	Age      int
	Contact  ContactInfo
}

type ContactInfo struct {
	Zip     int
	City    string
	Country string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: Person{
				Name: "John",
			},
			ExpectedCalls: []string{"John"},
		},
		{
			Name: "Struct with two string fields",
			Input: Person{
				Name:     "John",
				Nickname: "Johny",
			},
			ExpectedCalls: []string{"John", "Johny"},
		},
		{
			Name: "Struct with non string field",
			Input: Person{
				Name:     "John",
				Nickname: "Johny",
				Age:      99,
			},
			ExpectedCalls: []string{"John", "Johny"},
		},
		{
			Name: "Nested fields",
			Input: Person{
				Name:     "John",
				Nickname: "Johny",
				Age:      99,
				Contact: ContactInfo{
					Zip:     20001,
					City:    "DC",
					Country: "US",
				},
			},
			ExpectedCalls: []string{"John", "Johny", "DC", "US"},
		},
		{
			Name: "Pointer to struct",
			Input: &Person{
				Name:     "John",
				Nickname: "Johny",
				Age:      99,
				Contact: ContactInfo{
					Zip:     20001,
					City:    "DC",
					Country: "US",
				},
			},
			ExpectedCalls: []string{"John", "Johny", "DC", "US"},
		},
		{
			Name: "Slices",
			Input: []Person{
				{
					Name:     "John",
					Nickname: "Johny",
					Age:      99,
					Contact: ContactInfo{
						Zip:     20001,
						City:    "DC",
						Country: "US",
					},
				},
				{
					Name:     "Reg",
					Nickname: "Reggie",
					Age:      99,
					Contact: ContactInfo{
						Zip:     20001,
						City:    "DC",
						Country: "US",
					},
				},
			},
			ExpectedCalls: []string{"John", "Johny", "DC", "US", "Reg", "Reggie", "DC", "US"},
		},
		{
			Name: "Arrays",
			Input: [2]Person{
				{
					Name:     "John",
					Nickname: "Johny",
					Age:      99,
					Contact: ContactInfo{
						Zip:     20001,
						City:    "DC",
						Country: "US",
					},
				},
				{
					Name:     "Reg",
					Nickname: "Reggie",
					Age:      99,
					Contact: ContactInfo{
						Zip:     20001,
						City:    "DC",
						Country: "US",
					},
				},
			},
			ExpectedCalls: []string{"John", "Johny", "DC", "US", "Reg", "Reggie", "DC", "US"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Baz": "Boz",
			"Foo": "Bar",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan string)

		go func() {
			aChannel <- "foo"
			aChannel <- "bar"
			close(aChannel)
		}()

		var got []string

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "foo")
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (string, string) {
			return "foo", "bar"
		}
		var got []string

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "foo")
		assertContains(t, got, "bar")
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
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
