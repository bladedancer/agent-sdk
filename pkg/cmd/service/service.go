package service

import (
	"bytes"
	"os/exec"
	"strings"
	"text/template"

	"github.com/takama/daemon"

	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/util/log"
)

var (
	// Name -
	Name string
	// Description -
	Description string

	dependencies = []string{"network"}

	// GlobalAgentService -
	GlobalAgentService AgentService
)

var systemDConfig = `[Unit]
Description={{.Description}}
Requires={{.Dependencies}}
After={{.Dependencies}}

[Service]
PIDFile=/var/run/{{.Name}}.pid
ExecStartPre=/bin/rm -f /var/run/{{.Name}}.pid
ExecStart={{.Path}} {{.Args}}
User={{.User}}
Group={{.Group}}
Restart=on-failure

[Install]
WantedBy=multi-user.target
`

//AgentService -
type AgentService struct {
	service     daemon.Daemon
	Name        string
	Description string
	Path        string
	PathArg     string
	User        string
	Group       string
}

func init() {
	service, err := daemon.New(Name, Description, dependencies...)
	if err != nil {
		log.Errorf("error hit creating the service definition: %s", err.Error())
	}

	GlobalAgentService = AgentService{
		service:     service,
		Name:        Name,
		Description: Description,
	}
}

// HandleServiceFlag - handles the action needed based ont eh service flag value
func (a *AgentService) HandleServiceFlag(command string) int {
	var err error
	var status string
	result := 0
	// complete teh appropriate action for the service
	switch strings.ToLower(command) {
	case "install":
		_, err = a.serviceinstall()
	case "remove":
		log.Debug("removing the agent service")
		_, err = a.service.Remove()
	case "start":
		log.Debug("starting the agent service")
		_, err = a.service.Start()
	case "stop":
		log.Debug("stoping the agent service")
		_, err = a.service.Stop()
	case "status":
		log.Debug("getting the agent service status")
		status, err = a.service.Status()
	case "enable":
		_, err = a.serviceEnableReboot()
	default:
		log.Errorf("unknown value of '%s' given", command)
		result = 1
	}

	// error hit
	if err != nil {
		log.Errorf("service %s command failed: %s", strings.ToLower(command), err.Error())
		result = 1
	}

	// log success
	if result == 0 {
		log.Debugf("service %s command succeeded", strings.ToLower(command))
		if status != "" {
			log.Info(status)
		}
	}
	return result
}

func (a *AgentService) serviceinstall() (string, error) {
	log.Debug("installing the agent service")
	log.Infof("service will look for config file at %s", a.Path)

	// Create a template to fill in the variables
	temp, err := template.New("systemDConfig").Parse(systemDConfig)
	if err != nil {
		return "Install could not create template", err
	}

	var newTemplate bytes.Buffer
	if err := temp.Execute(&newTemplate,
		// Execute expects all values to be replaced, adding back in the template variable names for values the daemon library will handle
		&struct {
			Name, Description, Dependencies, Path, Args, User, Group string
		}{
			"{{.Name}}",
			"{{.Description}}",
			"{{.Dependencies}}",
			"{{.Path}}",
			"{{.Args}}",
			a.User,
			a.Group,
		},
	); err != nil {
		return "Install could not populate template", err
	}

	a.service.SetTemplate(newTemplate.String())

	_, err = a.service.Install(a.PathArg, a.Path)

	return "", err
}

func (a *AgentService) serviceEnableReboot() (string, error) {
	// Check the status
	log.Debug("setting the agent to start on reboot")
	status, err := a.service.Status()
	if err != nil {
		return status, err
	}

	output, err := exec.Command("systemctl", "enable", a.Name+".service").Output()
	return string(output), err
}
