package setting

import "time"

const (
	DEV    string = "dev"    // 开发环境
	SIM    string = "sim"    // sim环境
	ONLINE string = "online" // 生产环境
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	Name                  string
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type LinkerdSettingS struct {
	Host    string
	Token   string
	AppName string
}

type SkywalkingSettingS struct {
	AdminAddress string
	ServerHost   string
	ServerName   string
}

type RabbitmqSettingS struct {
	AdminAddress string
	ServerHost   string
	User         string
	PassWord     string
}

type RedisSettingS struct {
	Host         string
	Port         int
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
	TimeOut      int
}

type AccountCardsSettingS struct {
	IosAuditHidden []string `json:"ios_audit_hidden"`
	HiddenTip      []string `json:"hidden_tip"`
	//UserBalanceInfo UserBalanceInfoStruct `json:"user_balance_info"`
	Balance   CardBalanceStruct   `json:"balance"`
	Recommend CardRecommendStruct `json:"recommend"`
	CardList  []CardListStruct    `json:"card_list"`
}

type UserBalanceInfoStruct struct {
	UserBalance        int    `json:"user_balance"`
	CoinBalance        int    `json:"coin_balance"`
	DisplayCoinBalance string `json:"display_coin_balance"`
	RestClass          string `json:"rest_class"`
	RestMainClass      string `json:"rest_main_class"`
	RestNotMainClass   string `json:"rest_not_main_class"`
}

type BalanceItemStruct struct {
	Name     string `json:"name"`
	ItemCode string `json:"item_code"`
	IsH5     bool   `json:"is_h5"`
	Icon     string `json:"icon"`
	Link     string `json:"link"`
}
type CardBalanceStruct struct {
	Title      string              `json:"title"`
	ModuleCode string              `json:"module_code"`
	Arrange    int                 `json:"arrange"`
	Items      []BalanceItemStruct `json:"items"`
}

/*
type CardBalance struct {
	Class struct {
		Name string `json:"name"`
		IsH5 bool   `json:"is_h5"`
		Icon string `json:"icon"`
		Link string `json:"link"`
	} `json:"class"`
	Balance struct {
		Name string `json:"name"`
		IsH5 bool   `json:"is_h5"`
		Icon string `json:"icon"`
		Link string `json:"link"`
	} `json:"balance"`
	Coin struct {
		Name string `json:"name"`
		IsH5 bool   `json:"is_h5"`
		Icon string `json:"icon"`
		Link string `json:"link"`
	} `json:"coin"`
	MainClass struct {
		Name string `json:"name"`
		IsH5 bool   `json:"is_h5"`
		Icon string `json:"icon"`
		Link string `json:"link"`
	} `json:"main_class"`
	NoMainClass struct {
		Name string `json:"name"`
		IsH5 bool   `json:"is_h5"`
		Icon string `json:"icon"`
		Link string `json:"link"`
	} `json:"no_main_class"`
}
*/

type CardRecommendStruct struct {
	Title      string                `json:"title"`
	ModuleCode string                `json:"module_code"`
	Arrange    int                   `json:"arrange"`
	Items      []RecommendItemStruct `json:"items"`
}

type RecommendItemStruct struct {
	Name       string `json:"name"`
	ItemCode   string `json:"item_code"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ButtonText string `json:"button_text"`
	IsH5       bool   `json:"is_h5"`
	Icon       string `json:"icon"`
	Link       string `json:"link"`
	Tip        bool   `json:"tip"`
	TipType    int    `json:"tip_type"`
	TipContent string `json:"tip_content"`
	HelpID     string `json:"help_id,omitempty"`
}

type CardListItem struct {
	Name       string `json:"name"`
	ItemCode   string `json:"item_code"`
	ModuleCode string `json:"module_code"`
	IsH5       bool   `json:"is_h5"`
	Icon       string `json:"icon"`
	Link       string `json:"link"`
	Tip        bool   `json:"tip,omitempty"`
	TipType    int    `json:"tip_type,omitempty"`
	TipContent string `json:"tip_content,omitempty"`
	TipID      string `json:"tip_id,omitempty"`
}

type CardListStruct struct {
	Title      string         `json:"title"`
	ModuleCode string         `json:"module_code"`
	Arrange    int            `json:"arrange"`
	Items      []CardListItem `json:"items"`
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
