package task

import (
	"github.com/kakoprog/sandbox_go_api/repository"
)

func TaskByID(id int64) (*repository.Task, error) {
	sqlite := repository.ConnectSqlite()
	task, err := repository.TaskByID(sqlite, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func Tasks() ([]*repository.Task, error) {
	sqlite := repository.ConnectSqlite()
	tasks, err := repository.Tasks(sqlite)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Save(t *repository.Task) (*repository.Task, error) {
	sqlite := repository.ConnectSqlite()
	err := t.Save(sqlite)
	if err != nil {
		return nil, err
	}

	task, err := repository.TaskByID(sqlite, t.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func Delete(t *repository.Task) error {
	sqlite := repository.ConnectSqlite()
	return t.Delete(sqlite)
}
