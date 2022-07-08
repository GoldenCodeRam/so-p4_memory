package utils

import (
	"errors"
	"so-p4_memory/src/view/lang"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func ExtractIntFromEntry(entry *gtk.Entry) (int, error) {
	text, err := ExtractTextFromEntry(entry)
	if err != nil {
		return 0, err
	} else {
		number, err := strconv.Atoi(text)
		if err != nil {
			return 0, errors.New(lang.ERROR_INVALID_NUMBER)
		} else {
			return number, nil
		}
	}
}

func ExtractTextFromEntry(entry *gtk.Entry) (string, error) {
	text, err := entry.GetText()
	if err != nil {
		return "", errors.New(lang.ERROR_EMPTY_TEXT)
	} else {
		return text, nil
	}
}
