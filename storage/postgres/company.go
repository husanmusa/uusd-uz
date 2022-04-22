package postgres

import (
	"database/sql"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/pkg/utils"
	"github.com/husanmusa/uusd-uz/storage/repo"
	"github.com/jmoiron/sqlx"
	"time"
)

type companyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) repo.CompanyRepoInterface {
	return &companyRepository{
		db: db,
	}
}

func (c companyRepository) CreateCompany(company structs.CompanyStruct) (structs.CompanyStruct, error) {
	cover := utils.StringToNullTime(company.Cover)
	err := c.db.QueryRow(`INSERT INTO companies(name, cover, slogan)
	 VALUES ($1, $2, $3) returning id`, company.Name, cover, company.Slogan).Scan(&company.Id)
	if err != nil {
		return structs.CompanyStruct{}, err
	}
	company, err = c.GetCompany(company.Id)
	if err != nil {
		return structs.CompanyStruct{}, err
	}

	return company, nil
}

func (c companyRepository) GetCompany(id int) (structs.CompanyStruct, error) {
	var (
		company structs.CompanyStruct
		cover   sql.NullString
	)

	err := c.db.QueryRow(`select  id, name, cover, slogan, created_at, updated_at from companies
	where deleted_at is null and id=$1`, id).
		Scan(&company.Id,
			&company.Name,
			&cover,
			&company.Slogan,
			&company.CreatedAt,
			&company.UpdatedAt)
	if err != nil {
		return structs.CompanyStruct{}, err
	}
	if !cover.Valid {
		company.Cover = cover.String
	}
	return company, nil
}

func (c companyRepository) GetListCompanies() ([]structs.CompanyStruct, error) {
	rows, err := c.db.Queryx(`
		SELECT id, name, cover, slogan, created_at, updated_at FROM companies WHERE deleted_at IS NULL order by id
		`)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		companies []structs.CompanyStruct
	)

	for rows.Next() {
		var (
			company structs.CompanyStruct
			cover   sql.NullString
		)
		err = rows.Scan(
			&company.Id,
			&company.Name,
			&cover,
			&company.Slogan,
			&company.CreatedAt,
			&company.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		if !cover.Valid {
			company.Cover = cover.String
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (c companyRepository) UpdateCompany(company structs.CompanyStruct) (structs.CompanyStruct, error) {
	result, err := c.db.Exec(`UPDATE companies SET name=$1, cover=$2, slogan=$3, updated_at=$4 WHERE id=$5`,
		&company.Name,
		&company.Cover,
		&company.Slogan,
		time.Now().UTC(),
		&company.Id)
	if err != nil {
		return structs.CompanyStruct{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return structs.CompanyStruct{}, sql.ErrNoRows
	}

	company, err = c.GetCompany(company.Id)
	if err != nil {
		return structs.CompanyStruct{}, err
	}

	return company, err
}

func (c companyRepository) DeleteCompany(id int) error {
	result, err := c.db.Exec(`UPDATE companies SET deleted_at = $1 WHERE id = $2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
