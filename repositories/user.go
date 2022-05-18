package repositories

import (
	"time"
	"xd_working_trial/db"
)

const (
	TableNameUserAccess = "user_access"
)

type UserAccess struct {
	ID        int64     `gorm:"column:id"`
	UserInfo  string    `gorm:"column:user_info"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserAccess) TableName() string {
	return TableNameUserAccess
}

type UserAccessRepository interface {
	Set(*db.DB, *UserAccess) error
	Get(*db.DB) ([]*UserAccess, error)
}

type userAccessRepo struct {
}

func NewUserAccessRepository() UserAccessRepository {
	return &userAccessRepo{}
}

func (u *userAccessRepo) Set(db *db.DB, item *UserAccess) error {
	return db.DB().Table(TableNameUserAccess).Create(item).Error
}

func (u *userAccessRepo) Get(db *db.DB) ([]*UserAccess, error) {
	var items []*UserAccess
	err := db.DB().Table(TableNameUserAccess).Find(&items).Error
	return items, err
}
