package service

import (
	"site-backend/database/entities"
	"site-backend/database/repository"
)

func CreateForm(form entities.FormEntity) (*entities.FormEntity, error) {
	repo := repository.GetFormRepository()
	return repo.CreateForm(form)
}
