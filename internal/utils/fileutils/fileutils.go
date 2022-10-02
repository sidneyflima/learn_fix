package fileutils

import (
	"errors"
	"fmt"
	"os"
	"sidney/examples/learn_fix/internal/utils/stringutils"
)

func OpenFile(filename string) (*os.File, error) {
	if len(filename) == 0 {
		return nil, errors.New(fmt.Sprint("invalid file name", filename))
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetFormattedFullFileName(workDir, filePath string) string {

	if stringutils.EndsWith(workDir, byte(os.PathSeparator)) {
		workDir = stringutils.RemoveLastCharacter(workDir)
	}

	if stringutils.StartsWith(filePath, byte(os.PathSeparator)) {
		filePath = stringutils.RemoveFirstCharacter(filePath)
	}

	return fmt.Sprintf("%s%c%s", workDir, os.PathSeparator, filePath)
}
