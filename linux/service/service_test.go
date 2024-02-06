package service_test
import (
	"testing"
	"github.com/guestc/golib/linux/service"
)

func Test(t *testing.T){
	service.StopService("nginx")
	service.StartService("nginx")
	service.ReloadService("nginx")
	service.ReloadServiceConfig("nginx")
	service.EnableService("nginx")
	service.DisableService("nginx")
	service.SetupServiceSimple("nginx", "nginx", "nginx")
}