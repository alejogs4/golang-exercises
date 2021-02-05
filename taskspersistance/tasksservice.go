package taskspersistance

type Task struct {
	Key   int
	Value string
}

// TasksService will define abstractaly how we will interact with any task storage system
type TasksService interface {
	CreateTask(task string) error
	ListAllTasks() ([]Task, error)
	DeleteTask(taskID int) error
}
