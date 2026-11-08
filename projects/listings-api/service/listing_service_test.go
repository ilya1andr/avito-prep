package service

import (
	"context"
	"testing"

	"github.com/ilya1andr/avito-prep/projects/listings-api/repository"
)

type mockRepo struct{}

func (m *mockRepo) Create(ctx context.Context, title string, price int) (*repository.Listing, error) {
	return &repository.Listing{ID: 1, Title: title, Price: price}, nil
}

func (m *mockRepo) GetByID(ctx context.Context, id int) (*repository.Listing, error) {
	return nil, nil
}

func (m *mockRepo) GetAll(ctx context.Context, minPrice *int, maxPrice *int) ([]repository.Listing, error) {
	return nil, nil
}

func (m *mockRepo) Update(ctx context.Context, id int, title string, price int) (*repository.Listing, error) {
	return nil, nil
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	return nil
}

func TestCreateListing_Success(t *testing.T) {
	mock := &mockRepo{}
	svc := NewListingService(mock)

	listing, err := svc.CreateListing(context.Background(), "iPhone", 50000)

	if err != nil {
		t.Errorf("Ожидали успех, но получили ошибку: %v", err)
	}

	if listing.ID != 1 {
		t.Errorf("Ожидали ID = 1, получили %d", listing.ID)
	}
}
func TestCreateListing_EmptyTitle(t *testing.T) {
	mock := &mockRepo{}
	svc := NewListingService(mock)

	_, err := svc.CreateListing(context.Background(), "   ", 1000)

	if err == nil {
		t.Error("Ожидали ошибку при пустом названии, но её нет")
	}
}

func TestCreateListing_TooLongTitle(t *testing.T) {
	svc := NewListingService(nil)

	longTitle := ""

	for i := 0; i < 101; i++ {
		longTitle += "a"
	}

	_, err := svc.CreateListing(context.Background(), longTitle, 1000)

	if err == nil {
		t.Error("Ожидали ошибку при длинном названии, но её нет")
	}
}
