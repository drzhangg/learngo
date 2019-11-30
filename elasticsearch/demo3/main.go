package demo3

//用户
type User struct {
	Id          int64  `json:"id"`
	NickName    string `json:"nickname"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	CreateAt    string `json:"create_at"`
	LastLoginAt int64  `json:"last_login_at"`
}

//每天的站点访问记录，用来聚合出每月的访问记录
type SiteVisitLog struct {
	Id     string `json:"id"`      //es随机id
	SiteId int64  `json:"site_id"` //站点id
	Day    string `json:"day"`     //每天一条记录，用每天凌晨代表一天
	Pv     int64  `json:"pv"`      //每访问一次就记录
	Uv     int64  `json:"uv"`      //使用cookie去重，每天(0-24h)内同一个cookie算一个
}

