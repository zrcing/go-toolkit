package maps

import (
	"testing"
)

func TestCloneMap(t *testing.T) {
	source := map[string]interface{}{
		"a": map[string]interface{}{
			"c": 1,
			"d": "aa",
		},
		"e": 123,
	}
	dest, err := CloneMap(source)
	if err != nil {
		t.Error(err)
	}
	dest["b"] = 1
	if _, ok := source["b"]; ok {
		t.Error("Doesn't copy.")
	}
	if source["e"] != dest["e"] {
		t.Error("Copy error.")
	}
}
