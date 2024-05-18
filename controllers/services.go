package controllers

import (
	"github/services/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allServices := models.GetAllServices()
	temp.ExecuteTemplate(w, "Index", allServices)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ServiceName := r.FormValue("serviceName")
		Url := r.FormValue("url")

		models.CreateNewService(ServiceName, Url)
	}
	successful := 301
	http.Redirect(w, r, "/", successful)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	serviceId := r.URL.Query().Get("id")
	models.DeleteService(serviceId)
	successful := 301
	http.Redirect(w, r, "/", successful)
}

func GetOnebyId(w http.ResponseWriter, r *http.Request) {
	serviceId := r.URL.Query().Get("id")
	service := models.FindServiceById(serviceId)
	temp.ExecuteTemplate(w, "Update", service)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Id := r.FormValue("id")
		ServiceName := r.FormValue("serviceName")
		Url := r.FormValue("serviceName")

		models.UpdateService(Id, ServiceName, Url)
	}
	successful := 301
	http.Redirect(w, r, "/", successful)
}
