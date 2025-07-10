package cohorts

import (
	"context"
	"errors"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type Service interface {
	CreateCohort(ctx context.Context, c *models.Cohort) error
	GetCohortByID(ctx context.Context, id string) (*models.Cohort, error)
	ListCohorts(ctx context.Context) ([]*models.Cohort, error)
	UpdateCohort(ctx context.Context, c *models.Cohort) error
	DeleteCohort(ctx context.Context, id string) error
}

type CohortService struct {
	repo CohortRepository
}

func NewCohortService(r CohortRepository) Service {
	return &CohortService{repo: r}
}

func (s *CohortService) CreateCohort(ctx context.Context, c *models.Cohort) error {
	if c.Name == "" || c.StartDate.IsZero() {
		return errors.New("missing required fields")
	}

	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return s.repo.Create(ctx, c)
}

func (s *CohortService) GetCohortByID(ctx context.Context, id string) (*models.Cohort, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CohortService) ListCohorts(ctx context.Context) ([]*models.Cohort, error) {
	return s.repo.List(ctx)
}

func (s *CohortService) UpdateCohort(ctx context.Context, c *models.Cohort) error {
	if c.ID == "" {
		return errors.New("id is required for update")
	}
	c.UpdatedAt = time.Now()
	return s.repo.Update(ctx, c)
}

func (s *CohortService) DeleteCohort(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required for deletion")
	}
	return s.repo.Delete(ctx, id)
}
