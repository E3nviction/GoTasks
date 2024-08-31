# GoTasks
GoTasks is a simple command-line interface (CLI) tool written in Go to help you manage your tasks. You can add, delete, edit, move, complete and view tasks easily. The tasks are stored in a CSV file for persistence.

### Features
Add Task: Add a new task with a description.
List Tasks: List tasks in a tabular format.
Delete Task: Delete a task by its ID.
Edit Task: Edit the description of a task by its ID.
Move Task: Change the ID of a task.
Get Task: Get the ID of a task by its name.
Complete Task: Mark a task as completed.
### Upcoming Features
Search Tasks: Search for tasks by keyword.
Prioritize Tasks: Set and manage task priorities.
Task Categories: Categorize tasks for better organization.
Import/Export: Import and export tasks to/from different formats.
Save/Load: Save and load tasks to/from a different file.


## Installation
Clone the repository:
```bash
git clone https://github.com/E3nviction/GoTasks.git
```
Change directory:
```bash
cd GoTasks
```
Build the project:
```bash
go build -o tasks
```
Run the executable:
```bash
./tasks
```
## Usage
The application provides several commands to manage your tasks. Here are the commands you can use:

Add Task
```bash
./tasks add "Your task description"
```
List Tasks
```bash
./tasks list
```
List All Tasks, including Completed
```bash
./tasks list --all # or ./tasks list -a
```
Delete Task
```bash
./tasks delete [id]
```
Edit Task
```bash
./tasks edit [id] "New description"
```
Move Task
```bash
./tasks move [id] [new-id]
```
Get Task ID by name
```bash
./tasks get "task name"
```
Complete Task
```bash
./tasks complete [id]
```