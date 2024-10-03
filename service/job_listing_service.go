package service

import (
	"github.com/appointment-api/graph/model"
	"github.com/appointment-api/repository"
	"github.com/appointment-api/utils"
)

type JobLisingService interface {
	ListAllJobs() ([]*model.JobListing, error)
	GetJobById(id string) (*model.JobListing, error)
	CreateJobListing(request model.CreateJobListingRequest) (*model.JobListing, error)
}

type jobListingService struct {
	store repository.Store
}

func NewJobListingService(store repository.Store) JobLisingService {
	return &jobListingService{
		store: store,
	}
}

func (s *jobListingService) ListAllJobs() ([]*model.JobListing, error) {
	return s.store.ListAllJobs()
}

func (s *jobListingService) GetJobById(id string) (*model.JobListing, error) {
	return s.store.GetJobById(id)
}

func (s *jobListingService) CreateJobListing(request model.CreateJobListingRequest) (*model.JobListing, error) {
	uuid, err := utils.GetUtils().GenerateUUID()
	if err != nil {
		return nil, err
	}
	job := &model.JobListing{
		ID:          uuid,
		Title:       request.Title,
		Description: request.Description,
		Company:     request.Company,
		URL:         request.URL,
	}
	return s.store.CreateJobListing(job)
}
