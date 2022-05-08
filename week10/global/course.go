package global

//主线课
var MainClassCourseTypes = []int{0, 1, 2, 3, 4, 5, 101, 102, 200, 201, 202, 203, 204, 205, 221, 222, 223, 224, 225, 701, 801, 802, 803, 901, 902, 903, 904, 905, 5001, 1223}

//非主线课
var NotMainClassCourseTypes = []int{1003, 2001, 2004, 1001, 1002, 3001, 5002}

// 课时包可用状态
const AVAILABLE int = 2

//课程类型
const CourseTypeInvalid int = 0
const CourseTypeV1PrekGk int = 1
const CourseTypeV1G1G2 int = 2
const CourseTypeV1G3G4 int = 3
const CourseTypeV1G5G6 int = 4
const CourseTypeV1G7More int = 5
const CourseTypeV2L1L3 int = 101
const CourseTypeV2L4More int = 102
const CourseTypeLevelStart int = 200
const CourseTypeV3PrekGk int = 201
const CourseTypeV3G1G2 int = 202
const CourseTypeV3G3G4 int = 203
const CourseTypeV3G5G6 int = 204
const CourseTypeV3G7More int = 205
const CourseTypeV3MinorPrekGk int = 221
const CourseTypeV3MinorG1G2 int = 222
const CourseTypeV3MinorG3G4 int = 223
const CourseTypeV3MinorG5G6 int = 224
const CourseTypeV3MinorG7More int = 225
const CourseTypePypMajor int = 701
const CourseTypeWnsA int = 801
const CourseTypeWnsB int = 802
const CourseTypeWnsC int = 803
const CourseTypeOralRegular int = 901
const CourseTypeWns2017 int = 902
const CourseTypeEbRegular int = 903
const CourseTypeCtRegular int = 904
const CourseTypePypTrial int = 905
const CourseTypeLevelCheck int = 1001
const CourseTypeTrial int = 1002
const CourseTypeCtExercise int = 1003
const CourseTypeRecordedLesson int = 2000
const CourseTypePublic int = 2001
const CourseTypeInduction int = 2002
const CourseTypeItTest int = 2003
const CourseTypeFreePublic int = 2004
const CourseTypeFestival int = 2005
const CourseTypeTeacherTraining int = 2006
const CourseTypeExtend int = 3001

//课程名称字典
var CourseTypeNameDict = map[int]string{
	//过时课程类型
	CourseTypeInvalid:         "无效类型",
	CourseTypeV1PrekGk:        "V1主课启蒙-基础",
	CourseTypeV1G1G2:          "V1主课G1-G2",
	CourseTypeV1G3G4:          "V1主课G3-G4",
	CourseTypeV1G5G6:          "V1主课G5-G6",
	CourseTypeV1G7More:        "V1主课G7以上",
	CourseTypeV2L1L3:          "V2主课低阶",
	CourseTypeV2L4More:        "V2主课高阶",
	CourseTypeWns2017:         "假期课(2017)",
	CourseTypeInduction:       "外教演示课",
	CourseTypeItTest:          "IT Test",
	CourseTypeFestival:        "节日课",
	CourseTypeTeacherTraining: "外教培训课",
	//在售课程类型
	CourseTypeLevelStart:     "鲸品小班外教启蒙课",
	CourseTypeV3PrekGk:       "鲸品小班主修课PREK-GK",
	CourseTypeV3G1G2:         "鲸品小班主修课G1-G2",
	CourseTypeV3G3G4:         "鲸品小班主修课G3-G4",
	CourseTypeV3G5G6:         "鲸品小班主修课G5-G6",
	CourseTypeV3G7More:       "鲸品小班主修课G7以上",
	CourseTypeV3MinorPrekGk:  "鲸品小班辅修课PREK-GK",
	CourseTypeV3MinorG1G2:    "鲸品小班辅修课G1-G2",
	CourseTypeV3MinorG3G4:    "鲸品小班辅修课G3-G4",
	CourseTypeV3MinorG5G6:    "鲸品小班辅修课G5-G6",
	CourseTypeV3MinorG7More:  "鲸品小班辅修课G7以上",
	CourseTypePypMajor:       "PYP主课",
	CourseTypeWnsA:           "鲸品小班假期课-初级",
	CourseTypeWnsB:           "鲸品小班假期课-中级",
	CourseTypeWnsC:           "鲸品小班假期课-高级",
	CourseTypeOralRegular:    "口语体验小班课",
	CourseTypeEbRegular:      "电商专属小班课",
	CourseTypeCtRegular:      "鲸品小班中教课",
	CourseTypePypTrial:       "PYP体验课",
	CourseTypeLevelCheck:     "测试课",
	CourseTypeTrial:          "体验课",
	CourseTypeCtExercise:     "中教练习课",
	CourseTypeRecordedLesson: "录播课",
	CourseTypePublic:         "鲸品公开课",
	CourseTypeFreePublic:     "免费公开课",
	CourseTypeExtend:         "其它",
}

const CourseTextbookTypeNormal int = 0
const CourseTextbookTypeReach int = 1

var CourseTextBookType = map[int]string{
	CourseTextbookTypeNormal: "常规教材",
	CourseTextbookTypeReach:  "Reach",
}
