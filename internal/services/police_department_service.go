package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
	"errors"
	"math"
	"strconv"
)

type PoliceDepartmentService struct {
	Repo *repositories.PoliceDepartmentRepository
}

func (s *PoliceDepartmentService) Create(ctx context.Context, pd models.PoliceDepartment) (models.PoliceDepartment, error) {
	id, err := s.Repo.Create(ctx, pd)
	if err != nil {
		return models.PoliceDepartment{}, err
	}
	pd.ID = id
	return pd, nil
}

func (s *PoliceDepartmentService) GetAll(ctx context.Context) ([]models.PoliceDepartment, error) {
	return s.Repo.GetAll(ctx)
}

func (s *PoliceDepartmentService) GetByID(ctx context.Context, id int) (models.PoliceDepartment, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *PoliceDepartmentService) Update(ctx context.Context, pd models.PoliceDepartment) (models.PoliceDepartment, error) {
	err := s.Repo.Update(ctx, pd)
	if err != nil {
		return models.PoliceDepartment{}, err
	}
	return s.Repo.GetByID(ctx, pd.ID)
}

func (s *PoliceDepartmentService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}

func (s *PoliceDepartmentService) FindNearestPolice(latStr, lonStr string) (models.PoliceDepartment, error) {
	policeList, err := s.Repo.GetAll(context.Background())
	if err != nil || len(policeList) == 0 {
		return models.PoliceDepartment{}, errors.New("no police departments available")
	}

	lat1, _ := strconv.ParseFloat(latStr, 64)
	lon1, _ := strconv.ParseFloat(lonStr, 64)

	minDist := math.MaxFloat64
	var nearest models.PoliceDepartment

	for _, pd := range policeList {
		lat2, _ := strconv.ParseFloat(pd.Latitude, 64)
		lon2, _ := strconv.ParseFloat(pd.Longitude, 64)
		dist := haversine(lat1, lon1, lat2, lon2)
		if dist < minDist {
			minDist = dist
			nearest = pd
		}
	}

	return nearest, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1 *= math.Pi / 180
	lat2 *= math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
