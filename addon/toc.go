package addon

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Not all fields implemented
// https://wowwiki.fandom.com/wiki/TOC_format
type TOC struct {
	Interface string
	Title     string
	Notes     string
	Author    string
	Version   string
	// defaultState string
	// savedVariables string
}

func tocReader(r io.Reader) TOC {
	scanner := bufio.NewScanner(r)
	var t = TOC{}
	// tocPrefix := "##"
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		if hasTag(line) {
			line := line[2:]
			if sep := strings.Index(line, ":"); sep != -1 {
				key := strings.TrimSpace(line[:sep])
				value := strings.TrimSpace(line[(sep + 1):])
				if err := t.setField(key, value); err != nil {
					fmt.Println(err.Error())
				}
			}
		}

	}
	return t
}

func hasTag(s string) bool {
	if len(s) > 2 {
		if s[:2] == "##" {
			return true
		}
	}
	return false

}

func (c *TOC) setField(key, value string) error {
	structVal := reflect.ValueOf(c).Elem()
	fieldVal := structVal.FieldByName(key)

	if !fieldVal.IsValid() {
		return fmt.Errorf("field %s not implemented", key)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("Cannot set field value %s", key)
	}

	// reflect.Value.Set takes reflect.Value
	val := reflect.ValueOf(value)
	fieldVal.Set(val)

	return nil
}
