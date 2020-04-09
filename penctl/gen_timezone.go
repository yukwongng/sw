// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

//
// usage:
//
// go run gen_timezone.go -output cmd/zoneinfo_windows_to_unix.go
//

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"text/template"
	"time"
)

var filename = flag.String("output", "../cmd/zoneinfo_windows_to_unix.go", "output file name")

// getAbbrs finds timezone abbreviations (standard and daylight saving time)
// for location l.
func getAbbrs(l *time.Location) (st, dt string) {
	t := time.Date(time.Now().Year(), 0, 1, 0, 0, 0, 0, l)
	abbr1, off1 := t.Zone()
	for i := 0; i < 12; i++ {
		t = t.AddDate(0, 1, 0)
		abbr2, off2 := t.Zone()
		if abbr1 != abbr2 {
			if off2-off1 < 0 { // southern hemisphere
				abbr1, abbr2 = abbr2, abbr1
			}
			return abbr1, abbr2
		}
	}
	return abbr1, abbr1
}

type zone struct {
	WinName  string
	UnixName string
	StTime   string
	DSTime   string
}

const wzURL = "https://raw.githubusercontent.com/unicode-org/cldr/master/common/supplemental/windowsZones.xml"

type MapZone struct {
	Other     string `xml:"other,attr"`
	Territory string `xml:"territory,attr"`
	Type      string `xml:"type,attr"`
}

type SupplementalData struct {
	Zones []MapZone `xml:"windowsZones>mapTimezones>mapZone"`
}

func readWindowsZones() ([]*zone, error) {
	r, err := http.Get(wzURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var sd SupplementalData
	err = xml.Unmarshal(data, &sd)
	if err != nil {
		return nil, err
	}
	zs := make([]*zone, 0)
	for _, z := range sd.Zones {
		if z.Territory != "001" {
			// to avoid dups. I don't know why.
			continue
		}
		l, err := time.LoadLocation(z.Type)
		if err != nil {
			return nil, err
		}
		st, dt := getAbbrs(l)
		zs = append(zs, &zone{
			WinName:  z.Other,
			UnixName: z.Type,
			StTime:   st,
			DSTime:   dt,
		})
	}
	return zs, nil
}

func main() {
	flag.Parse()
	zs, err := readWindowsZones()
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(zs, func(i, j int) bool {
		return zs[i].UnixName < zs[j].UnixName
	})
	var v = struct {
		URL string
		Zs  []*zone
	}{
		wzURL,
		zs,
	}
	var buf bytes.Buffer
	err = template.Must(template.New("prog").Parse(prog)).Execute(&buf, v)
	if err != nil {
		log.Fatal(err)
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(*filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

const prog = `
//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

package cmd

var timezones = map[string]string{
{{range .Zs}}	"{{.WinName}}": "{{.UnixName}}",
{{end}}}

`
