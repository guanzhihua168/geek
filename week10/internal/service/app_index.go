package service

import (
	"parent-api-go/internal/model/mop"
	"parent-api-go/pkg/util"
)

type AppIndexRequest struct {
	Platform int            `form:"platform" binding:"required"`
	Device   RequestDeviceP `json:"device" binding:"required"`
}

func (svc *Service) GetIndexBanners(userId uint32, platform int, appVersion string, lineId uint8) (banners []*mop.BannerModel, err error) {
	var (
		fields        = []string{"id", "position", "channel", "banner_name", "banner_img_url", "link", "line_id", "target_user_type"}
		latestVersion mop.AppVersionModel
		bannerVersion = mop.AppV2Version
	)

	if platform > 0 {
		latestVersion, err = svc.daoMopSlave.GetLatestVersionConfig(platform, appVersion)
		if err != nil {
			return
		}

		// There is no appreciate version with current version in backend, to display all audited banners.
		if latestVersion == (mop.AppVersionModel{}) {
			banners, err = svc.daoMopSlave.GetBannersToApp(fields, bannerVersion, lineId, 10)
		} else {
			if latestVersion.IsAudition == mop.IsAuditing {
				banners, err = svc.daoMopSlave.GetBannersToAppAuditing(fields, bannerVersion, lineId, 10)
			} else if latestVersion.IsAudition == mop.IsAudited {
				banners, err = svc.daoMopSlave.GetBannersToAppAudited(fields, bannerVersion, lineId, 10)
			}
		}
	} else {
		banners, err = svc.daoMopSlave.GetBannersToApp(fields, bannerVersion, lineId, 0)
	}

	if err != nil {
		return
	}

	return svc.BannerFormat(banners, userId, 0)
}

type subject struct {
	subjectId    int32
	offset       int
	limit        int
	identifyLine bool
}

type VideoResult struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	TypeShow    string              `json:"type_show"`
	Idx         int                 `json:"idx"`
	HasMorePage bool                `json:"has_more_page"`
	List        []*mop.ContentVideo `json:"list"`
}

type ArticleResult struct {
	ID          int32                 `json:"id"`
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	TypeShow    string                `json:"type_show"`
	CreatedAt   string                `json:"created_at"`
	Status      uint8                 `json:"status"`
	Idx         int                   `json:"idx"`
	HasMorePage bool                  `json:"has_more_page"`
	List        []*mop.ContentArticle `json:"list"`
}

func (svc *Service) GetIndexContent(lineId uint8) (videoResults []*VideoResult, articleResults []*ArticleResult, err error) {
	/*
		首页视频内容
	*/
	videoSubject := make(map[int]*subject)
	// 认识鲸鱼
	videoSubject[1] = &subject{
		subjectId:    1,
		offset:       0,
		limit:        20,
		identifyLine: true, // 是否区分欧美标
	}

	// 优秀课堂
	videoSubject[2] = &subject{
		subjectId:    2,
		offset:       0,
		limit:        4,
		identifyLine: false,
	}

	/*
		首页图文文章
	*/
	// 鲸鱼动态
	articleSubject := make(map[int]*subject)
	articleSubject[3] = &subject{
		subjectId:    5,
		offset:       0,
		limit:        3,
		identifyLine: false,
	}

	// 增加视频内容
	videoResults, err = svc.getIndexVideoContent(lineId, videoSubject)

	// 增加文章内容
	articleResults, err = svc.getIndexArticleContent(lineId, articleSubject)
	return
}

// 获取视频内容
func (svc *Service) getIndexVideoContent(lineId uint8, videoSubject map[int]*subject) (sliceResults []*VideoResult, err error) {
	var (
		contentSubject *mop.ContentSubject
		contentVideos  []*mop.ContentVideo
		videoCount     int
		hasMorePage    = false
	)

	for key, item := range videoSubject {
		contentSubject, err = svc.daoMopSlave.GetSubjectById(item.subjectId, []string{"id", "name", "type"})
		if contentSubject == nil || contentSubject == (&mop.ContentSubject{}) {
			//没找到也增加一个默认的结构占位
			continue
		}

		// 是否区分欧美标，0表示不区分，获取所有的
		if !item.identifyLine {
			lineId = 0
		}

		contentVideos, err = svc.daoMopSlave.GetVideoListBySubjectID(item.subjectId,
			item.offset, item.limit, "updated_at DESC", lineId)

		videoCount, err = svc.daoMopSlave.CountVideoListBySubjectID(item.subjectId, lineId)
		if videoCount > len(contentVideos) {
			hasMorePage = true
		}

		var results VideoResult
		results.ID = contentSubject.ID
		results.Name = contentSubject.Name
		results.Type = contentSubject.Type
		results.TypeShow = contentSubject.TypeShow
		results.Idx = key
		results.HasMorePage = hasMorePage
		results.List = contentVideos
		sliceResults = append(sliceResults, &results)
	}
	return
}

// 获取文章内容
func (svc *Service) getIndexArticleContent(lineId uint8, articleSubject map[int]*subject) (sliceResults []*ArticleResult, err error) {
	var (
		contentSubject  *mop.ContentSubject
		contentArticles []*mop.ContentArticle
		articleCount    int
		hasMorePage     = false
	)

	for key, item := range articleSubject {
		contentSubject, err = svc.daoMopSlave.GetSubjectById(item.subjectId, []string{})
		if contentSubject == nil || contentSubject == (&mop.ContentSubject{}) {
			continue
		}

		// 家长专区，不区分欧美标，暂时保持原样，无需处理
		contentArticles, err = svc.daoMopSlave.GetArticleListBySubjectID(item.subjectId,
			item.offset, item.limit, lineId,
			[]string{"id", "title", "subtitle", "cover"}, "updated_at DESC")

		articleCount, err = svc.daoMopSlave.CountArticleListBySubjectID(item.subjectId, lineId)
		if articleCount > len(contentArticles) {
			hasMorePage = true
		}

		var results ArticleResult
		results.ID = contentSubject.ID
		results.Name = contentSubject.Name
		results.Type = contentSubject.Type
		results.TypeShow = contentSubject.TypeShow
		results.CreatedAt = util.FormatDate(contentSubject.CreatedAt)
		results.Status = contentSubject.Status
		results.Idx = key
		results.HasMorePage = hasMorePage
		results.List = contentArticles
		sliceResults = append(sliceResults, &results)
	}
	return
}
