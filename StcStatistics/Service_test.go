package StcStatistics

import (
	"fmt"
	"github.com/cruzzan/stc-intel/StcStatistics/Entities"
	"testing"
)

func TestService_NoClasses(t *testing.T) {
	clubs := []Entities.Club{
		getClub(0, 0),
		getClub(1, 0),
		getClub(7, 7),
		getClub(0, 0),
		getClub(0, 0),
	}
	want := 3

	service := Service{clubs: clubs}
	got, err := service.NoClasses()

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if len(got) != want {
		t.Fatalf("Expected %d clubs to have no classes but got %d", want, len(got))
	}
}

func TestService_ClassCount(t *testing.T) {
	clubs := []Entities.Club{
		getClub(2, 0),
		getClub(1, 0),
		getClub(7, 7),
		getClub(1, 0),
		getClub(0, 0),
	}
	wantCount := 5
	wantClassCount := []int{2, 1, 7, 1, 0}

	service := Service{clubs: clubs}
	got, err := service.ClassCount()

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if len(got) != wantCount {
		t.Fatalf("Expected %d clubs to have no classes but got %d", wantCount, len(got))
	}

	for i, countResult := range got {
		if countResult.ClassCount != wantClassCount[i] {
			t.Fatalf("Wanted %d classes for Club %d, but got %d", wantClassCount[i], i, countResult.ClassCount)
		}
	}
}

func TestService_FullyBookedPercentage(t *testing.T) {
	clubs := []Entities.Club{
		getClub(2, 0),
		getClub(1, 0),
		getClub(7, 7),
		getClub(2, 1),
		getClub(10, 3),
	}
	wantPercentage := []int{0, 0, 100, 50, 30}

	service := Service{clubs: clubs}
	got, err := service.FullyBookedPercentage()
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	for i, percentageResult := range got {
		if percentageResult.FullyBookedPercentage != wantPercentage[i] {
			t.Fatalf("Wanted %d%% for Club %d, but got %d%%", wantPercentage[i], i, percentageResult.FullyBookedPercentage)
		}
	}
}

func getClub(classCount, bookedCount int) Entities.Club {
	var classes []Entities.Class
	for i := 0; i < classCount; i++ {
		spotsLeft := 5
		if bookedCount > 0 {
			spotsLeft = 0
			bookedCount--
		}

		classes = append(classes, Entities.Class{Name: fmt.Sprintf("Class %d", i), AvailableSpots: spotsLeft})
	}

	return Entities.Club{
		Name:    "Lorem ipsum",
		Id:      1337,
		Classes: classes,
	}
}
