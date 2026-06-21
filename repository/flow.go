package repository

import (
	"fmt"

	"github.com/Inflowenger/dev-backend/models"
)

func FlowIndexByInt(i uint64) string {
	return fmt.Sprintf("%s:%d", FLOW_INDEX_PREFIX, i)
}
func FlowIndexByString(i string) string {
	return fmt.Sprintf("%s:%s", FLOW_INDEX_PREFIX, i)
}
func UpsertFlow(f *models.FlowRecord) error {
	if f.ID == "" {
		index, err := Seq()
		if err != nil {
			return err
		}
		f.ID = FlowIndexByInt(index)
	} else {
		db := GetBadgerDb(models.FlowRecord{})
		rec, err := db.Get(f.ID)
		if err != nil {
			return err
		}
		f.CreatedAt = rec.CreatedAt

	}

	return UpsertWithKeys([]string{f.ID}, f)
}

func GetFlowById(key string) *models.FlowRecord {
	db := GetBadgerDb(models.FlowRecord{})
	f, _ := db.Get(key)
	return f
}
