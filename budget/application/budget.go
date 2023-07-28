package application

import (
	"context"
	"fmt"
	"time"
	"your-accounts-api/budget/domain"
	"your-accounts-api/project/application"
	project_dom "your-accounts-api/project/domain"
	"your-accounts-api/shared/domain/persistent"
)

//go:generate mockery --name IBudgetApp --filename budget-app.go
type IBudgetApp interface {
	Create(ctx context.Context, userId uint, name string) (uint, error)
	Clone(ctx context.Context, userId uint, baseId uint) (uint, error)
	FindById(ctx context.Context, id uint) (*domain.Budget, error)
	FindByUserId(ctx context.Context, userId uint) ([]*domain.Budget, error)
	Delete(ctx context.Context, id uint) error
}

type budgetApp struct {
	tm         persistent.TransactionManager
	budgetRepo domain.BudgetRepository
	projectApp application.IProjectApp
}

func (app *budgetApp) Create(ctx context.Context, userId uint, name string) (uint, error) {
	var id uint
	err := app.tm.Transaction(func(tx persistent.Transaction) error {
		projectId, err := app.projectApp.Create(ctx, userId, project_dom.TypeBudget, tx)
		if err != nil {
			return err
		}

		err = app.projectApp.CreateLog(ctx, "Creación", projectId, nil, tx)
		if err != nil {
			return err
		}

		budgetRepo := app.budgetRepo.WithTransaction(tx)
		now := time.Now()
		year := uint16(now.Year())
		month := uint8(now.Month())
		newBudget := domain.Budget{
			Name:      &name,
			Year:      &year,
			Month:     &month,
			ProjectId: &projectId,
		}
		id, err = budgetRepo.Save(ctx, newBudget)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *budgetApp) Clone(ctx context.Context, userId uint, baseId uint) (uint, error) {
	baseBudget, err := app.FindById(ctx, baseId)
	if err != nil {
		return 0, err
	}

	var id uint
	err = app.tm.Transaction(func(tx persistent.Transaction) error {
		projectId, err := app.projectApp.Create(ctx, userId, project_dom.TypeBudget, tx)
		if err != nil {
			return err
		}

		description := fmt.Sprintf("Creación a partir del presupuesto %s(%d)", *baseBudget.Name, baseId)
		detail := fmt.Sprintf(`{"cloneId": %d, "cloneName": "%s"}`, baseId, *baseBudget.Name)
		err = app.projectApp.CreateLog(ctx, description, projectId, &detail, tx)
		if err != nil {
			return err
		}

		budgetRepo := app.budgetRepo.WithTransaction(tx)
		name := fmt.Sprintf("%s Copia", *baseBudget.Name)
		newBudget := domain.Budget{
			Name:             &name,
			Year:             baseBudget.Year,
			Month:            baseBudget.Month,
			FixedIncome:      baseBudget.FixedIncome,
			AdditionalIncome: baseBudget.AdditionalIncome,
			ProjectId:        &projectId,
		}
		id, err = budgetRepo.Save(ctx, newBudget)
		if err != nil {
			return err
		}

		// TODO Pendiente la creación de AvailableBalances, Bills y BillShared

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *budgetApp) FindById(ctx context.Context, id uint) (*domain.Budget, error) {
	budget, err := app.budgetRepo.Search(ctx, id)
	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (app *budgetApp) FindByUserId(ctx context.Context, userId uint) ([]*domain.Budget, error) {
	budgets, err := app.budgetRepo.SearchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return budgets, nil
}

func (app *budgetApp) Delete(ctx context.Context, id uint) error {
	budget, err := app.FindById(ctx, id)
	if err != nil {
		return err
	}

	return app.tm.Transaction(func(tx persistent.Transaction) error {
		err = app.budgetRepo.Delete(ctx, *budget.ID) // TODO Validar si se puede eliminar toda la información
		if err != nil {
			return err
		}

		return app.projectApp.Delete(ctx, *budget.ProjectId, tx)
	})
}

func NewBudgetApp(tm persistent.TransactionManager, budgetRepo domain.BudgetRepository, projectApp application.IProjectApp) IBudgetApp {
	return &budgetApp{tm, budgetRepo, projectApp}
}
