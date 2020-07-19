package utilities

import (
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Remove all other unrecognised characters apart from
var (
	baseNameSeparators = regexp.MustCompile(`[./]`)
	separators         = regexp.MustCompile(`[ &_=+:]`)
	dashes             = regexp.MustCompile(`[\-]+`)
	illegalName        = regexp.MustCompile(`[^[:alnum:]-.]`)
)

// EndsWithSlash just check if a path ends with /, if not add
func EndsWithSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

// DelInitialSlash just delete (if exists) a initial /
func DelInitialSlash(path string) string {
	if path[0] == '/' {
		return path[1:]
	}
	return path
}

// Name makes a string safe to use in a file name by first finding the path basename, then replacing non-ascii characters.
func Name(str string) string {
	// Start with a Clean Base
	fileName := path.Clean(path.Base(str))
	// Remove illegal characters for names, replacing some common separators with -
	fileName = CleanString(fileName, illegalName)
	// NB this may be of length 0, caller must check
	return fileName
}

// CleanString replaces separators with - and removes characters listed in the regexp provided from string.
// Accents, spaces, and all characters not in A-Za-z0-9 are replaced.
func CleanString(str string, r *regexp.Regexp) string {
	// Remove any trailing space to avoid ending on -
	str = strings.Trim(str, " ")
	// Replace certain joining characters with a dash
	str = separators.ReplaceAllString(str, "-")
	// Remove all other unrecognised characters - NB we do allow any printable characters
	str = r.ReplaceAllString(str, "")
	// Remove any multiple dashes caused by replacements above
	str = dashes.ReplaceAllString(str, "-")

	return str
}

// BaseName makes a string safe to use in a file name, producing a sanitized basename replacing . or / with -.
// No attempt is made to normalise a path or normalise case.
func BaseName(str string) string {
	// Replace certain joining characters with a dash
	baseName := baseNameSeparators.ReplaceAllString(str, "-")
	// Remove illegal characters for names, replacing some common separators with -
	baseName = CleanString(baseName, illegalName)
	// NB this may be of length 0, caller must check
	return baseName
}

// FileNameISO8601 returns a valid fileName with ISO8606 timestamp on it
func FileNameISO8601(fileName string) string {
	// Stats with a clean base fileName
	fileName = Name(fileName)
	// Get file extension
	var extension = filepath.Ext(fileName)
	// Strip file extension
	var name = FilenameWithoutExtension(fileName)
	// Return name + Separator + ISO8601 Timestamp + extension
	return BaseName(name+"_"+time.Now().Format(time.RFC3339)) + extension
}

// ValidFileName returns a valid fileName without not-ascii characters
func ValidFileName(folderName, fileName string) string {

	return path.Join(folderName, Name(DelInitialSlash(fileName)))
}

// FilenameWithoutExtension returns a valid fileName without extension
func FilenameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
