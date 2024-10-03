package mariadb

import (
	"fmt"

	"github.com/appointment-api/graph/model"
)

func (s *Store) ListAllJobs() ([]*model.JobListing, error) {
	var jobs []*model.JobListing
	// Perform the query to retrieve all job listings
	if err := s.db.Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch job listings: %v", err)
	}
	return jobs, nil
}

func (s *Store) GetJobById(id string) (*model.JobListing, error) {
	var job *model.JobListing
	// Perform the query to retrieve all job listings
	if err := s.db.Find(&job).Where(`id=?`, id).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch job listings: %v", err)
	}
	return job, nil
}

func (s *Store) CreateJobListing(request *model.JobListing) (*model.JobListing, error) {
	err := s.db.Create(request).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}
