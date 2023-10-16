package StcStatistics

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cruzzan/stc-intel/StcStatistics/Entities"
	"math"
	"net/http"
	"net/url"
	"time"
)

type Service struct {
	baseUrl string
	client  *http.Client
	clubs   []Entities.Club
}

// NoClasses Finds all the clubs that offer no classes
func (s *Service) NoClasses() ([]string, error) {
	if len(s.clubs) == 0 {
		err := s.fetchClubData()
		if err != nil {
			return nil, fmt.Errorf("could not build Club cache, got error: %w", err)
		}
	}

	var res []string
	for _, club := range s.clubs {
		if club.CountClasses() == 0 {
			res = append(res, club.Name)
		}
	}

	return res, nil
}

type ClassCount struct {
	Club       string `json:"Club"`
	ClassCount int    `json:"ClassCount"`
}

// ClassCount Calculates the number of classes available at each Club
func (s *Service) ClassCount() ([]ClassCount, error) {
	if len(s.clubs) == 0 {
		err := s.fetchClubData()
		if err != nil {
			return nil, fmt.Errorf("could not build Club cache, got error: %w", err)
		}
	}

	var res []ClassCount
	for _, club := range s.clubs {
		res = append(res, ClassCount{Club: club.Name, ClassCount: club.CountClasses()})
	}

	return res, nil
}

type FullyBookedPercentage struct {
	Club                  string `json:"Club"`
	FullyBookedPercentage int    `json:"FullyBookedPercentage"`
}

// FullyBookedPercentage calculates the percentage of fully booked classes for each Club
func (s *Service) FullyBookedPercentage() ([]FullyBookedPercentage, error) {
	if len(s.clubs) == 0 {
		err := s.fetchClubData()
		if err != nil {
			return nil, fmt.Errorf("could not build Club cache, got error: %w", err)
		}
	}

	var res []FullyBookedPercentage
	for _, club := range s.clubs {
		if club.CountClasses() > 0 {
			percentage := math.Round((float64(club.CountFullyBookedClasses()) / float64(club.CountClasses())) * 100)
			res = append(res, FullyBookedPercentage{Club: club.Name, FullyBookedPercentage: int(percentage)})
		}
	}

	return res, nil
}

func (s *Service) fetchClubData() error {
	var err error
	s.clubs, err = s.fetchAllClubs()
	if err != nil {
		return fmt.Errorf("failed to fetch Club data %w", err)
	}

	from := time.Now()
	to := from.Add(7 * 24 * time.Hour)
	for i, club := range s.clubs {
		s.clubs[i].Classes, err = s.fetchClubClasses(club, from, to)
		if err != nil {
			return fmt.Errorf("failed to fetch classes for Club %s %w", club.Name, err)
		}
	}

	return nil
}

func (s *Service) fetchAllClubs() ([]Entities.Club, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/businessunits", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create fetch all clubs request %w", err)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute fetch all clubs request %w", err)
	}

	if res.StatusCode == http.StatusOK {
		var clubs []Entities.Club
		err := json.NewDecoder(res.Body).Decode(&clubs)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshall the fetch all clubs response %w", err)
		}

		return clubs, nil
	} else {
		return nil, errors.New(fmt.Sprintf("got undesired http status %d - %s while fetching all clubs", res.StatusCode, res.Status))
	}
}

func (s *Service) fetchClubClasses(club Entities.Club, from, to time.Time) ([]Entities.Class, error) {
	format := "2006-01-02T15:04:05.000Z"
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/businessunits/%d/groupactivities?period.start=%s&period.end=%s",
			s.baseUrl,
			club.Id,
			url.QueryEscape(from.Format(format)),
			url.QueryEscape(to.Format(format)),
		),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create fetch Club classes request %w", err)
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute fetch Club classes request %w", err)
	}

	if res.StatusCode == http.StatusOK {
		var classes []Entities.Class
		err := json.NewDecoder(res.Body).Decode(&classes)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshall the fetch Club classes response %w", err)
		}

		return classes, nil
	} else {
		return nil, errors.New(fmt.Sprintf("got undesired http status %d - %s while fetching classes for %s", res.StatusCode, res.Status, club.Name))
	}
}

func NewService(httpClient *http.Client) Service {
	return Service{
		baseUrl: "https://stc.brpsystems.com/brponline/api/ver3",
		client:  httpClient,
	}
}
