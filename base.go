package main

import (
	"testing"
)

type Step func(t TestContext)

type TestContext struct {
	name    string
	fixture map[string]interface{}
	t       *testing.T
}

func NewTestContext(t *testing.T, name string) TestContext {
	return TestContext{
		name:    name,
		fixture: map[string]interface{}{},
		t:       t,
	}
}

type TestCase struct {
	ctx   TestContext
	steps []Step
}

func (t TestCase) Run() {
	for _, j := range t.steps {
		j(t.ctx)
	}
}
