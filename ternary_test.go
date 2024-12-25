package ternary

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type T any
type R any

func TestTernary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestIf")
}

var _ = DescribeTable(
	"TestIf",
	func(condition bool, elemTrue, elemFalse, expectedResult T) {
		result := If(condition, elemTrue, elemFalse)
		Expect(result).To(BeEquivalentTo(expectedResult))
	},
	Entry("String True", true, "foo", "bar", "foo"),
	Entry("String False", false, "foo", "bar", "bar"),
	Entry("Int True", true, 1, 2, 1),
	Entry("Int False", false, 1, 2, 2),
	Entry("Float True", true, 1.1, 2.2, 1.1),
	Entry("Float False", false, 1.1, 2.2, 2.2),
	Entry("Bool True", true, true, false, true),
	Entry("Bool False", false, true, false, false),
	Entry("Struct True", true, struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "foo"}),
	Entry("Struct False", false, struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "bar"}),
	Entry("Slice True", true, []string{"foo"}, []string{"bar"}, []string{"foo"}),
	Entry("Slice False", false, []string{"foo"}, []string{"bar"}, []string{"bar"}),
	Entry("Map True", true, map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "foo"}),
	Entry("Map False", false, map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "bar"}),
	Entry("Interface True", true, interface{}("foo"), interface{}("bar"), interface{}("foo")),
	Entry("Interface False", false, interface{}("foo"), interface{}("bar"), interface{}("bar")),
)

var _ = DescribeTable(
	"TestIfFunc",
	func(condition bool, elemTrue, elemFalse T, expectedResult R) {
		testFunc := func(elem T) R {
			return elem
		}
		result := IfFunc(condition, elemTrue, elemFalse, testFunc)
		Expect(result).To(BeEquivalentTo(expectedResult))
	},
	Entry("String True", true, "foo", "bar", "foo"),
	Entry("String False", false, "foo", "bar", "bar"),
	Entry("Int True", true, 1, 2, 1),
	Entry("Int False", false, 1, 2, 2),
	Entry("Float True", true, 1.1, 2.2, 1.1),
	Entry("Float False", false, 1.1, 2.2, 2.2),
	Entry("Bool True", true, true, false, true),
	Entry("Bool False", false, true, false, false),
	Entry("Struct True", true, struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "foo"}),
	Entry("Struct False", false, struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "bar"}),
	Entry("Slice True", true, []string{"foo"}, []string{"bar"}, []string{"foo"}),
	Entry("Slice False", false, []string{"foo"}, []string{"bar"}, []string{"bar"}),
	Entry("Map True", true, map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "foo"}),
	Entry("Map False", false, map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "bar"}),
	Entry("Interface True", true, interface{}("foo"), interface{}("bar"), interface{}("foo")),
	Entry("Interface False", false, interface{}("foo"), interface{}("bar"), interface{}("bar")),
)

var _ = DescribeTable(
	"TestThird",
	func(value, third string, elemTrue, elemFalse, elemThird, expectedResult T) {
		result, err := Third(value, third, elemTrue, elemFalse, elemThird, "")
		Expect(result).To(BeEquivalentTo(expectedResult))
		Expect(err).To(BeNil())
	},
	Entry("String True", "true", "third", "foo", "bar", "baz", "foo"),
	Entry("String False", "false", "third", "foo", "bar", "baz", "bar"),
	Entry("String Third 1", "third", "third", "foo", "bar", "baz", "baz"),
	Entry("String Third 2", "strict", "strict", "foo", "bar", "baz", "baz"),
	Entry("Int True", "true", "third", 1, 2, 3, 1),
	Entry("Int False", "false", "third", 1, 2, 3, 2),
	Entry("Int Third 1", "third", "third", 1, 2, 3, 3),
	Entry("Int Third 2", "strict", "strict", 1, 2, 3, 3),
	Entry("Float True", "true", "third", 1.1, 2.2, 3.3, 1.1),
	Entry("Float False", "false", "third", 1.1, 2.2, 3.3, 2.2),
	Entry("Float Third 1", "third", "third", 1.1, 2.2, 3.3, 3.3),
	Entry("Float Third 2", "strict", "strict", 1.1, 2.2, 3.3, 3.3),
	Entry("Struct True", "true", "third", struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "baz"}, struct{ foo string }{foo: "foo"}),
	Entry("Struct False", "false", "third", struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "baz"}, struct{ foo string }{foo: "bar"}),
	Entry("Struct Third 1", "third", "third", struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "baz"}, struct{ foo string }{foo: "baz"}),
	Entry("Struct Third 2", "strict", "strict", struct{ foo string }{foo: "foo"}, struct{ foo string }{foo: "bar"},
		struct{ foo string }{foo: "baz"}, struct{ foo string }{foo: "baz"}),
	Entry("Slice True", "true", "third", []string{"foo"}, []string{"bar"}, []string{"baz"}, []string{"foo"}),
	Entry("Slice False", "false", "third", []string{"foo"}, []string{"bar"}, []string{"baz"}, []string{"bar"}),
	Entry("Slice Third 1", "third", "third", []string{"foo"}, []string{"bar"}, []string{"baz"}, []string{"baz"}),
	Entry("Slice Third 2", "strict", "strict", []string{"foo"}, []string{"bar"}, []string{"baz"}, []string{"baz"}),
	Entry("Map True", "true", "third", map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "baz"}, map[string]string{"foo": "foo"}),
	Entry("Map False", "false", "third", map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "baz"}, map[string]string{"foo": "bar"}),
	Entry("Map Third 1", "third", "third", map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "baz"}, map[string]string{"foo": "baz"}),
	Entry("Map Third 2", "strict", "strict", map[string]string{"foo": "foo"}, map[string]string{"foo": "bar"},
		map[string]string{"foo": "baz"}, map[string]string{"foo": "baz"}),
	Entry("Interface True", "true", "third", interface{}("foo"), interface{}("bar"), interface{}("baz"), interface{}("foo")),
	Entry("Interface False", "false", "third", interface{}("foo"), interface{}("bar"), interface{}("baz"), interface{}("bar")),
	Entry("Interface Third 1", "third", "third", interface{}("foo"), interface{}("bar"), interface{}("baz"), interface{}("baz")),
	Entry("Interface Third 2", "strict", "strict", interface{}("foo"), interface{}("bar"), interface{}("baz"), interface{}("baz")),
)

var _ = DescribeTable(
	"ErrTestThird",
	func(value string, elemTrue, elemFalse, elemThird T, customError, expectedError string) {
		result, err := Third(value, "third", elemTrue, elemFalse, elemThird, customError)
		Expect(result).To(BeNil())
		Expect(err).To(MatchError(expectedError))
	},
	Entry("Error Default", "enabled", "foo", "bar", "baz", "", "value does not match options"),
	Entry("Error Custom", "enabled", "foo", "bar", "baz", "custom error", "custom error"),
)
