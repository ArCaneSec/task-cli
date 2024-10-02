## Introduction
**task-cli** is a simple command-line interface task manager application that allows users to manage tasks through basic operations like adding, listing, deleting with multiple statuses. The data is stored in a JSON file. This project is based on [this backend project of roadmap.sh](https://roadmap.sh/projects/task-tracker).
## Features
-   **Add a task**: Create new tasks.
-   **List tasks**: Display all tasks with their current status, supports filtering on status.
-   **Change status of a task**: You can mark tasks as **todo**, **in-progress** and **done**.
-   **Delete a task**: Remove tasks that are no longer needed.
-   **Task persistence**: Tasks are stored in a JSON file, so your progress is saved between sessions.
## Installation
1. clone the repo:
	```bash
	git clone https://github.com/ArCaneSec/task-cli.git
	```
2. navigate to project and build
	```bash
	cd task-cli && go build -o task-cli ./cmd/
	```
## Usage

### Add a task
To add a new task:
```bash
./task-cli -add "my new task"
```
### Delete a task
to delete a task:
```bash
./task-cli -del <id>
```
### Update task
to update description of a task:
```bash
./task-cli update -id <id> -val "updated task"
```
### List tasks
to list all tasks:
```bash
./task-cli list
```
to list all todo tasks:
```bash
./task-cli list -todo 
```
list all in-progress tasks:
```bash
./task-cli  list -ip
```
list all done tasks:
```bash
./task-cli list -done 
```
### Change status
to change status of a task to todo:
```bash
./task-cli -todo <id>
```
change status to in-progress:
```bash
./task-cli  -ip <id>
```
change status to done:
```bash
./task-cli -done <id>
```
## Contributing

Feel free to contribute to this project! If you have ideas for new features, feel free to open a pull request or issue.