package main

import (
	"flag"
	"log"
	"os"

	"github.com/ArCaneSec/task-cli/internal/task"
)

func main() {
	var (
		add string

		update      *flag.FlagSet
		updateId    int
		updateValue string

		delete int

		inProgress int
		done       int

		list           *flag.FlagSet
		listDone       bool
		listInProgress bool
		listTodo       bool
		listAll        bool
	)
	flag.StringVar(&add, "add", "", "add a task to todo list")
	flag.IntVar(&delete, "del", 0, "delete a task")
	flag.IntVar(&inProgress, "ip", 0, "mark a task as in progress")
	flag.IntVar(&done, "done", 0, "mark a task as done")

	update = flag.NewFlagSet("update", flag.ExitOnError)
	update.IntVar(&updateId, "id", 0, "id of the task that getting updated")
	update.StringVar(&updateValue, "val", "", "new value of the task that getting updated")

	list = flag.NewFlagSet("list", flag.ExitOnError)
	list.BoolVar(&listDone, "done", false, "only list done tasks")
	list.BoolVar(&listInProgress, "ip", false, "only list in progress tasks")
	list.BoolVar(&listTodo, "todo", false, "only list todo tasks")
	list.BoolVar(&listAll, "all", true, "list all tasks")

	flag.Parse()

	switch {
	case add != "":

		task, err := task.AddTask(add)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("[*] Task added successfully with id: %d\n", task.Id)

	case delete != 0:
		err := task.DeleteTask(delete)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("[#] Task number %d got deleted successfully\n", delete)

	case len(os.Args) > 2:
		switch os.Args[1] {
		case "update":
			update.Parse(os.Args[2:])

			if updateValue == "" || updateId == 0 {
				log.Fatalln("[!] Please provide id and value for updating task.")
			}

			newTask, err := task.UpdateTask(updateId, updateValue)
			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("[+] Successfully updated task number %d's description to %s\n", newTask.Id, newTask.Description)
		}
	}
}
