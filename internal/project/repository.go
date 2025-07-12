package project

import (
	"context"
	"database/sql"
	"time"

	// "time"
	"fmt"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/models"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, c *models.Project) error
	GetByProjectId(ctx context.Context, id string) error
	ListProjects(ctx context.Context) error
	UpdateProject(ctx context.Context, c *models.Project) error
	DeleteProject(ctx context.Context, id string) error
}

type ProjectRepo struct {
	DB *sql.DB
}

func (r *ProjectRepo) CreateProject(ctx context.Context, c *models.Project) error {
	c.Id = uuid.New().String()
	c.Createdat = time.Now()
	c.Updatedat = time.Now()
	query := `
					INSERT INTO projects(id, name, created_at, updated_at)
					VALUES ($1, $2, $3, $4)
	      `
	_, err := r.DB.ExecContext(ctx, query, c.Id, c.Name, c.Createdat, c.Updatedat)
	if err != nil {
		return fmt.Errorf("Failed to execute query : %v", err)
	}
	return nil
}

func (r *ProjectRepo) GetByProjectId (ctx context.Context, id string) (*models.Project,error) {
	query := `
				SELECT id, name , created_at, updated_at 
				FROM project 
				WHERE id=$1
	`
	var c models.Project
  
	row := r.DB.QueryRowContext(ctx,query,id) 

	err := row.Scan(&c.Id, &c.Name, &c.Createdat, &c.Updatedat)
	if err != nil {
			return nil,fmt.Errorf("Failed to fetch project : %w",err)
	}
	return &c,nil
}
func NewProjectRepo(db *sql.DB) ProjectRepository {
	return &ProjectRepo{DB: db}
}
