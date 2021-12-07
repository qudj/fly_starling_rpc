package service

import (
	"encoding/json"
	"fmt"
	"github.com/qudj/fly_starling_rpc/models"
	"time"
)

func SaveHistory(pre, cur interface{}, table, obKey, obType, opId string) error {
	historyByte, _ := json.Marshal(pre)
	changeByte, _ := json.Marshal(cur)
	curTime := time.Now().Unix()
	history := models.StarlingHistoryLog{
		Table:       table,
		ObjectKey:   obKey,
		ObjectType:  obType,
		OpId:        opId,
		ChangeData:  string(changeByte),
		HistoryData: string(historyByte),
		CreateTime:  curTime,
	}
	if err := models.SaveHistory(&history); err != nil {
		fmt.Println(fmt.Sprintf("save histroy error=%v", err))
		return err
	}
	return nil
}
