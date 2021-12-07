package handler

import (
	"context"
	"errors"
	"github.com/qudj/fly_starling_rpc/config"
	"github.com/qudj/fly_starling_rpc/models"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
	"github.com/qudj/fly_starling_rpc/service"
	"gorm.io/gorm"
	"time"
)

func SaveProject(ctx context.Context, req *servbp.SaveProjectRequest) error {
	if req.Project == nil {
		return errors.New("param error")
	}
	pre := &models.StarlingProject{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).Debug().Where("project_key = ?", req.Project.ProjectKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurProject(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveProject(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurProject(pre *models.StarlingProject, req *servbp.SaveProjectRequest) (*models.StarlingProject, error) {
	if req.Project.ProjectKey == "" {
		return nil, errors.New("project_key need")
	}
	cur := &models.StarlingProject{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.Project.ProjectName == "" {
			return nil, errors.New("add project need project_key")
		}
		if req.Project.Description == "" {
			return nil, errors.New("add project need description")
		}
		if req.Project.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.Project.ProjectKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.Project.ProjectName != "" {
		cur.ProjectName = req.Project.ProjectName
	}
	if req.Project.Description != "" {
		cur.Description = req.Project.Description
	}
	if req.Project.Status != 0 {
		cur.Status = req.Project.Status
	}
	cur.UpdateTime = curTime
	return cur, nil
}

func SaveGroup(ctx context.Context, req *servbp.SaveGroupRequest) error {
	if req.Group == nil {
		return errors.New("param error")
	}
	pro := &models.StarlingProject{}
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ?", req.Group.ProjectKey).Last(pro).Error; err != nil {
		return err
	}
	pre := &models.StarlingGroup{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.Group.ProjectKey, req.Group.GroupKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurGroup(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveGroup(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurGroup(pre *models.StarlingGroup, req *servbp.SaveGroupRequest) (*models.StarlingGroup, error) {
	if req.Group.ProjectKey == "" || req.Group.GroupKey == "" {
		return nil, errors.New("project_key and group_key need")
	}
	cur := &models.StarlingGroup{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.Group.GroupName == "" {
			return nil, errors.New("add project need project_name")
		}
		if req.Group.Description == "" {
			return nil, errors.New("add project need description")
		}
		if req.Group.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.Group.ProjectKey
		cur.GroupKey = req.Group.GroupKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.Group.GroupName != "" {
		cur.GroupName = req.Group.GroupName
	}
	if req.Group.Description != "" {
		cur.Description = req.Group.Description
	}
	if req.Group.Status != 0 {
		cur.Status = req.Group.Status
	}
	cur.UpdateTime = curTime
	return cur, nil
}

func SaveOriginLg(ctx context.Context, req *servbp.SaveOriginLgRequest) error {
	if req.OriginLang == nil {
		return errors.New("param error")
	}
	gro := &models.StarlingGroup{}
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ?", req.OriginLang.ProjectKey, req.OriginLang.GroupKey).Last(gro).Error; err != nil {
		return err
	}

	pre := &models.StarlingOrigin{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).
		Where("project_key = ? and group_key = ? and conf_key = ?", req.OriginLang.ProjectKey, req.OriginLang.GroupKey, req.OriginLang.LangKey).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurOriginLg(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveStarlingOriginLg(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurOriginLg(pre *models.StarlingOrigin, req *servbp.SaveOriginLgRequest) (*models.StarlingOrigin, error) {
	if req.OriginLang.ProjectKey == "" || req.OriginLang.GroupKey == "" || req.OriginLang.Lang == "" {
		return nil, errors.New("project_key and group_key and conf_key need")
	}
	cur := &models.StarlingOrigin{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.OriginLang.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.OriginLang.ProjectKey
		cur.GroupKey = req.OriginLang.GroupKey
		cur.LangKey = req.OriginLang.LangKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.OriginLang.Status != 0 {
		cur.Status = req.OriginLang.Status
	}

	cur.UpdateTime = curTime
	return cur, nil
}

func SaveTransLg(ctx context.Context, req *servbp.SaveTransLgRequest) error {
	if req.TransLang == nil {
		return errors.New("param error")
	}
	ori := &models.StarlingOrigin{}
	if err := config.FccWriteDB.WithContext(ctx).Where("project_key = ? and group_key = ? and lang_key", req.TransLang.ProjectKey, req.TransLang.GroupKey, req.TransLang.LangKey).Last(ori).Error; err != nil {
		return err
	}

	pre := &models.StarlingTranslation{}
	objectType := "update"
	if err := config.FccWriteDB.WithContext(ctx).
		Where("project_key = ? and group_key = ? and lang_key = ? and lang = ?", req.TransLang.ProjectKey, req.TransLang.GroupKey, req.TransLang.LangKey, req.TransLang.Lang).Last(pre).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		objectType = "add"
	}
	cur, err := GetCurTransLg(pre, req)
	if err != nil {
		return err
	}
	if err := models.SaveStarlingTransLg(cur); err != nil {
		return err
	}
	_ = service.SaveHistory(pre, cur, cur.TableName(), cur.ProjectKey, objectType, req.OpId)
	return nil
}

func GetCurTransLg(pre *models.StarlingTranslation, req *servbp.SaveTransLgRequest) (*models.StarlingTranslation, error) {
	if req.TransLang.ProjectKey == "" || req.TransLang.GroupKey == "" || req.TransLang.LangKey == "" || req.TransLang.Lang == "" {
		return nil, errors.New("project_key and group_key and conf_key need")
	}
	cur := &models.StarlingTranslation{}
	curTime := time.Now().Unix()
	if pre == nil || pre.Id == 0 {
		if req.TransLang.Status == 0 {
			return nil, errors.New("add project need status")
		}
		cur.ProjectKey = req.TransLang.ProjectKey
		cur.GroupKey = req.TransLang.GroupKey
		cur.LangKey = req.TransLang.LangKey
		cur.CreateTime = curTime
	} else {
		cur = pre
	}
	if req.TransLang.Status != 0 {
		cur.Status = req.TransLang.Status
	}

	cur.UpdateTime = curTime
	return cur, nil
}
