package ignore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

function: Matches
purpose: check if a given path matches any of the given conditions

? Test cases
success
	| allowed
			| #1 => empty ignore list
			| #2 => filled ignore list
	| not allowed
			| #1 => std case

*/

func TestAllowed(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		t.Run("allowed", func(t *testing.T) {

			table := []struct {
				name  string
				i     Ignore
				input string
			}{
				{"empty ignore list", []string{}, "./test"},
				{"filled ignore list", []string{"./test/tmp", "node_modules", "**/*.go"}, "README.md"},
			}

			for _, test := range table {
				t.Run(test.name, func(t *testing.T) {
					assert.True(t, test.i.Allowed(".", test.input))
				})
			}
		})

		t.Run("not allowed", func(t *testing.T) {
			table := []struct {
				name    string
				i       Ignore
				rootDir string
				input   string
			}{
				{"dissallow all", []string{"*"}, ".", "ignore.go"},
				{"by file type", []string{"*.go"}, "../games", "game.go"},
				{"folder", []string{".git"}, "../../../../", ".git"},
			}

			for _, test := range table {
				t.Run(test.name, func(t *testing.T) {
					assert.False(t, test.i.Allowed(test.rootDir, test.input))
				})
			}
		})
	})
}
