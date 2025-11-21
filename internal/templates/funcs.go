package templates

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

// buildFuncMap constructs the FuncMap used in templates. Keeping separate
func buildFuncMap() template.FuncMap {
	return template.FuncMap{
		"sub":            func(a, b int) int { return a - b },
		"mod":            func(a, b int) int { return a % b },
		"add":            func(a, b int) int { return a + b },
		"dateFormat":     dateFormat,
		"durationFormat": durationFormat,
	}
}

func dateFormat(t interface{}) string {
	if t == nil {
		return ""
	}
	switch v := t.(type) {
	case *time.Time:
		if v == nil || v.IsZero() {
			return ""
		}
		return v.Format("Jan 2006")
	case time.Time:
		if v.IsZero() {
			return ""
		}
		return v.Format("Jan 2006")
	default:
		return ""
	}
}

func durationFormat(start time.Time, end *time.Time) string {
	var endTime time.Time
	if end == nil || end.IsZero() {
		endTime = time.Now()
	} else {
		endTime = *end
	}
	years, months := diff(start, endTime)
	if years < 0 || months < 0 {
		return "Less than a month"
	}
	var parts []string
	if years > 0 {
		if years == 1 {
			parts = append(parts, "1 year")
		} else {
			parts = append(parts, fmt.Sprintf("%d years", years))
		}
	}
	if months > 0 {
		if months == 1 {
			parts = append(parts, "1 month")
		} else {
			parts = append(parts, fmt.Sprintf("%d months", months))
		}
	}
	if len(parts) == 0 {
		return "Less than a month"
	}
	return strings.Join(parts, ", ")
}

// diff computes the difference in years and remaining months between two times.
// It assumes a <= b ordering.
func diff(a, b time.Time) (years, months int) {
	if a.After(b) {
		a, b = b, a
	}
	years = b.Year() - a.Year()
	months = int(b.Month()) - int(a.Month())
	if b.Day() >= a.Day() {
		months++
	}
	if months < 0 {
		years--
		months += 12
	}
	if months >= 12 {
		years += months / 12
		months = months % 12
	}
	return
}
