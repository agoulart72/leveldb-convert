package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConvertLeveldbToJson(files []string, output string) error {

	for _, fileName := range files {

		outputFileName := getOutputName(fileName, output, ".json")

		db, err := leveldb.OpenFile(fileName, nil)
		if err != nil {
			return fmt.Errorf("error opening LevelDB : %w", err)
		}
		defer db.Close()

		jsonMap := make(map[string]any)

		iter := db.NewIterator(nil, nil)
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			if strings.HasPrefix(string(value), "{") &&
				strings.HasSuffix(string(value), "}") {
				jsonMap[string(key)], err = valueUnmarshal(value)
				if err != nil {
					fmt.Println("Warn : error unmarshall value : ", err)
					jsonMap[string(key)] = string(value)
				}
			} else {
				jsonMap[string(key)] = string(value)
			}
		}
		iter.Release()
		err = iter.Error()
		if err != nil {
			return fmt.Errorf("error iterating over LevelDB : %w", err)
		}

		dat, err := json.Marshal(jsonMap)
		if err != nil {
			return fmt.Errorf("error marshalling to Json : %w", err)
		}

		err = os.WriteFile(outputFileName, dat, 0644)
		if err != nil {
			return fmt.Errorf("error writing json file : %w", err)
		}

	}

	return nil
}

func valueUnmarshal(value []byte) (map[string]any, error) {

	var resp map[string]any

	err := json.Unmarshal(value, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
