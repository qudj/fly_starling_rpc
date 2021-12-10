package handler

import (
	"context"
	"github.com/qudj/fly_starling_rpc/models"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
)

func FetchProjects(ctx context.Context, req *servbp.FetchProjectsRequest) (*servbp.FetchProjectsRet, error) {
	filter := make(map[string]interface{})
	if proKey, ok := req.Filter["project_key"]; ok {
		filter["project_key"] = proKey
	}
	if proName, ok := req.Filter["project_name"]; ok {
		filter["project_name"] = proName
	}
	list, count, err := models.GetProjects(ctx, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &servbp.FetchProjectsRet{
		Total: count,
		List:  FormatProjectRetList(list),
	}
	return ret, nil
}

func FormatProjectRetList(res []*models.StarlingProject) []*servbp.Project {
	ret := make([]*servbp.Project, 0, len(res))
	for _, v := range res {
		one := &servbp.Project{
			ProjectKey:  v.ProjectKey,
			ProjectName: v.ProjectName,
			Description: v.Description,
			Status:      v.Status,
			OpId:        v.OpId,
			OpName:      v.OpName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		ret = append(ret, one)
	}
	return ret
}

func FetchGroups(ctx context.Context, req *servbp.FetchGroupsRequest) (*servbp.FetchGroupsRet, error) {
	filter := make(map[string]interface{})
	if groKey, ok := req.Filter["group_key"]; ok {
		filter["group_key"] = groKey
	}
	if groName, ok := req.Filter["group_name"]; ok {
		filter["group_name"] = groName
	}
	list, count, err := models.GetGroups(ctx, req.ProjectKey, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &servbp.FetchGroupsRet{
		Total: count,
		List:  FormatGroupRetList(list),
	}
	return ret, nil
}

func FormatGroupRetList(res []*models.StarlingGroup) []*servbp.Group {
	ret := make([]*servbp.Group, 0, len(res))
	for _, v := range res {
		one := &servbp.Group{
			ProjectKey:  v.ProjectKey,
			GroupKey:    v.GroupKey,
			GroupName:   v.GroupName,
			Description: v.Description,
			Status:      v.Status,
			OpId:        v.OpId,
			OpName:      v.OpName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		ret = append(ret, one)
	}
	return ret
}

func FetchOriginLgs(ctx context.Context, req *servbp.FetchOriginLgsRequest) (*servbp.FetchOriginLgsRet, error) {
	filter := make(map[string]interface{})
	if confKey, ok := req.Filter["conf_key"]; ok {
		filter["conf_key"] = confKey
	}
	list, count, err := models.GetStarlingOriginLgs(ctx, req.ProjectKey, req.GroupKey, filter, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &servbp.FetchOriginLgsRet{
		Total: count,
		List:  FormatOriginLgsRet(list),
	}
	return ret, nil
}

func FormatOriginLgsRet(res []*models.StarlingOrigin) []*servbp.OriginLg {
	ret := make([]*servbp.OriginLg, 0, len(res))
	for _, v := range res {
		one := &servbp.OriginLg{
			ProjectKey: v.ProjectKey,
			GroupKey:   v.GroupKey,
			LangKey:    v.LangKey,
			Lang:       v.Lang,
			OriginText: v.OriginText,
			Status:     v.Status,
			OpId:        v.OpId,
			OpName:      v.OpName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		ret = append(ret, one)
	}
	return ret
}

func FetchTransLgs(ctx context.Context, req *servbp.FetchTransLgsRequest) (*servbp.FetchTransLgsRet, error) {
	filter := make(map[string]interface{})
	if confKey, ok := req.Filter["conf_key"]; ok {
		filter["conf_key"] = confKey
	}
	list, count, err := models.GetStarlingTransLgs(ctx, req.ProjectKey, req.GroupKey, req.LangKey, int(req.Offset), int(req.Limit), "id")
	if err != nil {
		return nil, err
	}
	ret := &servbp.FetchTransLgsRet{
		Total: count,
		List:  FormatTransLgsRet(list),
	}
	return ret, nil
}

func FormatTransLgsRet(res []*models.StarlingTranslation) []*servbp.TransLg {
	ret := make([]*servbp.TransLg, 0, len(res))
	for _, v := range res {
		one := &servbp.TransLg{
			ProjectKey:    v.ProjectKey,
			GroupKey:      v.GroupKey,
			LangKey:       v.LangKey,
			Lang:          v.Lang,
			TranslateText: v.TranslateText,
			Status:        v.Status,
			OpId:        v.OpId,
			OpName:      v.OpName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		}
		ret = append(ret, one)
	}
	return ret
}
