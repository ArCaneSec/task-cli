package options

import (
	"flag"
	"log"
	"os"

	"github.com/ArCaneSec/task-cli/internal/task"
)

func ParseFlags() error {
	flag.StringVar(&add, "add", "", "add a task to todo list")
	flag.IntVar(&delete, "del", 0, "delete a task")
	flag.IntVar(&markTodo, "todo", 0, "mark a task as todo")
	flag.IntVar(&markInProgress, "ip", 0, "mark a task as in progress")
	flag.IntVar(&markDone, "done", 0, "mark a task as done")

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
			return err
		}

		log.Printf("[*] Task added successfully with id: %d\n", task.Id)

	case delete != 0:
		err := task.DeleteTask(delete)
		if err != nil {
			return err
		}

		log.Printf("[#] Task number %d got deleted successfully\n", delete)

	case markTodo != 0:
		reqTask, allTasks, err := task.GetTasks(markTodo)
		if err != nil {
			return err
		}
		
		reqTask.Status = task.TODO
		_, err = task.UpdateTasks(reqTask, allTasks)
		if err != nil {
			return err
		}

		log.Printf("[+] Successfully updated task number %d's status to %s\n", reqTask.Id, task.TODO)
		
	case markInProgress != 0:
		reqTask, allTasks, err := task.GetTasks(markInProgress)
		if err != nil {
			return err
		}
		
		reqTask.Status = task.IN_PROGRESS
		_, err = task.UpdateTasks(reqTask, allTasks)
		if err != nil {
			return err
		}

		log.Printf("[+] Successfully updated task number %d's status to %s\n", reqTask.Id, task.IN_PROGRESS)

	case markDone != 0:
		reqTask, allTasks, err := task.GetTasks(markDone)
		if err != nil {
			return err
		}
		
		reqTask.Status = task.DONE
		_, err = task.UpdateTasks(reqTask, allTasks)
		if err != nil {
			return err
		}

		log.Printf("[+] Successfully updated task number %d's status to %s\n", reqTask.Id, task.DONE)

	case len(os.Args) > 1:
		switch os.Args[1] {
		case "update":
			update.Parse(os.Args[2:])

			if updateValue == "" || updateId == 0 {
				log.Fatalln("[!] Please provide id and value for updating task.")
			}
			
			reqTask, allTasks, err := task.GetTasks(updateId)
			if err != nil {
				return err
			}
			reqTask.Description = updateValue

			newTask, err := task.UpdateTasks(reqTask, allTasks)
			if err != nil {
				return err
			}

			log.Printf("[+] Successfully updated task number %d's description to %s\n", newTask.Id, newTask.Description)

		case "list":
			list.Parse(os.Args[2:])
			switch {
			case listTodo:
				if err := task.ListSpecificTasks(task.TODO); err != nil {
					return err
				}

			case listInProgress:
				if err := task.ListSpecificTasks(task.IN_PROGRESS); err != nil {
					return err
				}

			case listDone:
				if err := task.ListSpecificTasks(task.DONE); err != nil {
					return err
				}

			case listAll:
				if err := task.ListAll(); err != nil {
					return err
				}
			}

		}
	}
	return nil
}
