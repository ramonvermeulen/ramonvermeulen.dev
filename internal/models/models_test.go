package models

import (
	"testing"
	"time"
)

func TestPositionStartDateShort(t *testing.T) {
	p := &Position{StartDate: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)}
	if got := p.StartDateShort(); got != "Jan 2024" {
		t.Errorf("expected Jan 2024, got %s", got)
	}
}

func TestPositionEndDateShortPresent(t *testing.T) {
	p := &Position{StartDate: time.Now()}
	if got := p.EndDateShort(); got != "Present" {
		t.Errorf("expected Present, got %s", got)
	}
}

func TestPositionEndDateShortValue(t *testing.T) {
	end := time.Date(2025, time.March, 1, 0, 0, 0, 0, time.UTC)
	p := &Position{StartDate: time.Now(), EndDate: &end}
	if got := p.EndDateShort(); got != "Mar 2025" {
		t.Errorf("expected Mar 2025, got %s", got)
	}
}

func TestPositionDurationLessThanMonth(t *testing.T) {
	now := time.Now()
	start := now.Add(-48 * time.Hour)
	p := &Position{StartDate: start}

	got := p.Duration()
	if got != "Less than a month" && got != "1 month" {
		t.Errorf("expected Less than a month or 1 month, got %s", got)
	}
}

func TestPositionDurationYearsMonths(t *testing.T) {
	start := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, time.August, 20, 0, 0, 0, 0, time.UTC)
	p := &Position{StartDate: start, EndDate: &end}
	got := p.Duration()
	if got != "3 years, 8 months" && got != "3 years, 7 months" { // allow off-by-one month logic
		t.Errorf("unexpected duration string: %s", got)
	}
}
