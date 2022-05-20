package postgres

import (
	"database/sql"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/repo"
	"github.com/jmoiron/sqlx"
	"time"
)

type setRepository struct {
	db *sqlx.DB
}

func NewSetRepo(db *sqlx.DB) repo.SetRepoInterface {
	return &setRepository{
		db: db,
	}
}

func (s setRepository) CreateSet(set structs.CreateSet) (structs.SetStruct, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO sets(name, description, service_id)
	 VALUES ($1, $2, $3) returning id`, set.Name, set.Description, set.ServiceId).Scan(&id)
	if err != nil {
		return structs.SetStruct{}, err
	}

	setNew, err := s.GetSet(id)
	if err != nil {
		return structs.SetStruct{}, err
	}

	return setNew, nil
}

func (s setRepository) GetSet(id int) (structs.SetStruct, error) {
	var set structs.SetStruct
	err := s.db.QueryRow(`select  id, name, description, service_id, created_at, updated_at from sets
	where deleted_at is null and id=$1`, id).
		Scan(&set.Id,
			&set.Name,
			&set.Description,
			&set.ServiceId,
			&set.CreatedAt,
			&set.UpdatedAt)
	if err != nil {
		return structs.SetStruct{}, err
	}

	return set, nil
}

func (s setRepository) GetListSets(serviceId int) ([]structs.SetStruct, error) {
	rows, err := s.db.Queryx(`
		SELECT id, name, description, service_id, created_at, updated_at FROM sets WHERE deleted_at 
		    IS NULL AND service_id=$1 order by id
		`, serviceId)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		sets []structs.SetStruct
	)

	for rows.Next() {
		var set structs.SetStruct
		err = rows.Scan(
			&set.Id,
			&set.Name,
			&set.Description,
			&set.ServiceId,
			&set.CreatedAt,
			&set.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}

	return sets, nil
}

func (s setRepository) UpdateSet(set structs.SetStruct) (structs.SetStruct, error) {
	result, err := s.db.Exec(`UPDATE sets SET name=$1, description=$2, service_id=$3, updated_at=$4 WHERE id=$5`,
		&set.Name,
		&set.Description,
		&set.ServiceId,
		time.Now().UTC(),
		&set.Id)
	if err != nil {
		return structs.SetStruct{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return structs.SetStruct{}, sql.ErrNoRows
	}

	set, err = s.GetSet(set.Id)
	if err != nil {
		return structs.SetStruct{}, err
	}

	return set, err
}

func (s setRepository) DeleteSet(id int) error {
	result, err := s.db.Exec(`UPDATE sets SET deleted_at = $1 WHERE id = $2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
