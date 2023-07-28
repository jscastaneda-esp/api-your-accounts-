package application

import (
	"context"
	"your-accounts-api/project/domain"
	"your-accounts-api/shared/domain/persistent"
)

//go:generate mockery --name IProjectApp --filename project-app.go
type IProjectApp interface {
	Create(ctx context.Context, userId uint, typeProject domain.ProjectType, tx persistent.Transaction) (uint, error)
	Delete(ctx context.Context, id uint, tx persistent.Transaction) error
	CreateLog(ctx context.Context, description string, projectId uint, detail *string, tx persistent.Transaction) error
	FindLogsByProject(ctx context.Context, projectId uint) ([]*domain.ProjectLog, error)
}

type projectApp struct {
	tm             persistent.TransactionManager
	projectRepo    domain.ProjectRepository
	projectLogRepo domain.ProjectLogRepository
}

func (app *projectApp) Create(ctx context.Context, userId uint, typeProject domain.ProjectType, tx persistent.Transaction) (uint, error) {
	projectRepo := app.projectRepo.WithTransaction(tx)
	newProject := domain.Project{
		UserId: userId,
		Type:   typeProject,
	}
	id, err := projectRepo.Save(ctx, newProject)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *projectApp) Delete(ctx context.Context, id uint, tx persistent.Transaction) error {
	projectRepo := app.projectRepo.WithTransaction(tx)
	return projectRepo.Delete(ctx, id)
}

func (app *projectApp) CreateLog(ctx context.Context, description string, projectId uint, detail *string, tx persistent.Transaction) error {
	projectLogRepo := app.projectLogRepo.WithTransaction(tx)
	newLog := domain.ProjectLog{
		Description: description,
		ProjectId:   projectId,
		Detail:      detail,
	}
	_, err := projectLogRepo.Save(ctx, newLog)
	if err != nil {
		return err
	}

	return nil
}

func (app *projectApp) FindLogsByProject(ctx context.Context, projectId uint) ([]*domain.ProjectLog, error) {
	example := domain.ProjectLog{
		ProjectId: projectId,
	}
	logs, err := app.projectLogRepo.SearchAllByExample(ctx, example)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func NewProjectApp(tm persistent.TransactionManager, projectRepo domain.ProjectRepository, projectLogRepo domain.ProjectLogRepository) IProjectApp {
	return &projectApp{tm, projectRepo, projectLogRepo}
}
