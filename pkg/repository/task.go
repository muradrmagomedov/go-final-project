package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/muradrmagomedov/final-project/pkg/todo"
	"github.com/sirupsen/logrus"
)

// AddTask() adds task to database
func (r *Repository) AddTask(task todo.Task) (string, error) {
	const op = "repository.AddTask"

	var id string
	query := fmt.Sprintf("INSERT INTO %s (date,title,comment,repeat) VALUES ($1,$2,$3,$4) RETURNING id;", schedulerTable)
	row := r.db.QueryRow(query, task.Date, task.Title, task.Comment, task.Repeat)
	err := row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("Не получилось добавить задачу в таблицу. %s:%v", op, err)
	}
	err = row.Err()
	if err != nil {
		return "", fmt.Errorf("Неизвестная ошибка при добавление задачи. %s:%v", op, err)
	}
	return id, nil
}

// GetAllTasks() returns all tasks from schedule table
func (r *Repository) GetAllTasks(searchStr string) (todo.Tasks, error) {
	const op = "repository.GetAllTasks"

	var tasks todo.Tasks
	var condition string
	const orderBy = "date"
	limit := 20
	tasks.Tasks = make([]todo.Task, 0, 1)
	var query string

	if searchStr != "" {
		if date, err := time.Parse("02.01.2006", searchStr); err != nil {
			condition = "WHERE title LIKE $1 OR comment LIKE $1"
			searchStr = `%` + searchStr + `%`
		} else {
			condition = "WHERE date=$1"
			searchStr = date.Format("20060102")
		}
	}

	query = fmt.Sprintf("SELECT * FROM %s %s ORDER BY %s LIMIT %d;", schedulerTable, condition, orderBy, limit)
	rows, err := r.db.Query(query, searchStr)
	if err != nil {
		return tasks, fmt.Errorf("Ошибка получения заданий из базы данных. %s:%v", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var task todo.Task
		if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			logrus.Errorf("%s: Ошибка чтения данных в структуру", op)
		}
		tasks.Tasks = append(tasks.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, fmt.Errorf("Неизвестная ошибка. %s:%v", op, err)
	}
	return tasks, nil
}

// GetTaskById() returns specified task by id if will find it in schedule table
func (r *Repository) GetTaskById(id string) (*todo.Task, error) {
	const op = "repository.GetTaskById"

	var task todo.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", schedulerTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Ошибка получения задачи. %s:%v", op, err)
	case err != nil:
		logrus.Errorf("Неизвестная ошибка. %s:%v", op, err)
		return nil, fmt.Errorf("%s:%v", op, err)
	default:
		return &task, nil
	}
}

// UpdateTask() updates task
func (r *Repository) UpdateTask(task todo.Task) error {
	const op = "repository.UpdateTask"

	query := fmt.Sprintf(`UPDATE %s SET 
	date = $1,
	title = $2,
	comment = $3,
	repeat = $4
	WHERE id = $5;`, schedulerTable)
	result, err := r.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.Id)
	if err != nil {
		return fmt.Errorf("Не получилось внести изменения в задачу. %s:%v", op, err)
	}
	num, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Не получилось внести изменения в задачу. %s:%v", op, err)
	}
	if num == 0 {
		return fmt.Errorf("Не получилось внести изменения в задачу. %s:%v", op, err)
	}
	return nil
}

// Delete() deletes specified task from schedule table
func (r *Repository) DeleteTask(id string) error {
	const op = "repository.DeleteTask"

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", schedulerTable)

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Возникла ошибка при удалении задания. %s:%v", op, err)
	}
	row, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Неизвестная ошибка. %s:%v", op, err)
	}
	if row == 0 {
		return fmt.Errorf("Запись для удаления не найдена. %s:%v", op, err)
	}
	return nil
}
