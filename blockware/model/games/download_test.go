package games

import (
	"testing"

	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
)

func TestSetupDownload(t *testing.T) {
	testutil.ShortTest(t)
	gamesTestSetup()
	defer gamesTestTeardown()

	_, err := setupTestDownload()
	if err != nil {
		t.Error(err)
	}
}

func TestFindBlock(t *testing.T) {
	testutil.ShortTest(t)

	t.Cleanup(gamesTestTeardown)

	t.Run("success", func(t *testing.T) {
		gamesTestSetup()
		// TODO
	})

}
