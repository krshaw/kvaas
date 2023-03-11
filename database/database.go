package database

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

var (
	// keep seperate indices for different key types
	StringIndex map[string]int
	IntIndex    map[int64]int

	indexByType = map[byte]interface{}{
		0: IntIndex,
		1: StringIndex,
	}
)

func Get(key string) ([]byte, error) {
	f, err := os.Open("/var/lib/kvaas/data")
	if err != nil {
		return nil, nil
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		// parse json
		var entry map[string]interface{}
		json.Unmarshal(line, &entry)
		if val, ok := entry[key]; ok {
			return json.Marshal(val)
		}
	}
	return nil, nil
}

func Create(pair []byte) error {
	f, err := os.OpenFile("/var/lib/kvaas/data", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	valid := json.Valid(pair)
	if !valid {
		return errors.New("Pair not valid json")
	}
	// write type information to the database
	// figure out if int64 or string
	var keyType byte = 0
	if err = json.Unmarshal(pair, &map[int]interface{}{}); err != nil {
		// try stringEntry instead
		if err = json.Unmarshal(pair, &map[string]interface{}{}); err != nil {
			return errors.New("Key not string or integer")
		} else {
			keyType = 1
		}
	}
	pair = append(pair, keyType)
	pair = append(pair, '\n')
	if _, err := f.Write(pair); err != nil {
		return err
	}
	return nil
}
