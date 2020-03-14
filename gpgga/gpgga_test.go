package gpgga

import (
	"reflect"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	type args struct {
		sentence GPGGA
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"",
			args{
				GPGGA{
					Time: time.Date(2020, 3, 3,
						2, 3, 4, 50*1000*1000, time.UTC),
					Latitude:      -44.44,
					Longitude:     55.55,
					FixQuality:    1,
					NumSatellites: 3,
					HDOP:          1.0,
					Altitude:      105,
					Separation:    0,
					DGPSAge:       0,
					DGPSId:        0,
				}},
			"$GPGGA,020304.050,4426.4000,S,05533.0000,E,1,03,1.0,105.0,M,0.0,M,0.0,0000*59",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Build(tt.args.sentence)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		sentence string
		v        *GPGGA
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"",
			args{"$GPGGA,020304.051,4426.4000,S,05533.0000,E,1,03,1.0,105.0,M,0.0,M,0.0,0000*58",
				&GPGGA{
					Time: time.Date(2020, 3, 3,
						2, 3, 4, 51*1000*1000, time.UTC),
					Latitude:      -44.44,
					Longitude:     55.55,
					FixQuality:    1,
					NumSatellites: 3,
					HDOP:          1.0,
					Altitude:      105,
					Separation:    0,
					DGPSAge:       0,
					DGPSId:        0,
				}},
			false,
		},
		{"",
			args{"$GPGGA,215602.966,5230.486,N,01324.171,E,1,12,1.0,0.0,M,0.0,M,,*68",
				&GPGGA{
					Time: time.Date(2020, 3, 3,
						21, 56, 02, 966*1000*1000, time.UTC),
					Latitude:      52.5081,
					Longitude:     13.40285,
					FixQuality:    1,
					NumSatellites: 12,
					HDOP:          1.0,
					Altitude:      0,
					Separation:    0,
					DGPSAge:       0,
					DGPSId:        0,
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &GPGGA{}
			if err := Parse(tt.args.sentence, out); !reflect.DeepEqual(*out, *tt.args.v) || (err != nil) != tt.wantErr {
				t.Errorf("got: %+v, want: %+v", *out, *tt.args.v)
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
