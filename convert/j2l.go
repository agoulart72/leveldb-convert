package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConvertJsonToLeveldb(files []string) error {

	for _, fileName := range files {

		p := strings.Split(fileName, "/")
		path := strings.Join(p[:len(p)-1], "/")

		fn := strings.Split(p[len(p)-1], ".")
		outputFileName := path + "/" + strings.Join(fn[:len(fn)-1], ".") + ".db"

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

		fmt.Println(fileName, " converted with success to ", outputFileName)
	}

	return nil
}