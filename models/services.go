package models

import (
	"github/services/database"
	"log"
)

type HealthCheckServices struct {
	Id          int
	ServiceName string
	Url         string
	Status      string
	LastCheck   string
}

func GetAllServices() []HealthCheckServices {
	db := database.DatabaseConnect()

	allServices, err := db.Query("select * from services")
	if err != nil {
		log.Fatal(err)
	}

	service := HealthCheckServices{}
	services := []HealthCheckServices{}

	for allServices.Next() {
		var id int
		var serviceName, url, status, lastCheck string

		err = allServices.Scan(&id, &serviceName, &url, &status, &lastCheck)
		if err != nil {
			log.Fatal(err)
		}

		service.Id = id
		service.Url = url
		service.ServiceName = serviceName
		service.Status = status
		service.LastCheck = lastCheck

		services = append(services, service)
	}
	defer db.Close()
	return services
}

func CreateNewService(serviceName, url string) {
	db := database.DatabaseConnect()

	newService, err := db.Prepare("INSERT INTO services(\"serviceName\",url, status, \"lastCheck\") values ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	newService.Exec(serviceName, url, "NO DATA", "NO DATA")
	defer db.Close()
}

func DeleteService(id string) {
	db := database.DatabaseConnect()
	deleteService, err := db.Prepare("delete from services where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	deleteService.Exec(id)
	defer db.Close()
}

func FindServiceById(id string) HealthCheckServices {
	db := database.DatabaseConnect()
	getService, err := db.Query("select * from services where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	service := HealthCheckServices{}
	for getService.Next() {
		var id int
		var serviceName, url, status, lastCheck string
		err = getService.Scan(&id, &serviceName, &url, &status, &lastCheck)
		if err != nil {
			log.Fatal(err)
		}
		service.Id = id
		service.Url = url
		service.ServiceName = serviceName
		service.Status = status
		service.LastCheck = lastCheck
	}
	defer db.Close()
	return service
}

func UpdateService(id, serviceName, url string) {
	db := database.DatabaseConnect()
	updatedService, err := db.Prepare("update services set \"serviceName\"=$1, url=$2 where id=$3")
	if err != nil {
		log.Fatal(err)
	}
	updatedService.Exec(serviceName, url, id)
	defer db.Close()
}
