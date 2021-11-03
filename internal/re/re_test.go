package re

import (
	"testing"

	"github.com/celicoo/ttest"
)

func TestCompile(t *testing.T) {
	ttest.
		NewTest(t).
		SetSubject(Compile).
		Run([]ttest.TestCase{})
}

func TestRE_Match(t *testing.T) {
	ttest.
		NewTest(t).
		SetSubject((*RE).Match).
		Run([]ttest.TestCase{})
}
