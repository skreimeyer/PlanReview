// Package comment templates response letters for grading permit and subdivision
// applications to the Public Works for the City of Little Rock.
package comment

import (
	"os"
	"text/template"
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
}

type geo struct {
	Address string
	Acres   float64
}

type streetClass int

// Street classifications
const (
	RES streetClass = iota
	MINRES
	COLL
	COMM
	MINART
	ART
)

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

type flood struct {
	Class    []FloodHaz
	Floodway bool
}

type zone struct {
	Class    string
	File     string
	Multifam bool
}

/*
functions needed:
gpfee -> $120 + $60/acre
bump -> i++
state -> if ARDOT==true for any
isFH -> if flood.class contains AE/A/Floodway
*/

// Render takes a `main` struct and produces a templated response letter
func Render(m master) {
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
	}
	t, err := template.New("CommentLetter").Funcs(funcMap).ParseFiles("civil.tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
