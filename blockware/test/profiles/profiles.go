package profiles

import (
	"errors"

	profile1 "github.com/t02smith/part-iii-project/toolkit/test/profiles/1"
	profile2 "github.com/t02smith/part-iii-project/toolkit/test/profiles/2"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profiles are cmd based instances of this application
and will run a predefined set of steps to mimic a
typically peer whilst automating certain aspects.

Each profile has a specific purpose that is explained
in the file of each one and its own custom config file
to be parsed initially.

*/

// run a given profile by its ID number
func RunProfile(profileNumber uint) error {
	util.Logger.Infof("Profile number %d selected. Attempting to start profile", profileNumber)

	switch profileNumber {
	case 1:
		profile1.Run()
	case 2:
		profile2.Run()
	default:
		return errors.New("unknown profile")
	}

	return nil
}
