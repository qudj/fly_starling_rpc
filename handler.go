package main

import (
	"context"
	"github.com/qudj/fly_starling_rpc/handler"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
)

type StarlingService struct{}

func NewStarlingServiceServer() servbp.StarlingServiceServer {
	return StarlingService{}
}

func (f StarlingService) FetchProjects(ctx context.Context, req *servbp.FetchProjectsRequest) (*servbp.FetchProjectsResponse, error) {
	ret := &servbp.FetchProjectsResponse{
		BaseRet: &servbp.BaseRet{},
	}
	data, err := handler.FetchProjects(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f StarlingService) FetchGroups(ctx context.Context, req *servbp.FetchGroupsRequest) (*servbp.FetchGroupsResponse, error) {
	ret := &servbp.FetchGroupsResponse{
		BaseRet: &servbp.BaseRet{},
		Data:    &servbp.FetchGroupsRet{},
	}
	data, err := handler.FetchGroups(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f StarlingService) FetchOriginLgs(ctx context.Context, req *servbp.FetchOriginLgsRequest) (*servbp.FetchOriginLgsResponse, error) {
	ret := &servbp.FetchOriginLgsResponse{
		BaseRet: &servbp.BaseRet{},
		Data:    &servbp.FetchOriginLgsRet{},
	}
	data, err := handler.FetchOriginLgs(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f StarlingService) FetchTransLgs(ctx context.Context, req *servbp.FetchTransLgsRequest) (*servbp.FetchTransLgsResponse, error) {
	ret := &servbp.FetchTransLgsResponse{
		BaseRet: &servbp.BaseRet{},
		Data:    &servbp.FetchTransLgsRet{},
	}
	data, err := handler.FetchTransLgs(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = data
	return ret, nil
}

func (f StarlingService) SaveProject(ctx context.Context, req *servbp.SaveProjectRequest) (*servbp.SaveProjectResponse, error) {
	ret := &servbp.SaveProjectResponse{
		BaseRet: &servbp.BaseRet{},
	}
	err := handler.SaveProject(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f StarlingService) SaveGroup(ctx context.Context, req *servbp.SaveGroupRequest) (*servbp.SaveGroupResponse, error) {
	ret := &servbp.SaveGroupResponse{
		BaseRet: &servbp.BaseRet{},
	}
	err := handler.SaveGroup(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f StarlingService) SaveOriginLg(ctx context.Context, req *servbp.SaveOriginLgRequest) (*servbp.SaveOriginLgResponse, error) {
	ret := &servbp.SaveOriginLgResponse{
		BaseRet: &servbp.BaseRet{},
	}
	err := handler.SaveOriginLg(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f StarlingService) SaveTransLg(ctx context.Context, req *servbp.SaveTransLgRequest) (*servbp.SaveTransLgResponse, error) {
	ret := &servbp.SaveTransLgResponse{
		BaseRet: &servbp.BaseRet{},
	}
	err := handler.SaveTransLg(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	return ret, nil
}

func (f StarlingService) FetchTransLg(ctx context.Context, req *servbp.FetchTransLgRequest) (*servbp.FetchTransLgResponse, error) {
	ret := &servbp.FetchTransLgResponse{
		BaseRet: &servbp.BaseRet{},
	}
	value, err := handler.FetchTransLg(ctx, req)
	if err != nil {
		ret.BaseRet.Code = 400
		ret.BaseRet.Msg = err.Error()
		return ret, nil
	}
	ret.Data = value
	return ret, nil
}
