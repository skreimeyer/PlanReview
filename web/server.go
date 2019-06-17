package main

import (
	"fmt"
	"github.com/skreimeyer/PlanReview/pkg/comment"
	"github.com/skreimeyer/PlanReview/pkg/esri"
	"html/template"
	"net/http"
	"strconv"
)

//go:generate go run /usr/local/go/src/crypto/tls/generate_cert.go -host "FIXME"

type result struct {
	Letter   string
	Warnings []string
}

func main() {
	// Main page

	tmpl := template.Must(template.ParseFiles("index.html", "letter.html"))

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
		// Geocode
		gc := esri.Geocode(r.FormValue("Address")) //esri.Location
		par, err := esri.FetchParcel(gc)           //esri.Parcel
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to fetch parcel data.")
		}
		env := esri.MakeEnvelope(esri.GetRing(par), 0.05)
		streets, err := esri.FetchRoads(env)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to fetch street data.")
		}
		geo := comment.Geo{
			Address: r.FormValue("Address"),
			Acres:   par.Features[0].Attributes.CalcAcre,
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
			res.Warnings = append(res.Warnings, "Failed to fetch zoning.")
		}
		zNum, err := esri.FetchCases(env)
		if err != nil {
			res.Warnings = append(res.Warnings, "Failed to find case number.")
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
			res.Warnings = append(res.Warnings, "Failed to write template.")
			fmt.Println(err)
		}
		res.Letter = letter
		tmpl.ExecuteTemplate(w, "letter.html", res)
	})

	// Serve

	fmt.Println("serving...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
