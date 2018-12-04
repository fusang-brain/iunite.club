package handler

import "github.com/emicklei/go-restful"

type Store struct {
	BaseHandler
}

// GetReportTemplates 获取汇报模板
func (s *Store) GetReportTemplates(req *restful.Request, rsp *restful.Response) {
	// TODO get report templates in store
}

// GetApprovedTemplates 获取审批模板
func (s *Store) GetApprovedTemplates(req *restful.Request, rsp *restful.Response) {
	// TODO get approved templates in store
}

// Templates 获取商店的模板列表
func (s *Store) Templates(req *restful.Request, rsp *restful.Response) {
	// TODO get templates in store
}

// PostTemplate 发布模板
func (s *Store) PostTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO post template in store
}

// UpdateTemplate 更新模板
func (s *Store) UpdateTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO update template in store
}

// DeleteTemplate 删除模板
func (s *Store) DeleteTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO delete template in store
}

// BuyTemplate 购买一个模板
func (s *Store) BuyTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO buy one template
}

// EnableOneTemplate 启用一个模板
func (s *Store) EnableOneTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO enable one template
}

// DisableOneTemplate 停用一个模板
func (s *Store) DisableOneTemplate(req *restful.Request, rsp *restful.Response) {
	// TODO disable one template
}