package repository

import "github.com/appointment-api/graph/model"

type Store interface {
	ListAllJobs() ([]*model.JobListing, error)
	GetJobById(id string) (*model.JobListing, error)
	CreateJobListing(request *model.JobListing) (*model.JobListing, error)
}
