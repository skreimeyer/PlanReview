package comment

import "testing"

func TestRender(t *testing.T) {
	type args struct {
		m Master
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			args: args{m: Master{
				Meta: Meta{
					Sub:         true,
					AppName:     "Mr. Gerald Dingalingus, P.E.,",
					AppTitle:    "Mr. Dingus",
					AppCompany:  "ABC Architects",
					AppAdd:      "100 1st Street",
					AppCSZ:      "Little Rock, AR 72201",
					ProjectName: "A PROJECT",
					Approved:    true,
					GP:          false,
					Franchise:   false,
					Storm:       false,
					Wall:        false,
				},
				Geo: Geo{
					Address: "1 Project Lane",
					Acres:   2.5,
				},
				Street: []Street{
					Street{
						Name:  "Project Lane",
						Class: Residential,
						Row:   25,
						Alt:   false,
						ARDOT: false,
					},
				},
				Flood: Flood{
					Class:    []FloodHaz{X},
					Floodway: false,
				},
				Zone: Zone{
					Class:    "R2",
					File:     "Z-12345",
					Multifam: false,
				},
			}},
			wantErr: true, // make true to write template letter in tmp
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Render(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
