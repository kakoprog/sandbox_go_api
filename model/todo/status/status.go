package status

import (
	"github.com/kakoprog/sandbox_go_api/repository"
)

func StatusByID(id int64) (*repository.Status, error) {
	sqlite := repository.ConnectSqlite()
	status, err := repository.StatusByID(sqlite, id)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func Statuses() ([]*repository.Status, error) {
	sqlite := repository.ConnectSqlite()
	statuses, err := repository.Statuses(sqlite)
	if err != nil {
		return nil, err
	}

	return statuses, nil
}
