package common

type SimpleProxy struct {
	SQLModel `json:",inline"`
	Type     string `json:"type" gorm:"column:type;"`
	IP       string `json:"ip" gorm:"column:ip;"`
	Port     string `json:"port" gorm:"column:port;"`
}

func (SimpleProxy) TableName() string {
	return "proxies"
}

func (p *SimpleProxy) Mask() {
	p.GenUID(DbTypeProxy)
}
