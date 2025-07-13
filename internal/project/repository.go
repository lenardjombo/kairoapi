package project

import (
	"context"
	"database/sql"
	"time"
	"fmt"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/models"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, c *models.Project) error
	GetByProjectId(ctx context.Context, id string) (*models.Project,error)
	ListProjects(ctx context.Context) ([]*models.Project,error)
	UpdateProject(ctx context.Context, c *models.Project) error
	DeleteProject(ctx context.Context, id string) error
}

type ProjectRepo struct {
	DB *sql.DB
}

//Create a Project
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
		return fmt.Errorf("failed to execute query : %v", err)
	}
	return nil
}

//Fetch one project
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
			return nil,fmt.Errorf("failed to fetch project : %w",err)
	}
	return &c,nil
}

//List Projects 
func (r *ProjectRepo) ListProjects (ctx context.Context)([]*models.Project,error){
	query := `
			SELECT id, name, created_at, updated_at
			FROM project
	`

	rows,err := r.DB.QueryContext(ctx,query)
	if err != nil {
		return nil,fmt.Errorf("failed to fetch projects : %w",err)
	}
	defer rows.Close()

	var Categories []*models.Project
	for rows.Next() {
		var category models.Project
		if err := rows.Scan(&category.Id, &category.Name, &category.Createdat, &category.Updatedat); err != nil {
			return nil,fmt.Errorf("failed to scan row : %w",err)
		}
		Categories = append(Categories, &category)
	}
	if err = rows.Err(); err != nil {
		return  nil, fmt.Errorf("row iteration : %w",err)
	}
	return  Categories,nil
}

//Update a Project
func (r *ProjectRepo) UpdateProject(ctx context.Context, c *models.Project)(error){
	c.Updatedat = time.Now()

	query := `
			UPDATE project
			SET name=$1,updated_at=$2
			WHERE id=$3
	`
	rows,err := r.DB.ExecContext(ctx,query,&c.Id, &c.Name, &c.Createdat, &c.Updatedat)
	if err != nil {
		return fmt.Errorf("failed to update category : %w",err)
	}

	affected ,err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected : %w",err)
	}

	if affected == 0 {
		return fmt.Errorf("no rows Affected.Failed to fetch category")
	}
	return nil
}

//Delete a Project
func (r *ProjectRepo) DeleteProject(ctx context.Context, id string)(error){
	query := `
				DELETE FROM project WHERE id = $1
	`
	
	rows, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project : %w",err)
	}

	affected,err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows : %w",err)
	}
	if affected == 0 {
		return fmt.Errorf("failed to delete project,Id not found : %w",err)
	}
	return nil
}

func NewProjectRepo(db *sql.DB) ProjectRepository {
	return &ProjectRepo{DB: db}
}
