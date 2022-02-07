package proxymodel

import (
	"errors"
	"nevad/common"
)

const EntityName = "Proxy"

type Proxy struct {
	common.SQLModel `json:",inline"`
	Type            string `json:"type" gorm:"column:type;"`
	IP              string `json:"ip" gorm:"column:ip;"`
	Port            string `json:"port" gorm:"column:port;"`
	UserName        string `json:"username" gorm:"column:username;"`
	Password        string `json:"password" gorm:"column:password;"`
}

func (Proxy) TableName() string {
	return "proxies"
}

func (p *Proxy) Mask() {
	p.GenUID(common.DbTypeProxy)
}

var (
	ErrIpExisted = common.NewCustomError(
		errors.New("IP has already existed"),
		"IP has already existed",
		"ErrIpExisted",
	)
)
