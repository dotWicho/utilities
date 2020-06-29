package utilities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWithSlash(t *testing.T) {

	t.Run("return equal string if ends with a slash", func(t *testing.T) {
		// Define some vars
		var str string = "/home/user/"
		// fire up
		result := EndsWithSlash(str)
		// if result equals to expected?
		assert.Equal(t, str, result)
	})

	t.Run("return a string with a slash at end", func(t *testing.T) {
		// Define some vars
		var str string = "/home/user"
		var expected string = str + "/"
		// fire up
		result := EndsWithSlash(str)
		// if result equals to expected?
		assert.Equal(t, expected, result)
	})
}

func TestDelInitialSlash(t *testing.T) {

	t.Run("return equal string if first char is not slash", func(t *testing.T) {
		// Define some vars
		var str string = "home/user/"
		// fire up
		result := DelInitialSlash(str)
		// if result equals to expected?
		assert.Equal(t, str, result)
	})

	t.Run("return a string without a slash at first char", func(t *testing.T) {
		// Define some vars
		var str string = "/home/user/"
		var expected string = str[1:]
		// fire up
		result := DelInitialSlash(str)
		// if result equals to expected?
		assert.Equal(t, expected, result)
	})
}

func TestValidFileName(t *testing.T) {

	// Define some vars
	folderName := "/home/user/"
	fileName := "/file.ext"

	// Const expected
	const expected string = "/home/user/file.ext"

	// fire up
	result := ValidFileName(folderName, fileName)

	// if result equals to expected?
	assert.Equal(t, expected, result)
}

func TestName(t *testing.T) {

	// Define some vars
	fileName := "\\file - testing.ext"

	// Const expected
	const expected string = "file-testing.ext"

	// fire up
	result := Name(fileName)

	// if result equals to expected?
	assert.Equal(t, expected, result)
}

func TestBaseName(t *testing.T) {

	// Define some vars
	fileName := "\\file - testing.ext \\ .ok"

	// Const expected
	const expected string = "file-testing-ext-ok"

	// fire up
	result := BaseName(fileName)

	// if result equals to expected?
	assert.Equal(t, expected, result)
}

func TestCleanString(t *testing.T) {

	// Define some vars
	fileName := "\\file - testing.ext \\ .ok"

	// fire up
	result := CleanString(fileName, baseNameSeparators)

	// if result equals to expected?
	assert.Equal(t, "\\file-testingext-\\-ok", result)

	// remove this &_=+:
	fileName = "\\file_-=testing&before:after\\ +ok.txt"

	// fire up
	result = CleanString(fileName, separators)

	// if result equals to expected?
	assert.Equal(t, "\\file-testing-before-after\\-ok.txt", result)

	// remove this -
	fileName = "\\file-testing--_---before-after\\---ok.txt"

	// fire up
	result = CleanString(fileName, dashes)

	// if result equals to expected?
	assert.Equal(t, "\\filetestingbeforeafter\\ok.txt", result)

	// remove this -
	fileName = "\\file-testing--_---before-after\\---ok.txt"

	// fire up
	result = CleanString(fileName, dashes)

	// if result equals to expected?
	assert.Equal(t, "\\filetestingbeforeafter\\ok.txt", result)

	// remove this -
	fileName = "\\file-testing--_---before-after\\---ok.txt"

	// fire up
	result = CleanString(fileName, illegalName)

	// if result equals to expected?
	assert.Equal(t, "file-testing-before-after-ok.txt", result)
}
