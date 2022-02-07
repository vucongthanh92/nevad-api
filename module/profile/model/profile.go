package profilemode

import (
	"errors"
	"nevad/common"
)

const EntityName = "Profile"

// define avairable to return error
var (
	ErrProfileExisted = common.NewCustomError(
		errors.New("Profile has already existed"),
		"Profile has already existed",
		"ErrProfileExisted",
	)
	ErrProxyNotExist = common.NewCustomError(
		errors.New("proxy is not existed"),
		"Proxy is not existed",
		"ErrProxyNotExist",
	)
)

type Profile struct {
	common.SQLModel `json:",inline"`
	ProfileName     string              `json:"profile_name" gorm:"column:profile_name"`
	UserAgent       string              `json:"user_agent" gorm:"column:user_agent"`
	ProxyId         uint32              `json:"-" gorm:"column:proxy_id"`
	Proxy           *common.SimpleProxy `json:"proxy"`
	Status          bool                `json:"status" gorm:"column:status;default:1"`
}

func (Profile) TableName() string {
	return "profiles"
}

func (p *Profile) Mask() {

	p.GenUID(common.DbTypeProfile)

	if proxy := p.Proxy; proxy != nil {
		p.Proxy.Mask()
	}
}

// structure for create profile
type CreateProfile struct {
	common.SQLModel `json:",inline"`
	ProfileName     *string `json:"profile_name" gorm:"column:profile_name"`
	UserAgent       *string `json:"user_agent" gorm:"column:user_agent"`
	ProxyId         *uint32 `json:"-" gorm:"column:proxy_id"`
	FakeProxyId     *string `json:"proxy_id" gorm:"-"`
}

func (CreateProfile) TableName() string {
	return "profiles"
}

// GetProxyIdFromFakeProxy method get proxy Id from fake proxy and convert it
func (p *CreateProfile) GetProxyIdFromFakeProxy() {

	if *p.FakeProxyId != "" {
		proxyObj, err := common.FromBase58(*p.FakeProxyId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		proxyId := proxyObj.GetLocalID()
		p.ProxyId = &proxyId
	}
}

func (p *CreateProfile) Mask() {

	p.GenUID(common.DbTypeProfile)

}

// structure for update profile
type UpdateProfile struct {
	common.SQLModel `json:",inline"`
	ProfileName     *string `json:"profile_name" gorm:"column:profile_name"`
	UserAgent       *string `json:"user_agent" gorm:"column:user_agent"`
	ProxyId         *int    `json:"-" gorm:"column:proxy_id"`
	FakeProxyId     *string `json:"proxy_id" gorm:"-"`
}

func (UpdateProfile) TableName() string {
	return "profiles"
}

// GetProxyIdFromFakeProxy method get proxy Id from fake proxy and convert it
func (p *UpdateProfile) GetProxyIdFromFakeProxy() {

	if *p.FakeProxyId != "" {
		proxyObj, err := common.FromBase58(*p.FakeProxyId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		proxyId := int(proxyObj.GetLocalID())
		p.ProxyId = &proxyId
	}
}

func (p *UpdateProfile) Mask() {

	p.GenUID(common.DbTypeProfile)

}
