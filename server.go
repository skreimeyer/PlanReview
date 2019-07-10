package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/skreimeyer/PlanReview/pkg/comment"
	"github.com/skreimeyer/PlanReview/pkg/esri"
)

//go:generate go run /usr/local/go/src/crypto/tls/generate_cert.go -host "FIXME"

type result struct {
	Letter   string
	Warnings []string
}

func main() {
	// Main page

	tmpl := template.Must(template.ParseFiles("static/index.html", "static/letter.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// Serve static assets

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Request comments

	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		var res result // output to template. Errors will be sent here
		// Boolean values
		sub, _ := strconv.ParseBool(r.FormValue("Sub"))
		apr, _ := strconv.ParseBool(r.FormValue("Approved"))
		gp, _ := strconv.ParseBool(r.FormValue("GP"))
		fr, _ := strconv.ParseBool(r.FormValue("Franchise"))
		stm, _ := strconv.ParseBool(r.FormValue("Storm"))
		wl, _ := strconv.ParseBool(r.FormValue("Wall"))
		meta := comment.Meta{
			Sub:         sub,
			AppName:     r.FormValue("AppName"),
			AppTitle:    r.FormValue("AppTitle"),
			AppCompany:  r.FormValue("AppCompany"),
			AppAdd:      r.FormValue("AppAdd"),
			AppCSZ:      r.FormValue("AppCSZ"),
			ProjectName: r.FormValue("ProjectName"),
			Approved:    apr,
			GP:          gp,
			Franchise:   fr,
			Storm:       stm,
			Wall:        wl,
		}
		// Check Address for being PID-like
		match, err := regexp.MatchString(`\d{2}L\d+`, r.FormValue("Address"))
		if err != nil {
			res.Warnings = append(res.Warnings, "Regex test failed.", err.Error())
		}
		//init
		var gc esri.Location
		var par esri.PResponse
		if match {
			par, err = esri.FetchByPID(r.FormValue("Address"))
			if err != nil {
				res.Warnings = append(res.Warnings, "Parcel lookup failed.", err.Error())
			}
			ring := esri.GetRing(par)
			x, y := 0.0, 0.0
			for i := range ring {
				x += ring[i][0]
				y += ring[i][1]
			}
			x = x / float64(len(ring))
			y = y / float64(len(ring))
			gc = esri.Location{X: x, Y: y}
		} else {
			// Geocode
			gc, err = esri.Geocode(r.FormValue("Address")) //esri.Location
			if err != nil {
				res.Warnings = append(res.Warnings, "Failed to geocode.", err.Error())
			}
			par, err = esri.FetchParcel(gc) //esri.Parcel
			if err != nil {
				res.Warnings = append(res.Warnings, "Failed to fetch parcel data.", err.Error())
			}
		}
		acr := 0.0
		if len(par.Features) >= 1 {
			acr = par.Features[0].Attributes.CalcAcre
		}
		env := esri.MakeEnvelope(esri.GetRing(par), 0.05)
		streets, err := esri.FetchRoads(env)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to fetch street data.", err.Error())
		}
		geo := comment.Geo{
			Address: r.FormValue("Address"),
			Acres:   acr,
		}
		fldzn, err := esri.FloodData(env) // []FloodHaz
		fldzn = esri.Unique(fldzn)
		fldwy := esri.InFloodway(fldzn)
		flood := comment.Flood{
			Class:    fldzn,
			Floodway: fldwy,
		}
		zCode, err := esri.FetchZone(gc)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to fetch zoning.", err.Error())
		}
		zNum, err := esri.FetchCases(env)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to find case number.", err.Error())
		}
		mf := esri.IsMultifam(zCode)
		zone := comment.Zone{
			Class:    zCode,
			File:     zNum,
			Multifam: mf,
		}
		master := comment.Master{
			Meta:   meta,
			Geo:    geo,
			Street: streets,
			Flood:  flood,
			Zone:   zone,
		}
		// write template
		letter, err := comment.Render(master)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to write template.", err.Error())
			fmt.Println(err)
		}
		res.Letter = letter
		tmpl.ExecuteTemplate(w, "letter.html", res)
	})

	// Serve
	port := os.Getenv("PORT")
	// port := "8080" // uncomment for testing
	fmt.Println("serving...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
