package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

//Here we combine steps to create test suite
func TestCalculator(t *testing.T) {
	TestCase{
		ctx:   NewTestContext(t, "Performs divide by 0"),
		steps: []Step{TwoRandomNumbers, SetB(0), DivideStep, VerifyNoError},
	}.Run()

	TestCase{
		ctx:   NewTestContext(t, "Performs log with error"),
		steps: []Step{TwoRandomNumbers, LogStep, VerifyNoError},
	}.Run()
}

var SetB StepWithArgs = func(arg interface{}) Step {
	b := arg.(int)
	return func(t TestContext) {
		t.fixture["b"] = b
	}
}

//Define steps.
var TwoRandomNumbers Step = func(t TestContext) {
	t.fixture["a"] = rand.Intn(10)
	t.fixture["b"] = rand.Intn(10)
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
