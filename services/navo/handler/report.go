package handler

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
)

// Reports 汇报控制器
type Reports struct {
	BaseHandler
}

// NewReportsHandler 创建汇报控制器
func NewReportsHandler(c client.Client) *Reports {
	return &Reports{}
}

// PostReport 发布汇报
func (r *Reports) PostReport(req *restful.Request, rsp *restful.Response) {

}

// PostTemplateReport 发布模板汇报
func (r *Reports) PostTemplateReport(req *restful.Request, rsp *restful.Response) {

}

// Reports 获取汇报列表
func (r *Reports) Reports(req *restful.Request, rsp *restful.Response) {

}

// PendingReportTemplates 获取待汇报的模板列表
func (r *Reports) PendingReportTemplates(req *restful.Request, rsp *restful.Response) {

}

// DownloadReport 下载一个汇报
func (r *Reports) DownloadReport(req *restful.Request, rsp *restful.Response) {

}

// BatchDownload 批量下载汇报列表
func (r *Reports) BatchDownload(req *restful.Request, rsp *restful.Response) {

}

// Details 获取一个汇报详情
func (r *Reports) Details(req *restful.Request, rsp *restful.Response) {

}

// GetReportTemplate 获取模板汇报详情
func (r *Reports) GetReportTemplate(req *restful.Request, rsp *restful.Response) {

}

// PostTemplate 发布模版
func (r *Reports) PostTemplate(req *restful.Request, rsp *restful.Response) {

}

// func (r *Reports) Download(req *restful.)
