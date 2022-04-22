package postgres

import (
	"database/sql"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/repo"
	"github.com/jmoiron/sqlx"
	"time"
)

type packageRepository struct {
	db *sqlx.DB
}

func NewPackageRepo(db *sqlx.DB) repo.PackageRepoInterface {
	return &packageRepository{
		db: db,
	}
}

func (p packageRepository) CreatePackage(pack structs.PackageStruct) (structs.PackageStruct, error) {
	err := p.db.QueryRow(`INSERT INTO packages(name, description, capacity, cost, code, set_id)
	 VALUES ($1, $2, $3, $4, $5, $6) returning id`,
		pack.Name,
		pack.Description,
		pack.Capacity,
		pack.Cost,
		pack.Code,
		pack.SetId,
	).Scan(&pack.Id)
	if err != nil {
		return structs.PackageStruct{}, err
	}

	pack, err = p.GetPackage(pack.Id)
	if err != nil {
		return structs.PackageStruct{}, err
	}

	return pack, nil
}

func (p packageRepository) GetPackage(id int) (structs.PackageStruct, error) {
	var pack structs.PackageStruct
	err := p.db.QueryRow(`select  id, name, description, capacity, cost, code, set_id, created_at, updated_at from packages
	where deleted_at is null and id=$1`, id).
		Scan(&pack.Id,
			&pack.Name,
			&pack.Description,
			&pack.Capacity,
			&pack.Cost,
			&pack.Code,
			&pack.SetId,
			&pack.CreatedAt,
			&pack.UpdatedAt)
	if err != nil {
		return structs.PackageStruct{}, err
	}

	return pack, nil
}

func (p packageRepository) GetListPackages() ([]structs.PackageStruct, error) {
	rows, err := p.db.Queryx(`
		SELECT id, name, description, capacity, cost, code, set_id, created_at, updated_at from packages WHERE deleted_at IS NULL order by id
		`)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		packs []structs.PackageStruct
	)

	for rows.Next() {
		var pack structs.PackageStruct
		err = rows.Scan(
			&pack.Id,
			&pack.Name,
			&pack.Description,
			&pack.Capacity,
			&pack.Cost,
			&pack.Code,
			&pack.SetId,
			&pack.CreatedAt,
			&pack.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		packs = append(packs, pack)
	}

	return packs, nil
}

func (p packageRepository) UpdatePackage(pack structs.PackageStruct) (structs.PackageStruct, error) {
	result, err := p.db.Exec(`UPDATE packages SET name=$1, description=$2, capacity=$3, cost=$4, code=$5,
                    set_id=$6, updated_at=$7 WHERE id=$8`,
		&pack.Name,
		&pack.Description,
		&pack.Capacity,
		&pack.Cost,
		&pack.Code,
		&pack.SetId,
		time.Now().UTC(),
		&pack.Id)
	if err != nil {
		return structs.PackageStruct{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return structs.PackageStruct{}, sql.ErrNoRows
	}

	pack, err = p.GetPackage(pack.Id)
	if err != nil {
		return structs.PackageStruct{}, err
	}

	return pack, err
}

func (p packageRepository) DeletePackage(id int) error {
	result, err := p.db.Exec(`UPDATE packages SET deleted_at = $1 WHERE id = $2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
