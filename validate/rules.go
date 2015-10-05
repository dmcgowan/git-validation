package validate

import "github.com/vbatts/git-validation/git"

var (
	// RegisteredRules are the standard validation to perform on git commits
	RegisteredRules = []Rule{}
)

// RegisterRule includes the Rule in the avaible set to use
func RegisterRule(vr Rule) {
	RegisteredRules = append(RegisteredRules, vr)
}

// Rule will operate over a provided git.CommitEntry, and return a result.
type Rule struct {
	Name        string // short name for reference in in the `-run=...` flag
	Description string // longer Description for readability
	Run         func(git.CommitEntry) Result
}

// Commit processes the given rules on the provided commit, and returns the result set.
func Commit(c git.CommitEntry, rules []Rule) Results {
	results := Results{}
	for _, r := range rules {
		results = append(results, r.Run(c))
	}
	return results
}

// Result is the result for a single validation of a commit.
type Result struct {
	CommitEntry git.CommitEntry
	Pass        bool
	Msg         string
}

// Results is a set of results. This is type makes it easy for the following function.
type Results []Result

// PassFail gives a quick over/under of passes and failures of the results in this set
func (vr Results) PassFail() (pass int, fail int) {
	for _, res := range vr {
		if res.Pass {
			pass++
		} else {
			fail++
		}
	}
	return pass, fail
}
