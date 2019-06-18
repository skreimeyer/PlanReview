// Package comment templates response letters for grading permit and subdivision
// applications to the Public Works for the City of Little Rock.
package comment

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/skreimeyer/PlanReview/pkg/esri"
)

// Master is the struct passed to the template. It hold other structs for
// the sake of brevity.
type Master struct {
	Meta   Meta
	Geo    Geo
	Street []esri.Street
	Flood  Flood
	Zone   Zone
}

// Meta contains all information about the application itself.
type Meta struct {
	Sub         bool
	AppName     string
	AppTitle    string
	AppCompany  string
	AppAdd      string
	AppCSZ      string
	ProjectName string
	Approved    bool
	GP          bool
	Franchise   bool
	Storm       bool
	Wall        bool
}

// Geo is a container for geocoding information.
type Geo struct {
	Address string
	Acres   float64
}

// Flood contains a list of flood hazard area designations
type Flood struct {
	Class    []esri.FloodHaz
	Floodway bool
}

// Zone defines the zoning code of a parcel
type Zone struct {
	Class    string
	File     string
	Multifam bool
}

// Render takes a `master` struct and produces a templated response letter
func Render(m Master) (string, error) {
	funcMap := template.FuncMap{
		"bump": func(i int) int {
			return i + 1
		},
		"gpfee": func(a float64) float64 {
			switch {
			case a < 0.5:
				return 60.0
			case a < 1.0:
				return 120.0
			case a > 10.0:
				return 660.0
			default:
				return 120 + 60*(a-1)
			}
		},
		"state": func(streets []esri.Street) bool {
			for _, s := range streets {
				if s.ARDOT {
					return true
				}
			}
			return false
		},
		"isFH": func(f Flood) bool {
			for _, z := range f.Class {
				if z == esri.A || z == esri.AE || z == esri.FLOODWAY {
					return true
				}
			}
			return false
		},
		"today": func() string {
			return fmt.Sprint(time.Now().Date())
		},
	}
	t, err := template.New("civil.gotmpl").Funcs(funcMap).ParseFiles("civil.gotmpl")
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, m)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
