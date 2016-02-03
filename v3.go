package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/jberkhahn/v3_beta/commands"
)

type V3Plugin struct {
}

func main() {
	plugin.Start(new(V3Plugin))
}

func (v3plugin *V3Plugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "v3-push" {
		commands.Push(cliConnection, args)
	} else if args[0] == "v3-apps" {
		if len(args) == 1 {
			commands.Apps(cliConnection, args)
		} else {
			//print apps help
		}
	} else if args[0] == "v3-processes" {
		if len(args) == 1 {
			commands.Processes(cliConnection, args)
		} else {
			//print processes help
		}
	} else if args[0] == "v3-delete" {
		if len(args) == 2 {
			commands.Delete(cliConnection, args)
		} else {
			//print help
		}
	} else if args[0] == "v3-logs" {
		if len(args) == 2 {
			commands.Logs(cliConnection, args)
		} else {
			//			print help
		}
	} else if args[0] == "v3-tasks" {
		if len(args) == 2 {
			commands.Tasks(cliConnection, args)
		} else {
			//			print help
		}
	} else if args[0] == "v3-run-task" {
		if len(args) == 4 {
			commands.RunTask(cliConnection, args)
		} else {
			//			print help
		}
	} else if args[0] == "v3-bind-service" {
		if len(args) >= 3 {
			commands.BindService(cliConnection, args)
		} else {
			//			print help
		}
	}
}

func (v3plugin *V3Plugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "v3_beta",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 4,
			Build: 12,
		},
		Commands: []plugin.Command{
			{
				Name:     "v3-push",
				Alias:    "v3-p",
				HelpText: "pushes current dir as a v3 process",
				UsageDetails: plugin.Usage{
					Usage: "v3-push APPNAME",
					Options: map[string]string{
						"p":  "path to dir or zip to push",
						"di": "path to docker image to push",
					},
				},
			},
			{
				Name:     "v3-apps",
				Alias:    "v3-a",
				HelpText: "displays all v3 apps",
				UsageDetails: plugin.Usage{
					Usage:   "v3-apps",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-processes",
				Alias:    "",
				HelpText: "displays all v3 processes",
				UsageDetails: plugin.Usage{
					Usage:   "v3-processes",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-delete",
				Alias:    "v3-d",
				HelpText: "delete a v3 app",
				UsageDetails: plugin.Usage{
					Usage:   "v3-delete APPNAME",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-logs",
				Alias:    "",
				HelpText: "tail logs for a v3 app",
				UsageDetails: plugin.Usage{
					Usage:   "v3-logs APPNAME",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-tasks",
				Alias:    "v3-t",
				HelpText: "list tasks for a v3 app",
				UsageDetails: plugin.Usage{
					Usage:   "v3-tasks APPNAME",
					Options: map[string]string{},
				},
			},
			{
				Name:     "v3-bind-service",
				Alias:    "v3-bs",
				HelpText: "bind a service instance to a v3 app",
				UsageDetails: plugin.Usage{
					Usage: "v3-bind-service APPNAME SERVICEINSTANCE",
					Options: map[string]string{
						"c": "parameters as json",
					},
				},
			},
			{
				Name:     "v3-run-task",
				Alias:    "v3-rt",
				HelpText: "run a task on a v3 app",
				UsageDetails: plugin.Usage{
					Usage:   "v3-run-task APPNAME TASKNAME COMMAND",
					Options: map[string]string{},
				},
			},
		},
	}
}
