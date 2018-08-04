package subtest

import (
	"context"
	"testing"

	"gotest.tools/assert"
)

func TestRunCallsCleanup(t *testing.T) {
	calls := []int{}
	var ctx context.Context
	Run(t, "test-run-cleanup", func(t TestContext) {
		cleanup := func(n int) func() {
			return func() {
				calls = append(calls, n)
			}
		}
		ctx = t.Ctx()
		t.AddCleanup(cleanup(2))
		t.AddCleanup(cleanup(1))
		t.AddCleanup(cleanup(0))
	})
	assert.DeepEqual(t, calls, []int{0, 1, 2})
	assert.Equal(t, ctx.Err(), context.Canceled)
}
