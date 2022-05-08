package service

import (
	"parent-api-go/internal/model/mop"
	"parent-api-go/pkg/app"
)

type ContentVideoListRequest struct {
	SubjectID int32 `form:"subject_id" binding:"required,min=0,max=100"`
	Status    uint8 `form:"status,default=1" binding:"oneof=0 1"`
}

func (svc *Service) GetContentVideoList(param *ContentVideoListRequest, pager *app.Pager, order string, lineId uint8) (videoList []*mop.ContentVideo, videoCount int, err error) {
	videoCount, err = svc.daoMopSlave.CountVideoListBySubjectID(param.SubjectID, lineId)
	if err != nil {
		return
	}
	videoList, err = svc.daoMopSlave.GetVideoListBySubjectID(param.SubjectID, pager.Page, pager.PageSize, order, lineId)
	return
}

func (svc *Service) GetVideoDetail(id int32) (video *mop.ContentVideo, err error) {
	video, err = svc.daoMopSlave.GetVideoDetailById(id)
	return
}
