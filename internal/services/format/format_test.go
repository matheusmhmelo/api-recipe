package format

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFormat_Format(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
		err  error
	}{
		{
			"[format] Normal ingredients",
			"onions,garlic",
			[]string{"onions","garlic"},
			nil,
		},
		{
			"[format] Special characters ingredients",
			"onions_,^garlic=",
			[]string{"onions","garlic"},
			nil,
		},
		{
			"[format] Limit of ingredients reached",
			"onions,garlic,pepper,salsa",
			[]string(nil),
			errors.New("ingredient limit reached, choose only 3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.s)
			assert.Equal(t, got, tt.want)

			if tt.err != nil {
				if err == nil {
					t.Errorf("expected error: %v", tt.err.Error())
				} else if !strings.Contains(err.Error(), tt.err.Error()) {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}


func TestFormat_FormatIngredients(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			"[formatIngredients] Ingredients from API",
			"onions, garlic, pepper, salsa",
			[]string{"garlic","onions","pepper","salsa"},
		},
		{
			"[formatIngredients] Ingredients without format of API",
			"onions,garlic,pepper,salsa",
			[]string{"onions,garlic,pepper,salsa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatIngredients(tt.s)
			assert.Equal(t, got, tt.want)
		})
	}
}