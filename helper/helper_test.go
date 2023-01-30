package helper_test

import (
	"indonesian-heroes/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeFromInterface(t *testing.T) {
	testCase := []struct {
		Name        string
		Data        []interface{}
		Expectation int
	}{
		{
			Name:        "success string",
			Data:        []interface{}{"1234"},
			Expectation: 1234,
		}, {
			Name:        "success float64",
			Data:        []interface{}{float64(1234)},
			Expectation: 1234,
		},
	}

	for _, testCase := range testCase {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := helper.ChangeFromInterface(testCase.Data)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}
