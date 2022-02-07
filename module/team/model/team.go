package teammodel

import (
	"errors"
	"nevad/common"
	"strings"
)

const EntityName = "Team"

type Team struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	OwnerID         int    `json:"owner_id" gorm:"column:owner_id;"`
	State           string `json:"state" gorm:"column:state;"`
	Numbers         int    `json:"numbers" gorm:"column:numbers;"`
	Operating       string `json:"operating" gorm:"column:operating;"`
}

func (Team) TableName() string {
	return "teams"
}

func (t *Team) Mask(isAdmin bool) {
	t.GenUID(common.DbTypeTeam)
}

type TeamCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	OwnerID         int    `json:"owner_id" gorm:"column:owner_id;"`
	State           string `json:"state" gorm:"column:state;"`
	Numbers         int    `json:"numbers" gorm:"column:numbers;"`
	Operating       string `json:"operating" gorm:"column:operating;"`
}

func (TeamCreate) TableName() string {
	return Team{}.TableName()
}

func (t *TeamCreate) Mask(isAdmin bool) {
	t.GenUID(common.DbTypeTeam)
}

func (t *TeamCreate) Validate() error {
	t.Name = strings.TrimSpace(t.Name)
	if t.Name == "" {
		return ErrNotEmpty
	}

	return nil
}

type TeamUpdate struct {
	Name      *string `json:"name" gorm:"column:name;"`
	OwnerID   *int    `json:"owner_id" gorm:"column:owner_id;"`
	State     *string `json:"state" gorm:"column:state;"`
	Numbers   *int    `json:"numbers" gorm:"column:numbers;"`
	Operating *string `json:"operating" gorm:"column:operating;"`
}

func (TeamUpdate) TableName() string {
	return Team{}.TableName()
}

var (
	ErrNotEmpty = errors.New("Data not empty")
)
