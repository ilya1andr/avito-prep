package service

import (
	"context"
	"errors"

	"github.com/ilya1andr/avito-prep/projects/listings-api/repository"
)

type ListingService struct {
	repo *repository.ListingRepository
}

func NewListingService(repo *repository.ListingRepository) *ListingService {
	return &ListingService{repo: repo}
}

func (s *ListingService) CreateListing(ctx context.Context, title string, price int) (*repository.Listing, error) {
	if title == "" {
		return nil, errors.New("название не может быть пустым")
	}
	if price <= 0 {
		return nil, errors.New("цена должна быть больше нуля")
	}

	return s.repo.Create(ctx, title, price)
}
