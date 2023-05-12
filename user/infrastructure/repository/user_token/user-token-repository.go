package user_token

import (
	"api-your-accounts/shared/domain/persistent"
	"api-your-accounts/user/domain"
	"api-your-accounts/user/infrastructure/entity"
	"context"

	"gorm.io/gorm"
)

type gormUserTokenRepository struct {
	db *gorm.DB
}

func (r *gormUserTokenRepository) WithTransaction(tx persistent.Transaction) domain.UserTokenRepository {
	if tx, ok := tx.Get().(*gorm.DB); ok {
		return NewRepository(tx)
	}

	return r
}

func (r *gormUserTokenRepository) Create(ctx context.Context, userToken *domain.UserToken) (*domain.UserToken, error) {
	model := &entity.UserToken{
		Token:       userToken.Token,
		UserId:      userToken.UserId,
		ExpiresAt:   userToken.ExpiresAt,
		RefreshedAt: userToken.RefreshedAt,
		RefreshedBy: userToken.RefreshedBy,
	}

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	return &domain.UserToken{
		ID:          model.ID,
		Token:       model.Token,
		UserId:      model.UserId,
		RefreshedBy: model.RefreshedBy,
		CreatedAt:   model.CreatedAt,
		ExpiresAt:   model.ExpiresAt,
		RefreshedAt: model.RefreshedAt,
	}, nil
}

func (r *gormUserTokenRepository) FindByTokenAndUserId(ctx context.Context, token string, userId uint) (*domain.UserToken, error) {
	where := &entity.UserToken{
		Token:  token,
		UserId: userId,
	}
	model := new(entity.UserToken)
	if err := r.db.WithContext(ctx).Where(where).First(model).Error; err != nil {
		return nil, err
	}

	return &domain.UserToken{
		ID:          model.ID,
		Token:       model.Token,
		UserId:      model.UserId,
		RefreshedBy: model.RefreshedBy,
		CreatedAt:   model.CreatedAt,
		ExpiresAt:   model.ExpiresAt,
		RefreshedAt: model.RefreshedAt,
	}, nil
}

func (r *gormUserTokenRepository) Update(ctx context.Context, userToken *domain.UserToken) (*domain.UserToken, error) {
	model := new(entity.UserToken)
	if err := r.db.WithContext(ctx).First(model, userToken.ID).Error; err != nil {
		return nil, err
	}

	model.Token = userToken.Token
	model.UserId = userToken.UserId
	model.RefreshedBy = userToken.RefreshedBy
	model.ExpiresAt = userToken.ExpiresAt
	model.RefreshedAt = userToken.RefreshedAt
	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		return nil, err
	}

	return &domain.UserToken{
		ID:          model.ID,
		Token:       model.Token,
		UserId:      model.UserId,
		RefreshedBy: model.RefreshedBy,
		CreatedAt:   model.CreatedAt,
		ExpiresAt:   model.ExpiresAt,
		RefreshedAt: model.RefreshedAt,
	}, nil
}

func NewRepository(db *gorm.DB) domain.UserTokenRepository {
	return &gormUserTokenRepository{db}
}
