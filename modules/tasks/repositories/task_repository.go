package repositories

import (
	"database/sql"
	"fmt"

	"github.com/nuttapon-first/task-clean-arch/modules/entities"
)

type TaskRepo struct {
	Db *sql.DB
}

func NewTaskRepository(db *sql.DB) entities.TaskRepository {
	return &TaskRepo{
		Db: db,
	}
}

func (r *TaskRepo) NewTask(req *entities.NewTaskReq) (*entities.NewTaskRes, error) {
	insert := `
	INSERT INTO tasks(name,status)
	VALUES (?,?);
	`

	res := &entities.NewTaskRes{}
	stmt, err := r.Db.Prepare(insert)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(req.Name, 0)
	if err != nil {
		return res, err
	}

	taskId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	row := r.Db.QueryRow("SELECT id, name, status FROM tasks WHERE id=?", taskId)
	if err := row.Scan(&res.ID, &res.Name, &res.Status); err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("TaskById %d: no such task", taskId)
		}
		return res, fmt.Errorf("TaskById %d: %v", taskId, err)
	}

	return res, nil
}

func (r *TaskRepo) GetTaskList(req *entities.GetTaskListReq) (*[]entities.GetTaskListRes, error) {
	var taskList = []entities.GetTaskListRes{}
	rows, err := r.Db.Query("SELECT * FROM tasks")
	if err != nil {
		return &taskList, err
	}
	defer rows.Close()
	for rows.Next() {
		var task entities.GetTaskListRes
		if err := rows.Scan(&task.ID, &task.Name, &task.Status); err != nil {
			return &taskList, fmt.Errorf("task error: %v", err)
		}
		taskList = append(taskList, task)
	}
	if err := rows.Err(); err != nil {
		return &taskList, fmt.Errorf("task error: %v", err)
	}
	return &taskList, nil
}

func (r *TaskRepo) UpdateTaskById(req *entities.UpdateTaskListReq) (*entities.UpdateTaskListRes, error) {
	var task = entities.UpdateTaskListRes{}
	stmt, err := r.Db.Prepare(`
	UPDATE tasks
	SET name=?, status=?
	WHERE id=?
	`)
	if err != nil {
		return &task, err
	}
	_, err = stmt.Exec(req.Name, req.Status, req.ID)
	if err != nil {
		return &task, fmt.Errorf("update error: %#v", err)
	}

	row := r.Db.QueryRow("SELECT id, name, status FROM tasks WHERE id=?", req.ID)
	if err := row.Scan(&task.ID, &task.Name, &task.Status); err != nil {
		if err == sql.ErrNoRows {
			return &task, fmt.Errorf("TaskById %d: no such task", req.ID)
		}
		return &task, fmt.Errorf("TaskById %d: %v", req.ID, err)
	}
	return &task, nil
}

func (r *TaskRepo) DeleteTaskById(req *entities.DeleteTaskListReq) error {
	fmt.Printf("%#v\n", req)
	statement := `DELETE FROM tasks WHERE id = $1;`
	_, err := r.Db.Exec(statement, req.ID)
	if err != nil {
		return err
	}
	return nil
}
