package controller

import (
	"application/model"
	"context"
	"errors"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var DATABASE = []model.Todo{
	{
		ID:         "1",
		Title:      "Satu",
		IsDone:     false,
		CreatedAdd: time.Now(),
	},
}

// * = Pointer
func TodoCreate(ctx context.Context, title string) (*model.Todo, error) {
	// create new model
	// id, _ :=
	newTodo := model.Todo{
		ID:         gonanoid.Must(),
		Title:      title,
		IsDone:     false,
		CreatedAdd: time.Now(),
	}
	// insert into database
	DATABASE = append(DATABASE, newTodo)
	// return new model
	return &newTodo, nil
}

func TodoFindAll(ctx context.Context) ([]model.Todo, error) {
	return DATABASE, nil
}

func TodoFindById(ctx context.Context, id string) (*model.Todo, error) {
	for _, todo := range DATABASE {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Todo not Found")
}

func TodoUpdate(ctx context.Context, todo model.Todo) (*model.Todo, error) {
	for idx, item := range DATABASE {
		if item.ID == todo.ID {
			DATABASE[idx] = todo

			return &todo, nil
		}
	}
	return nil, errors.New("Todo not Found")

}

func TodoDelete(ctx context.Context, id string) error {
	var tmp []model.Todo

	for _, item := range DATABASE {
		if item.ID != id {
			tmp = append(tmp, item)
		}
	}
	DATABASE = tmp

	return nil
}
