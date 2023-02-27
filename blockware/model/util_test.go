package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	util "github.com/t02smith/part-iii-project/toolkit/util"
)

// * test utilities

func TestMain(m *testing.M) {
	util.InitLogger()

	os.Exit(m.Run())
}

func beforeEachAfterEach(t *testing.T) {
	t.Helper()

	// * setup
	testutil.ClearTmp("../")

	// * teardown
	t.Cleanup(func() {
		testutil.ClearTmp("../")
	})
}

// * tests

/*

function: CreateDirecotryIfNoExist
purpose: to create a new folder at the given location if it doesn't exist

? test cases:
success
	| #1 => folder is created successfully when it doesn't exist
	| #2 => folder is untouched if it already exists

failure
	| illegal arguments
			| #1 => Parent directories do not exist
			| #2 => File exists with the same name

*/

func TestCreateDirectoryIfNotExistCorrect(t *testing.T) {
	beforeEachAfterEach(t)

	t.Run("success", func(t *testing.T) {

		created := t.Run("folder created", func(t *testing.T) {
			err := CreateDirectoryIfNotExist("../test/data/tmp/hello-there")
			assert.Nil(t, err, err)

			f, err := os.Stat("../test/data/tmp/hello-there")
			assert.Nil(t, err, err)
			assert.True(t, f.IsDir(), "created data object is not a directory")
		})

		// ! create folder if previous test fails
		if !created {
			err := os.Mkdir("../test/data/tmp/hello-there", 0644)
			if err != nil {
				t.Fatal(err)
			}
		}

		t.Run("folder already exists", func(t *testing.T) {
			err := CreateDirectoryIfNotExist("../test/data/tmp/hello-there")
			assert.Nil(t, err, err)
		})

	})

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("parent directory not found", func(t *testing.T) {
				assert.NotNil(t,
					CreateDirectoryIfNotExist("./test/directory/doesnt/work"),
					"Will not create directory if parent dir doesnt exist")
			})

			t.Run("file exists with same name", func(t *testing.T) {
				assert.NotNil(t,
					CreateDirectoryIfNotExist("../test/data/tmp/.gitkeep"),
					"Should fail if a file exists with that name")
			})
		})
	})

}

/*

function: SetupToolkitEnvironment
purpose: create the folders used for storing data by this application

? test cases
success
	| #1 => Folder and sub folders are created
	| environment already exists
			| #1 => complete folder
			| #2 => environment incomplete

failure
	| illegal arguments
			| #1 => invalid directory specified in config

*/

func TestSetupToolkitEnvironment(t *testing.T) {
	beforeEachAfterEach(t)
	viper.Set("meta.directory", "../test/data/tmp/.toolkit")

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {
			t.Run("invalid input directory", func(t *testing.T) {
				viper.Set("meta.directory", "./fake/directory/shouldnt/exist")
				t.Cleanup(func() {
					viper.Set("meta.directory", "../test/data/tmp/.toolkit")
				})

				assert.NotNil(t, SetupToolkitEnvironment(), "should catch invalid root directory")
			})
		})
	})

	t.Run("success", func(t *testing.T) {
		beforeEachAfterEach(t)

		verifier := func() bool {
			f, err := os.Open(viper.GetString("meta.directory"))
			if err != nil {
				return false
			}
			defer f.Close()

			gamesFound, hashesFound := false, false
			dirs, err := f.Readdirnames(0)
			if err != nil {
				return false
			}

			for _, dir := range dirs {
				switch dir {
				case "games":
					gamesFound = true
				case "hashes":
					hashesFound = true
				}
			}

			return gamesFound && hashesFound
		}

		passed := t.Run("basic case", func(t *testing.T) {
			err := SetupToolkitEnvironment()
			assert.Nil(t, err, err)
			assert.True(t, verifier(), "invalid dir created")
		})

		// ! smoke
		if !passed {
			t.FailNow()
		}

		t.Run("env already exists", func(t *testing.T) {
			t.Run("is complete", func(t *testing.T) {
				err := SetupToolkitEnvironment()
				assert.Nil(t, err, err)
				assert.True(t, verifier(), "invalid dirs created")
			})

			t.Run("is incomplete", func(t *testing.T) {

				err := os.Remove(filepath.Join(viper.GetString("meta.directory"), "games"))
				assert.Nil(t, err, err)

				err = SetupToolkitEnvironment()
				assert.Nil(t, err, err)
				assert.True(t, verifier(), "invalid dirs created")

			})
		})

	})

	// run function to be tested
	err := SetupToolkitEnvironment()
	if err != nil {
		t.Error(err)
		return
	}

	// assert directories created
	_, err = os.Stat("../test/data/tmp/.toolkit")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = os.Stat("../test/data/tmp/.toolkit/hashes")
	if err != nil {
		t.Error(err)
		return
	}
}

/*

function: ZipDirectory
purpose: create a zip archive from a given directory

? Test cases
success
	| #1 => archive created successfully

failure
	| illegal arguments
			| #1 => invalid archive path
				 \ #1 => teardown of incomplete archive is completed
			| #2 => invalid output path

*/

func TestZipDirectory(t *testing.T) {
	beforeEachAfterEach(t)

	t.Run("failure", func(t *testing.T) {
		t.Run("illegal arguments", func(t *testing.T) {

			t.Run("archive path", func(t *testing.T) {
				err := ZipDirectory("./fake/directory/for/test", "../test/data/tmp/testdir.zip")
				assert.NotNil(t, err, "Fake direcotry not discovered")

				t.Run("teardown", func(t *testing.T) {
					if _, err := os.Stat("../test/data/tmp/testdir.zip"); !os.IsNotExist(err) {
						t.Error("teardown incomplete")
					}
				})
			})

			t.Run("output path", func(t *testing.T) {
				err := ZipDirectory("../test/data/testdir", "./fake/file/to/output/to.zip")
				assert.NotNil(t, err, "Invalid output location not caught")

			})

		})
	})

	t.Run("success", func(t *testing.T) {
		err := ZipDirectory("../test/data/testdir", "../test/data/tmp/testdir.zip")
		if err != nil {
			t.Fatal(err)
		}

		// ? archive created
		stat, err := os.Stat("../test/data/tmp/testdir.zip")
		assert.Nil(t, err, err)
		assert.Greater(t, stat.Size(), int64(0), "Data not written to archive")
	})

	testutil.ClearTmp("../")
}
