package maps

import (
	"bytes"
	"encoding/gob"
)

func CloneMap(source map[string]interface{}) (dest map[string]interface{}, err error) {
	var buffer bytes.Buffer
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
	encoder := gob.NewEncoder(&buffer)
	decoder := gob.NewDecoder(&buffer)
	err = encoder.Encode(source)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(&dest)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
