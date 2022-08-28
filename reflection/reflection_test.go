package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{Name: "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Chris", City: "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{Name: "Chris", Age: 30},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Nested fields",
			Input: struct {
				Name    string
				Age     int
				Contact struct {
					Zip     int
					City    string
					Country string
				}
			}{
				Name: "Chris",
				Age:  30,
				Contact: struct {
					Zip     int
					City    string
					Country string
				}{
					Zip:     20001,
					City:    "DC",
					Country: "US",
				},
			},
			ExpectedCalls: []string{"Chris", "DC", "US"},
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
}
