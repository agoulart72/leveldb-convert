package convert

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConvertJsonToLeveldb(files []string, output string) error {

	for _, fileName := range files {

		outputFileName := getOutputName(fileName, output, ".db")

		dat, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}

		var jsonMap map[string]any
		if !json.Valid([]byte(dat)) {
			return fmt.Errorf("invalid JSON file : %s", fileName)
		}

		db, err := leveldb.OpenFile(outputFileName, nil)
		if err != nil {
			return fmt.Errorf("error opening LevelDB %w", err)
		}
		defer db.Close()

		json.Unmarshal([]byte(dat), &jsonMap)

		batch := new(leveldb.Batch)
		for key, value := range jsonMap {
			batch.Put([]byte(key), []byte(value.(string)))
		}
		err = db.Write(batch, nil)
		if err != nil {
			return fmt.Errorf("error writing LevelDB %w", err)
		}

		fmt.Println(fileName, " converted with success to ", outputFileName)
	}

	return nil
}
