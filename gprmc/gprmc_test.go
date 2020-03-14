package gprmc

import (
	"reflect"
	"testing"
	"time"
)

func Test_buildgprmc(t *testing.T) {
	type args struct {
		sentence GPRMC
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		/*{
			"",
			args{
				GPRMC{
					DateTime: time.Date(98, 9, 13, 8, 18, 36, 0, time.UTC),
				},
			},
			"$GPRMC,081836,A,3751.65,S,14507.36,E,000.0,360.0,130998,011.3,E*62",
			false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Build(tt.args.sentence)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildgprmc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildgprmc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsegprmc(t *testing.T) {
	type args struct {
		sentence string
		v        *GPRMC
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"",
			args{
				"$GPRMC,081836,A,3751.63,S,14507.35,E,000.0,360.0,130998,011.3,E*67",
				&GPRMC{
					DateTime:          time.Date(1998, 9, 13, 8, 18, 36, 0, time.UTC),
					ReceiverStatus:    "A",
					Latitude:          -37.8605,
					Longitude:         145.1225,
					Speed:             0,
					Course:            360,
					MagneticVariation: 11.3,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &GPRMC{}
			if err := Parse(tt.args.sentence, out); !reflect.DeepEqual(*out, *tt.args.v) || (err != nil) != tt.wantErr {
				t.Errorf("got: %+v, want: %+v", *out, *tt.args.v)
				t.Errorf("parsegprmc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
