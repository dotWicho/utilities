package utilities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWithSlash(t *testing.T) {

	// Define some vars
	string1 := "/home/user/"
	string2 := "/usr/bin"

	// fire up
	result1 := EndsWithSlash(string1)
	result2 := EndsWithSlash(string2)

	// if result equals to expected?
	assert.Equal(t, string1, result1)
	assert.Equal(t, "/usr/bin/", result2)
}

func TestDelInitialSlash(t *testing.T) {

	// Define some vars
	string1 := "/home/user/"
	string2 := "usr/bin/"

	// fire up
	result1 := EndsWithSlash(string1)
	result2 := EndsWithSlash(string2)

	// if result equals to expected?
	assert.Equal(t, string1, result1)
	assert.Equal(t, "usr/bin/", result2)
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
