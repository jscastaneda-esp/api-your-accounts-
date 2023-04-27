package entity

import (
	project "api-your-accounts/project/infrastructure/entity"
	"api-your-accounts/shared/infrastructure/db/entity"
	"time"
)

type User struct {
	entity.BaseModel
	entity.BaseUpdateModel
	UUID       string            `gorm:"not null;size:32;unique"`
	Email      string            `gorm:"not null;unique"`
	Projects   []project.Project `gorm:"foreignKey:UserId;references:ID"`
	UserTokens []UserToken       `gorm:"foreignKey:UserId;references:ID"`
}

type UserToken struct {
	entity.BaseModel
	Token          string `gorm:"not null;size:2000;unique"`
	UserId         uint   `gorm:"not null"`
	RefreshedId    uint
	ExpiresAt      time.Time `gorm:"not null"`
	RefreshedAt    time.Time
	RefreshedToken *UserToken `gorm:"foreignKey:RefreshedId"`
}
