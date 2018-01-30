package main

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	movies, err := Handler(Request{
		ID: 28,
	})
	fmt.Println(err)
	if err != nil {
		//fmt.Printf("%s", t)
		fmt.Println(err)
	}
	
	
	assert.IsType(t, nil, err)
	assert.NotEqual(t, 0, len(movies))
}