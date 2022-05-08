package dao

import (
	"parent-api-go/internal/model/mop"
)

func (d *Dao) CountVideoListBySubjectID(subjectID int32, lineId uint8) (int, error) {
	contentVideo := mop.ContentVideo{SubjectID: subjectID, LineID: lineId}
	return contentVideo.CountVideoListBySubjectID(d.engine)
}

func (d *Dao) GetVideoListBySubjectID(subjectID int32, offset int, limit int, order string, lineId uint8) ([]*mop.ContentVideo, error) {
	contentVideo := mop.ContentVideo{SubjectID: subjectID, LineID: lineId}
	return contentVideo.GetVideoListBySubjectID(d.engine, offset, limit, order)
}

func (d *Dao) GetVideoDetailById(id int32) (*mop.ContentVideo, error) {
	contentVideo := mop.ContentVideo{}
	contentVideo.ID = id
	return contentVideo.GetVideoDetailById(d.engine)
}
