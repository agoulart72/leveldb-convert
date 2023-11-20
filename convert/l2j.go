package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConvertLeveldbToJson(files []string) error {

	for _, fileName := range files {

		p := strings.Split(fileName, "/")
		path := strings.Join(p[:len(p)-1], "/")

		fn := strings.Split(p[len(p)-1], ".")
		outputFileName := path + "/" + strings.Join(fn[:len(fn)-1], ".") + ".json"

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
			jsonMap[string(key)] = value
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
