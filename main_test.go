package main

import "testing"

func TestContact(t *testing.T) {
	ts := []struct {
		target   string
		expected int
	}{
		{"https://google.fr/", 200},
		{"https://freeboard.tech/post", 200},
		{"https://chocolatize.xyz/brantbjork", 404},
	}

	for _, tc := range ts {
		t.Run(tc.target, func(t *testing.T) {
			res, _ := contact(tc.target)
			if res != tc.expected {
				t.Errorf("got %d, expected %d", res, tc.expected)
			}
		})
	}
}
