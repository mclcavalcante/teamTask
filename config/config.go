package config

import (
	"github.com/mclcavalcante/teamTask/controller"
	service "github.com/mclcavalcante/teamTask/services"
)

type Initialization struct {
	Repo       service.Repository
	Svc        service.Service
	Controller controller.Controller
}

func NewInitialization(repo service.Repository, svc service.Service, controller controller.Controller) *Initialization {
	return &Initialization{
		Repo:       repo,
		Svc:        svc,
		Controller: controller,
	}
}
