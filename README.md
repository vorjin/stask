# stask

A command-line task manager built with Go. Manage your TODOs directly from the terminal with a simple, fast interface backed by an embedded BoltDB database.

## Features

- **Add tasks** with free-form descriptions
- **Complete tasks** by ID
- **Delete tasks** (soft-delete) by ID
- **List tasks** filtered by status (`todo`, `completed`, `deleted`)
- **Time-based filtering** for completed/deleted tasks (e.g., last 24 hours)
- **Colorized output** for quick visual feedback
- **Zero configuration** -- data is stored locally at `~/.stask/task-manager.db`

## Installation

### From source

Requires Go 1.25+.

```bash
git clone https://github.com/vorjin/CLI-task-manager.git
cd CLI-task-manager
go build -o stask .
```

Move the binary to a directory in your `PATH`:

```bash
mv stask /usr/local/bin/
```

## Usage

### Add a task

```bash
stask add Buy groceries
stask add "Finish the report by Friday"
```

### List tasks

```bash
# List all TODO tasks (default)
stask list

# List completed tasks from the last 24 hours (default)
stask list -s completed

# List deleted tasks from the last 48 hours
stask list -s deleted -t 48
```

**Flags:**
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--status` | `-s` | `todo` | Filter by status: `todo`, `completed`, `deleted` |
| `--time` | `-t` | `24` | Hours to look back (applies to `completed` and `deleted`) |

### Complete a task

```bash
stask do 1
stask do 1 2 3
```

### Delete a task

```bash
stask del 1
stask del 1 2 3
```

## Project Structure

```
stask/
├── main.go          # Entry point -- initializes DB and runs CLI
├── cmd/             # Cobra command definitions (add, do, del, list)
├── db/              # BoltDB persistence layer (implements TaskStore)
├── model/           # Domain types (Task, TaskStatus, TaskStore interface)
├── go.mod
└── go.sum
```

## Tech Stack

- **Go** -- core language
- **[Cobra](https://github.com/spf13/cobra)** -- CLI framework
- **[BoltDB](https://github.com/boltdb/bolt)** -- embedded key/value database
- **[fatih/color](https://github.com/fatih/color)** -- terminal color output

## License

See [LICENSE](LICENSE).
