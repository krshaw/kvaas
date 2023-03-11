package database

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

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
	pair = append(pair, '\n')
	valid := json.Valid(pair)
	if !valid {
		return errors.New("Pair not valid json")
	}
	if _, err := f.Write(pair); err != nil {
		return err
	}
	return nil
}
