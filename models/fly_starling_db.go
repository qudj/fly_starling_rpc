package models

import (
	"context"
	"github.com/qudj/fly_lib/tools"
	"github.com/qudj/fly_starling_rpc/config"
)

type StarlingProject struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

type StarlingGroup struct {
	Id          int64  `json:"id"`
	ProjectKey  string `json:"project_key"`
	GroupKey    string `json:"group_key"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

type StarlingOrigin struct {
	Id            int64  `json:"id"`
	ProjectKey    string `json:"project_key"`
	GroupKey      string `json:"group_key"`
	LangKey       string `json:"lang_key"`
	Lang          string `json:"lang"`
	OriginText string `json:"origin_text"`
	Status        int64  `json:"status"`
	UpdateTime    int64  `json:"update_time"`
	CreateTime    int64  `json:"create_time"`
}

type StarlingTranslation struct {
	Id         int64  `json:"id"`
	ProjectKey string `json:"project_key"`
	GroupKey   string `json:"group_key"`
	LangKey    string `json:"lang_key"`
	Lang       string `json:"lang"`
	TranslateText string `json:"translate_text"`
	Status     int64  `json:"status"`
	UpdateTime int64  `json:"update_time"`
	CreateTime int64  `json:"create_time"`
}

type StarlingHistoryLog struct {
	Id          int64  `json:"id"`
	Table       string `json:"table"`
	ObjectKey   string `json:"object_key"`
	ObjectType  string `json:"object_type"`
	OpId        string `json:"op_id"`
	ChangeData  string `json:"change_data"`
	HistoryData string `json:"history_data"`
	CreateTime  int64  `json:"create_time"`
}

func (StarlingProject) TableName() string {
	return "starling_project"
}

func (StarlingGroup) TableName() string {
	return "starling_group"
}

func (StarlingOrigin) TableName() string {
	return "starling_origin"
}

func (StarlingTranslation) TableName() string {
	return "starling_translation"
}

func (StarlingHistoryLog) TableName() string {
	return "starling_history_log"
}

func GetProjects(ctx context.Context, filter map[string]interface{}, offset, limit int, orderBy string) ([]*StarlingProject, int64, error) {
	var ret []*StarlingProject
	var count int64
	whereStr := "id > 0"
	whereArgs := make([]interface{}, 0)
	if v, ok := filter["project_key"]; ok {
		whereStr += " and project_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if v, ok := filter["project_name"]; ok {
		whereStr += " and project_name = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.StarlingReadDB.Table("starling_project").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		tools.LogCtxError(ctx, "GetProjects error=%v", err)
		return nil, 0, err
	}
	return ret, count, nil
}

func GetGroups(ctx context.Context, proKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*StarlingGroup, int64, error) {
	var ret []*StarlingGroup
	var count int64
	whereStr := "project_key = ?"
	whereArgs := []interface{}{proKey}
	if v, ok := filter["group_key"]; ok {
		whereStr += " and group_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if v, ok := filter["group_name"]; ok {
		whereStr += " and group_name = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.StarlingReadDB.Table("starling_group").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		tools.LogCtxError(ctx, "GetGroups error=%v", err)
		return nil, 0, err
	}
	return ret, count, nil
}

func GetStarlingOriginLgs(ctx context.Context, proKey, grKey string, filter map[string]interface{}, offset, limit int, orderBy string) ([]*StarlingOrigin, int64, error) {
	var ret []*StarlingOrigin
	var count int64
	whereStr := "project_key = ? and group_key = ?"
	whereArgs := []interface{}{proKey, grKey}
	if v, ok := filter["lang_key"]; ok {
		whereStr += " and lang_key = ?"
		whereArgs = append(whereArgs, v)
	}
	if err := config.StarlingReadDB.Table("starling_origin").WithContext(ctx).Where(whereStr, whereArgs...).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		tools.LogCtxError(ctx, "GetStarlingOriginLgs error=%v", err)
		return nil, 0, err
	}
	return ret, count, nil
}

func GetStarlingTransLgs(ctx context.Context, proKey, grKey, lgKey string, offset, limit int, orderBy string) ([]*StarlingTranslation, int64, error) {
	var ret []*StarlingTranslation
	var count int64
	if err := config.StarlingReadDB.Table("starling_translation").WithContext(ctx).
		Where("project_key = ? and group_key = ? and lang_key = ?", proKey, grKey, lgKey).Debug().Count(&count).
		Order(orderBy).Offset(offset).Limit(limit).Find(&ret).Error; err != nil {
		tools.LogCtxError(ctx, "GetStarlingTransLgs error=%v", err)
		return nil, 0, err
	}
	return ret, count, nil
}

func SaveProject(ctx context.Context, project *StarlingProject) error {
	if err := config.StarlingWriteDB.Debug().Save(project).Error; err != nil {
		tools.LogCtxError(ctx, "SaveProject error=%v", err)
		return err
	}
	return nil
}

func SaveGroup(ctx context.Context, group *StarlingGroup) error {
	if err := config.StarlingWriteDB.Debug().Save(group).Error; err != nil {
		tools.LogCtxError(ctx, "SaveGroup error=%v", err)
		return err
	}
	return nil
}

func SaveStarlingOriginLg(ctx context.Context, origin *StarlingOrigin) error {
	if err := config.StarlingWriteDB.Debug().Save(origin).Error; err != nil {
		tools.LogCtxError(ctx, "SaveStarlingOriginLg error=%v", err)
		return err
	}
	return nil
}

func SaveStarlingTransLg(ctx context.Context, trans *StarlingTranslation) error {
	if err := config.StarlingWriteDB.Debug().Save(trans).Error; err != nil {
		tools.LogCtxError(ctx, "SaveStarlingTransLg error=%v", err)
		return err
	}
	return nil
}

func SaveHistory(ctx context.Context, history *StarlingHistoryLog) error {
	if err := config.StarlingWriteDB.Debug().Save(history).Error; err != nil {
		tools.LogCtxError(ctx, "SaveHistory error=%v", err)
		return err
	}
	return nil
}

func GetStarlingTransLgsByKey(ctx context.Context, proKey, groKey, lang string, lgKeys []string) ([]*StarlingTranslation, error) {
	rets := make([]*StarlingTranslation, 0)
	if err := config.StarlingReadDB.WithContext(ctx).Debug().
		Where("project_key = ? and group_key = ? and lang_key in (?) and lang = ?", proKey, groKey, lgKeys, lang).Find(&rets).Error; err != nil {
		tools.LogCtxError(ctx, "GetStarlingTransLgsByKey error=%v", err)
		return nil, err
	}
	return rets, nil
}
