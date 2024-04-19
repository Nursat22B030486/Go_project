package model

import (
	_ "fmt"
	"math"
	"strings"

	"github.com/Nursat22B030486/go_project/pkg/read-it/validator"
)

// Define a new Metadata struct for holding the pagination metadata.
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalrecords, page, page_size int) Metadata {
	if totalrecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     page_size,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalrecords) / float64(page_size))),
		TotalRecords: totalrecords,
	}
}

type Filters struct {
	Page         int
	Page_size    int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than 0")
	v.Check(f.Page <= 10_000_000, "page", "must be maximum of 10 million")
	v.Check(f.Page_size > 0, "page_size", "must be greater than 0")
	v.Check(f.Page_size <= 100, "page_size", "must be maximum of 100")

	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

// Check that the client-provided Sort field matches one of the entries in our safelist
// and if it does, extract the column name from the Sort field by stripping the leading
// hyphen character (if one exists).
func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("Unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) limit() int {
	return f.Page_size
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.Page_size
}
