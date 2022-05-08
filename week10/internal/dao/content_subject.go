package dao

import (
	"parent-api-go/internal/model/mop"
	"strconv"
)

// 通过id获取详情
func (d *Dao) GetSubjectById(id int32, fields []string) (*mop.ContentSubject, error) {
	contentSubject := mop.ContentSubject{ID: id}
	subject, err := contentSubject.GetSubjectById(d.engine, fields)
	if err != nil {
		return nil, err
	}
	return formatSubject(subject)
}

// 格式化详情
func formatSubject(subject *mop.ContentSubject) (*mop.ContentSubject, error) {
	if subject == (&mop.ContentSubject{}) || subject == nil {
		return nil, nil
	}
	subject.TypeShow = ""
	if subject.Type != "" {
		subjectType, err := strconv.Atoi(subject.Type)
		if err != nil {
			return nil, err
		}
		if val, ok := mop.SubjectTypeConfig[subjectType]; ok {
			subject.TypeShow = val
		}
	}
	return subject, nil
}
