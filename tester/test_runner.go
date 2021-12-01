package tester

import (
	"reflect"
	"strings"
	"testing"
)

// Preparer 测试用例准备
type Preparer struct{}

// Setup 测试用例准备：前置操作
func (Preparer) Setup(t *testing.T) {
	t.Log("set up ...")
}

// Teardown 测试用例准备：后置清理
func (Preparer) Teardown(t *testing.T) {
	t.Log("teardown ...")
}

const (
	TestPrefix = "Test"
)

// RunTests 执行测试用例
func RunTests(t *testing.T, testSuites interface{}) {
	typ := reflect.TypeOf(testSuites)
	val := reflect.ValueOf(testSuites)
	setup := getTestFunc(val, "Setup")
	teardown := getTestFunc(val, "Teardown")
	for i := 0; i < typ.NumMethod(); i++ {
		name := typ.Method(i).Name
		if !strings.HasPrefix(name, TestPrefix) {
			continue
		}
		f := getTestFunc(val, name)
		t.Run(strings.TrimPrefix(name, TestPrefix), func(t *testing.T) {
			setup(t)
			defer teardown(t)
			f(t)
		})
	}
}

func getTestFunc(value reflect.Value, name string) func(t *testing.T) {
	if !value.IsValid() {
		return func(t *testing.T) {}
	}
	m := value.MethodByName(name)
	if !m.IsValid() {
		return func(t *testing.T) {}
	}
	f, ok := m.Interface().(func(t *testing.T))
	if ok {
		return f
	}
	return func(t *testing.T) {}
}
