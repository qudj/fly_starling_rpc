package handler

import (
	"context"
	"fmt"
	"github.com/qudj/fly_starling_rpc/models"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
	"golang.org/x/sync/singleflight"
	"sort"
	"strings"
)

var gsf singleflight.Group

func FetchTransLgsByKey(ctx context.Context, req *servbp.FetchTransLgsByKeyRequest) ([]*servbp.TransLg, error) {
	sort.Strings(req.LangKeys)
	key := fmt.Sprintf("TSLG:%s_%s_%s_%s", req.ProjectKey, req.GroupKey, strings.Join(req.LangKeys, "|"), req.Lang)
	gRes, err, _ := gsf.Do(key, func() (interface{}, error) {
		return models.GetStarlingTransLgsByKey(ctx, req.ProjectKey, req.GroupKey, req.Lang, req.LangKeys)
	})
	if err != nil {
		return nil, err
	}
	res := gRes.([]*models.StarlingTranslation)
	return FormatTransLgsRet(res), nil
}