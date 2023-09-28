package main

import "fmt"

type Pool struct {
	workersNum int
	cli        *ApiClient
}

func NewPool(workersNum int, cli *ApiClient) *Pool {
	return &Pool{
		workersNum: workersNum,
		cli:        cli,
	}
}

// AddTask добавляет в пул задачу, которую нужно взять в работу.
// Если пул забит, то нужно заблокировать поток управления до тех пор, пока не появится
// свободных воркеров.
func (p *Pool) AddTask(t Task) {
	// ...
}

// GetResult возвращает результат выполнения какой-либо задачи. Если пока результатов нет,
// то нужно заблокировать поток управления до тех пор, пока не появится какой-то
// результат от какого-либо воркера.
func (p *Pool) GetResult() User {
	// ...
	return User{}
}

type Task struct {
	UserID int
}

func main() {
	userIds := []int{1, 3, 8, 12, 17, 22, 24, 30, 40, 54, 57, 62, 78, 92, 102, 104}
	pool := NewPool(3, new(ApiClient))

	go func() {
		for _, userID := range userIds {
			pool.AddTask(Task{UserID: userID})
		}
	}()

	for {
		user := pool.GetResult()
		fmt.Println("User:", user)
	}
}
