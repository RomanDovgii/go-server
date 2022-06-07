package sqlstore

import (
	"database/sql"

	"github.com/RomanDovgii/go-restapi/internal/app/model"
	"github.com/RomanDovgii/go-restapi/internal/app/store"
)

type WorkRepository struct {
	store *Store
}

func (r *WorkRepository) Create(w *model.Work) error {
	if err := w.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (creator_id, name, description, document_links) VALUES ($1, $2, $3, $5) RETURNING id",
		w.CreatorId,
		w.Name,
		w.Description,
		w.DocumentLinks,
	).Scan(&w.ID)
}

func (r *WorkRepository) Find(id int) (*model.Work, error) {
	w := &model.Work{}
	if err := r.store.db.QueryRow(
		"SELECT id, creator_id, name, description, document_links FROM works WHERE id = $1",
		id,
	).Scan(
		&w.ID,
		&w.CreatorId,
		&w.Description,
		&w.DocumentLinks,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return w, nil
}

func (r *WorkRepository) FindByName(name string) (*model.Work, error) {
	w := &model.Work{}

	if err := r.store.db.QueryRow(
		"SELECT id, creator_id, name, description, document_links FROM works WHERE name = $1",
		name,
	).Scan(
		&w.ID,
		&w.CreatorId,
		&w.Name,
		&w.Description,
		&w.DocumentLinks,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return w, nil
}

func (r *WorkRepository) Delete(workId int, userId int) error {
	if err := r.store.db.QueryRow(
		"DELETE FROM works WHERE id='$1' AND creator_id='$2'",
		workId,
		userId,
	); err != nil {
		return store.ErrRecordNotFound
	}

	return nil
}

func (r *WorkRepository) FindAll(numberOfItems int, pageNumber int) ([]model.Work, error) {
	mw := &model.Work{}
	o := numberOfItems * pageNumber
	sw := []model.Work{}

	rows, err := r.store.db.Query(
		"SELECT id, creator_id, name, description, document_links FROM works ORDER BY id LIMIT $1 OFFSET $2",
		numberOfItems,
		o,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&mw.ID,
			&mw.CreatorId,
			&mw.Description,
			&mw.DocumentLinks,
		)
		if err != nil {
			panic(err)
		}
		sw = append(sw, *mw)
	}

	return sw, nil
}

func (r *WorkRepository) FindAllByName(name string, numberOfItems int, pageNumber int) ([]model.Work, error) {
	mw := &model.Work{}
	o := numberOfItems * pageNumber
	sw := []model.Work{}

	rows, err := r.store.db.Query(
		"SELECT id, creator_id, name, description, document_links FROM works WHERE name LIKE '%$1%' ORDER BY id LIMIT $2 OFFSET $3",
		name,
		numberOfItems,
		o,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&mw.ID,
			&mw.CreatorId,
			&mw.Description,
			&mw.DocumentLinks,
		)
		if err != nil {
			panic(err)
		}
		sw = append(sw, *mw)
	}

	return sw, nil
}
