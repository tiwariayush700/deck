package util

import (
	`testing`

	`github.com/stretchr/testify/assert`

	`deck/model`
)

func TestGetCode(t *testing.T) {

	code := GetCode(model.Spade, model.King)
	expectedCode := "KS"

	assert.Equal(t, expectedCode, code)

}
