package service

import (
	"fmt"
	"os/exec"
	"reflect"

	"github.com/guestc/golib/file"
)

type ServiceContent struct {
	Unit struct {
		Description   string
		Documentation string
		Requires      string
		After         string
		Before        string
	}
	Service struct {
		Type                   string
		ExecStart              string
		ExecReload             string
		ExecStop               string
		ExecStartPre           string
		ExecStartPost          string
		ExecStopPre            string
		ExecStopPost           string
		Restart                string
		RestartSec             string
		TimeoutStartSec        string
		TimeoutStopSec         string
		TimeoutSec             string
		RuntimeMaxSec          string
		WatchdogSec            string
		StartLimitInterval     string
		StartLimitBurst        string
		StartLimitAction       string
		FailureAction          string
		PermissionsStartOnly   string
		RootDirectoryStartOnly string
		NonBlocking            string
		NotifyAccess           string
		WorkingDirectory       string
		User                   string
		Group                  string
		PIDFile                string
	}
	Install struct {
		WantedBy   string
		RequiredBy string
	}
}

// stop service
func StopService(serviceName string) bool {
	cmd := exec.Command("systemctl", "stop", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to stop service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// start service
func StartService(serviceName string) bool {
	cmd := exec.Command("systemctl", "start", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to start service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// reload service
func ReloadService(serviceName string) bool {
	cmd := exec.Command("systemctl", "reload", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to reload service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// reload service config
func ReloadServiceConfig() bool {
	cmd := exec.Command("systemctl", "daemon-reload")
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to reload service config" + "\nerror:" + err.Error())
		return false
	}
	return true
}

// enable service
func EnableService(serviceName string) bool {
	cmd := exec.Command("systemctl", "enable", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to enable service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// disable service
func DisableService(serviceName string) bool {
	cmd := exec.Command("systemctl", "disable", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to disable service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// restart service
func RestartService(serviceName string) bool {
	cmd := exec.Command("systemctl", "restart", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to restart service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// status service if service not running return false
func StatusService(serviceName string) bool {
	cmd := exec.Command("systemctl", "status", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to get status of service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// setup new service to system
func SetupService(serviceName string, serviceContent string) bool {
	servicePath := "/etc/systemd/system/" + serviceName + ".service"
	if !file.WriteString(servicePath, serviceContent) {
		fmt.Println("failed to setup service :" + serviceName)
		return false
	}
	return true
}

// remove service from system
func RemoveService(serviceName string) bool {
	if StatusService(serviceName) {
		fmt.Println("service :" + serviceName + " can't be stopped.")
		return false
	}
	servicePath := "/etc/systemd/system/" + serviceName + ".service"
	cmd := exec.Command("rm", servicePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to remove service :" + serviceName + "\nerror:" + err.Error())
		return false
	}
	return true
}

// setup new service to system more details
func SetupServiceDetail(serviceName string, serviceContent ServiceContent) bool {
	buildServiceContent := "[Unit]"
	for i, t, v := 0, reflect.TypeOf(serviceContent.Unit), reflect.ValueOf(serviceContent.Unit); i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			buildServiceContent += "\n" + t.Field(i).Name + "=" + v.Field(i).String()
		}
	}

	buildServiceContent += "\n[Service]"
	for i, t, v := 0, reflect.TypeOf(serviceContent.Service), reflect.ValueOf(serviceContent.Service); i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			buildServiceContent += "\n" + t.Field(i).Name + "=" + v.Field(i).String()
		}
	}

	buildServiceContent += "\n[Install]"
	for i, t, v := 0, reflect.TypeOf(serviceContent.Install), reflect.ValueOf(serviceContent.Install); i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			buildServiceContent += "\n" + t.Field(i).Name + "=" + v.Field(i).String()
		}
	}
	if SetupService(serviceName, buildServiceContent) {
		return true
	}
	fmt.Println("failed to setup service :" + serviceName)
	return false
}

// setup new service to system simple
func SetupServiceSimple(serviceName string, execStart string, desc string) bool {
	serviceContent := ServiceContent{}
	serviceContent.Unit.Description = desc
	serviceContent.Service.ExecStart = execStart
	serviceContent.Install.WantedBy = "multi-user.target"
	return SetupServiceDetail(serviceName, serviceContent)
}
