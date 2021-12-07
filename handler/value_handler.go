package handler

import (
	"context"
	"fmt"
	"github.com/qudj/fly_starling_rpc/models"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
	"golang.org/x/sync/singleflight"
)

var gsf singleflight.Group

func FetchTransLg(ctx context.Context, req *servbp.FetchTransLgRequest) (*servbp.TransLg, error) {
	key := fmt.Sprintf("TSLG:%s_%s_%s_%s", req.ProjectKey, req.GroupKey, req.LangKey, req.Lang)
	gRes, err, _ := gsf.Do(key, func() (interface{}, error) {
		return models.GetStarlingTransLg(ctx, req.ProjectKey, req.GroupKey, req.LangKey, req.Lang)
	})
	if err != nil {
		return nil, err
	}
	res := gRes.(*models.StarlingTranslation)
	return FormatConfigRet(res), nil
}

func FormatConfigRet(conf *models.StarlingTranslation) *servbp.TransLg {
	ret := &servbp.TransLg{
		ProjectKey:  conf.ProjectKey,
		GroupKey:    conf.GroupKey,
		Status:      conf.Status,
	}
	return ret
}
