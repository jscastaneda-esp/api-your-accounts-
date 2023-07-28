package user_token

import (
	"context"
	"your-accounts-api/shared/domain/persistent"
	persistent_infra "your-accounts-api/shared/infrastructure/db/persistent"
	"your-accounts-api/user/domain"
	"your-accounts-api/user/infrastructure/entity"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func (r *gormRepository) WithTransaction(tx persistent.Transaction) domain.UserTokenRepository {
	return persistent_infra.DefaultWithTransaction[domain.UserTokenRepository](tx, NewRepository, r)
}

func (r *gormRepository) Save(ctx context.Context, userToken domain.UserToken) (uint, error) {
	model := &entity.UserToken{
		Token:     userToken.Token,
		UserId:    userToken.UserId,
		ExpiresAt: userToken.ExpiresAt,
	}

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (r *gormRepository) SearchByExample(ctx context.Context, example domain.UserToken) (*domain.UserToken, error) {
	where := &entity.UserToken{
		Token:  example.Token,
		UserId: example.UserId,
	}
	model := new(entity.UserToken)
	if err := r.db.WithContext(ctx).Where(where).First(model).Error; err != nil {
		return nil, err
	}

	return &domain.UserToken{
		ID:        model.ID,
		Token:     model.Token,
		UserId:    model.UserId,
		ExpiresAt: model.ExpiresAt,
	}, nil
}

func NewRepository(db *gorm.DB) domain.UserTokenRepository {
	return &gormRepository{db}
}
