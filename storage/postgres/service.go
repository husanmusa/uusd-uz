package postgres

import (
	"database/sql"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/repo"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewServiceRepo(db *sqlx.DB) repo.ServiceRepoInterface {
	return &serviceRepository{
		db: db,
	}
}

type serviceRepository struct {
	db *sqlx.DB
}

func (s serviceRepository) CreateService(service structs.CreateService) (structs.ServiceStruct, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO services(name, description, company_id)
	 VALUES ($1, $2, $3) returning id`, service.Name, service.Description, service.CompanyId).Scan(&id)
	if err != nil {
		return structs.ServiceStruct{}, err
	}

	serviceNew, err := s.GetService(id)
	if err != nil {
		return structs.ServiceStruct{}, err
	}

	return serviceNew, nil
}

func (s serviceRepository) GetService(id int) (structs.ServiceStruct, error) {
	var service structs.ServiceStruct
	err := s.db.QueryRow(`select  id, name, description, company_id, created_at, updated_at from services
	where deleted_at is null and id=$1`, id).
		Scan(&service.Id,
			&service.Name,
			&service.Description,
			&service.CompanyId,
			&service.CreatedAt,
			&service.UpdatedAt)
	if err != nil {
		return structs.ServiceStruct{}, err
	}

	return service, nil
}

func (s serviceRepository) GetListServices(companyId int) ([]structs.ServiceStruct, error) {
	rows, err := s.db.Queryx(`
		SELECT id, name, description, company_id, created_at, updated_at FROM services WHERE deleted_at IS NULL 
	   AND company_id=$1 order by id
		`, companyId)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		services []structs.ServiceStruct
	)

	for rows.Next() {
		var service structs.ServiceStruct
		err = rows.Scan(
			&service.Id,
			&service.Name,
			&service.Description,
			&service.CompanyId,
			&service.CreatedAt,
			&service.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (s serviceRepository) UpdateService(service structs.ServiceStruct) (structs.ServiceStruct, error) {
	result, err := s.db.Exec(`UPDATE services SET name=$1, description=$2, company_id=$3, updated_at=$4 WHERE id=$5`,
		&service.Name,
		&service.Description,
		&service.CompanyId,
		time.Now().UTC(),
		&service.Id)
	if err != nil {
		return structs.ServiceStruct{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return structs.ServiceStruct{}, sql.ErrNoRows
	}

	service, err = s.GetService(service.Id)
	if err != nil {
		return structs.ServiceStruct{}, err
	}

	return service, err
}

func (s serviceRepository) DeleteService(id int) error {
	result, err := s.db.Exec(`UPDATE services SET deleted_at = $1 WHERE id = $2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
