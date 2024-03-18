package service

import (
	"VkTask"
	"VkTask/package/repository"
)

type ActorManagementService struct {
	repo repository.ActorManagement
}

func NewActorService(repo repository.ActorManagement) *ActorManagementService {
	return &ActorManagementService{repo: repo}
}

func (s *ActorManagementService) CreateActor(actor VkTask.Actor) (int, error) {
	return s.repo.CreateActor(actor)
}

func (s *ActorManagementService) UpdateActor(id int, actor VkTask.Actor) error {
	return s.repo.UpdateActor(id, actor)
}

func (s *ActorManagementService) DeleteActor(id int) error {
	return s.repo.DeleteActor(id)
}

func (s *ActorManagementService) GetActorByID(id int) (VkTask.Actor, error) {
	return s.repo.GetActorByID(id)
}

func (s *ActorManagementService) GetAllActors() ([]VkTask.Actor, error) {
	return s.repo.GetAllActors()
}

func (s *ActorManagementService) SearchActors(nameFragment string) ([]VkTask.Actor, error) {
	return s.repo.SearchActors(nameFragment)
}
