package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}

var tasks = []Task{}

func ClearFile(filename string) bool {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	return true
}

func LoadFromFile(filename string) bool {
	file, err := LoadFile(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Skip header
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) != 4 {
			fmt.Println("Invalid record:", record)
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Invalid ID:", record[0])
			continue
		}

		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			fmt.Println("Invalid CreatedAt:", err)
			continue
		}

		completed, err := strconv.ParseBool(record[3])
		if err != nil {
			fmt.Println("Invalid Completed:", record[3])
			continue
		}

		task := Task{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
			Completed:   completed,
		}

		tasks = append(tasks, task)
	}

	if err := CloseFile(file); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SaveToFile(filename string) bool {
	file, err := LoadFile(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"ID", "Description", "CreatedAt", "Completed"})

	for i, task := range tasks {
		writer.Write([]string{
			fmt.Sprintf("%d", i+1),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			fmt.Sprintf("%t", task.Completed),
		})
	}
	writer.Flush()

	if err := CloseFile(file); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func LoadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func ListTasks(listall bool) bool {
	writer := tabwriter.NewWriter(
		os.Stdout, 0, 2, 4, ' ', 0,
	)
	if listall {
		writer.Write(
			[]byte("ID\tDescription\tDate\tCompleted\n"),
		)
	} else {
		writer.Write(
			[]byte("ID\tDescription\tDate\n"),
		)
	}

	for _, task := range tasks {
		if !listall && task.Completed {
			continue
		}
		timedifference := timediff.TimeDiff(task.CreatedAt)
		if listall {
			writer.Write(
				[]byte(fmt.Sprintf(
					"%d\t%s\t%s\t%t\n",
					task.ID,
					task.Description,
					timedifference,
					task.Completed,
				)),
			)
		} else {
			writer.Write(
				[]byte(fmt.Sprintf(
					"%d\t%s\t%s\n",
					task.ID,
					task.Description,
					timedifference,
				)),
			)
		}
	}
	writer.Flush()

	return true
}

func AddTask(description string) bool {
	id := len(tasks) + 1
	tasks = append(tasks, Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	})

	return true
}

func DeleteTask(id int) bool {
	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid ID")
		return false
	}
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found")
		return false
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return true
}

func EditTask(id int, description string) bool {
	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid ID")
		return false
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
		}
	}

	return true
}

func MoveTask(id int, id2 int) bool {
	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid ID")
		return false
	}
	if id2 < 1 || id2 > len(tasks) {
		fmt.Println("Invalid ID")
		return false
	}
	temp := tasks[id-1]
	// delete task
	tasks = append(tasks[:id-1], tasks[id:]...)
	// insert task
	tasks = append(tasks[:id2-1], append([]Task{temp}, tasks[id2-1:]...)...)
	return true
}

func GetTask(name string) int {
	if name == "" {
		return -1
	}
	for i, task := range tasks {
		if task.Description == name {
			return i
		}
	}
	return -1
}

func CompleteTask(id int) bool {
	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid ID")
		return false
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
		}
	}
	return true
}
