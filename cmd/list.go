package cmd

import (
	"fmt"

	"stask/model"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list tasks with a specific status",
		Long:  "list tasks with a specific status",
		RunE: func(cmd *cobra.Command, args []string) error {
			hours, err := cmd.Flags().GetInt("time")
			if err != nil {
				return err
			}

			status, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}

			var tasks []model.Task

			switch status {
			case "todo":
				tasks, err = app.DB.ListTasks(model.Todo, hours)
			case "completed":
				tasks, err = app.DB.ListTasks(model.Completed, hours)
			case "deleted":
				tasks, err = app.DB.ListTasks(model.Deleted, hours)
			}

			if err != nil {
				return err
			}

			if len(tasks) == 0 {
				color.Red("There are no %s tasks! Maaaan, time to do something..\n", status)
				return nil
			}

			switch status {
			case "todo":
				color.Green("These are your TODO tasks: \n")
			default:
				color.Cyan("These are your %s tasks for the last %d hours: \n", status, hours)
			}
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Name)
			}

			return nil
		},
	}

	cmd.Flags().IntP("time", "t", 24, "how many hours ago tasks were transitioned to the specified status")
	f := &statusFlag{"todo", []string{"todo", "completed", "deleted"}}
	cmd.Flags().VarP(f, "status", "s", "filter tasks by status (todo, completed, deleted)")
	return cmd
}

type statusFlag struct {
	value         string
	allowedValues []string
}

func (f *statusFlag) String() string {
	return f.value
}

func (f *statusFlag) Type() string {
	return "string"
}

func (f *statusFlag) Set(value string) error {
	for _, v := range f.allowedValues {
		if v == value {
			f.value = value
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s. Allowed values are: %v", value, f.allowedValues)
}
