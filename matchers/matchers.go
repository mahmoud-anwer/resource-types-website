package matchers

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type containSelectorMatcher struct {
	expected string
}

func (matcher containSelectorMatcher) Match(actual interface{}) (success bool, err error) {
	if doc, ok := actual.(*goquery.Document); ok {
		return doc.Find(matcher.expected).Length() != 0, nil
	} else {
		return false, fmt.Errorf("ContainSelector matcher expects a goquery.Document")
	}
}

func (matcher containSelectorMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\tdocument\nto contain the selector\n\t%s", matcher.expected)
}

func (matcher containSelectorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\tdocument\nnot to contain the selector\n\t%s", matcher.expected)
}

func ContainSelector(selector string) types.GomegaMatcher {
	return &containSelectorMatcher{
		expected: selector,
	}
}

type containSelectorWithTextMatcher struct {
	selector string
	matcher  types.GomegaMatcher
}

func ContainSelectorWithText(selector string, matcher types.GomegaMatcher) types.GomegaMatcher {
	return &containSelectorWithTextMatcher{selector, matcher}
}

func (c containSelectorWithTextMatcher) Match(actual interface{}) (success bool, err error) {
	if doc, ok := actual.(*goquery.Document); ok {
		return c.matcher.Match(doc.Find(c.selector).Text())
	} else {
		return false, fmt.Errorf("ContainSelectorWithText matcher expects a goquery.Document")
	}
}

func (c containSelectorWithTextMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\ttext of selector `%s`\nto match\n\t%s", c.selector, format.Object(c.matcher, 0))
}

func (c containSelectorWithTextMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\ttext of selector `%s`\nto not match\n\t%s", c.selector, format.Object(c.matcher, 0))
}
