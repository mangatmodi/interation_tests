package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCalculator(t *testing.T) {
	TestCase{
		ctx:   NewTestContext(t, "Performs divide by 0"),
		steps: []Step{TwoRandomNumbers, DivideStep, VerifyNoError},
	}.Run()

	TestCase{
		ctx:   NewTestContext(t, "Performs log with error"),
		steps: []Step{TwoRandomNumbers, LogStep, VerifyNoError},
	}.Run()
}

var TwoRandomNumbers Step = func(t TestContext) {
	a := rand.Intn(10)
	b := 0
	t.fixture["a"] = a
	t.fixture["b"] = b
}

var DivideStep Step = func(t TestContext) {
	a, ok := t.fixture["a"].(int)
	assert.True(t.t, ok)

	b, ok := t.fixture["b"].(int)
	assert.True(t.t, ok)

	t.fixture["result"], t.fixture["error"] = divide(a, b)
}

var LogStep Step = func(t TestContext) {
	a, ok := t.fixture["a"].(int)
	assert.True(t.t, ok)

	b, ok := t.fixture["b"].(int)
	assert.True(t.t, ok)

	t.fixture["result"], t.fixture["error"] = log(float64(a), float64(b))
}

var VerifyNoError Step = func(t TestContext) {
	_, ok := t.fixture["error"].(error)
	assert.False(t.t, ok)
}
