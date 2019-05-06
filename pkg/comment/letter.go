// Package comment templates response letters for grading permit and subdivision
// applications to the Public Works for the City of Little Rock.
package comment

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type master struct {
	Meta   meta
	Geo    geo
	Street []street
	Flood  flood
	Zone   zone
}

type meta struct {
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

type geo struct {
	Address string
	Acres   float64
}

type streetClass int

// Street classifications
const (
	Residential streetClass = iota
	MinorResidential
	Collector
	Commercial
	MinorArterial
	Arterial
)

//go:generate stringer -type=streetClass

type street struct {
	Name  string
	Class streetClass
	Row   int
	Alt   bool
	ARDOT bool
}

// FloodHaz is an enumeration of valid flood hazard area designations. Its use
// is an alternative to a hashmap
type FloodHaz int

// Flood Hazard Area classifications
const (
	X    FloodHaz = iota
	FIVE          // 0.2% annual chance
	A
	AE
	FLOODWAY
	LEVEE
)

//go:generate stringer -type=FloodHaz

type flood struct {
	Class    []FloodHaz
	Floodway bool
}

type zone struct {
	Class    string
	File     string
	Multifam bool
}

// Render takes a `main` struct and produces a templated response letter
func Render(m master) error {
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
		"state": func(streets []street) bool {
			for _, s := range streets {
				if s.ARDOT {
					return true
				}
			}
			return false
		},
		"isFH": func(f flood) bool {
			for _, z := range f.Class {
				if z == A || z == AE || z == FLOODWAY {
					return true
				}
			}
			return false
		},
		"today": func() string {
			return fmt.Sprint(time.Now().Date())
		},
	}
	t, err := template.New("civil.tmpl").Funcs(funcMap).ParseFiles("civil.tmpl")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("../../tmp/letter")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = t.Execute(f, m)
	if err != nil {
		panic(err)
	}
	return nil
}
