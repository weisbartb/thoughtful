package thoughtful_test

import (
	"github.com/weisbartb/thoughtful"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		width  float64
		height float64
		length float64
		mass   float64
	}
	tests := []struct {
		name      string
		args      args
		wantLabel string
	}{
		{
			name: "Bulky width",
			args: args{
				width:  150,
				height: 1,
				length: 1,
				mass:   1,
			},
			wantLabel: thoughtful.PackageLabelSpecial,
		},
		{
			name: "Bulky height",
			args: args{
				width:  1,
				height: 150,
				length: 1,
				mass:   1,
			},
			wantLabel: thoughtful.PackageLabelSpecial,
		},
		{
			name: "Bulky length",
			args: args{
				width:  1,
				height: 1,
				length: 150,
				mass:   1,
			},
			wantLabel: thoughtful.PackageLabelSpecial,
		},
		{
			name: "Rejected dimension",
			args: args{
				width:  0,
				height: 1,
				length: 150,
				mass:   20,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "Bulky dimension",
			args: args{
				width:  100,
				height: 100,
				length: 100,
				mass:   14,
			},
			wantLabel: thoughtful.PackageLabelSpecial,
		},
		{
			name: "Bulky weight",
			args: args{
				width:  99,
				height: 100,
				length: 100,
				mass:   20,
			},
			wantLabel: thoughtful.PackageLabelSpecial,
		},
		{
			name: "interdimensional package",
			args: args{
				width:  99,
				height: 100,
				length: 100,
				mass:   -1,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "interdimensional package",
			args: args{
				width:  -1,
				height: 100,
				length: 100,
				mass:   20,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "interdimensional package",
			args: args{
				width:  99,
				height: -1,
				length: 100,
				mass:   20,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "interdimensional package",
			args: args{
				width:  99,
				height: 100,
				length: -1,
				mass:   20,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "overflow",
			args: args{
				width:  math.MaxInt64,
				height: math.MaxInt64,
				length: math.MaxInt64,
				mass:   math.MaxInt64,
			},
			wantLabel: thoughtful.PackageLabelRejected,
		},
		{
			name: "normal",
			args: args{
				width:  20,
				height: 20,
				length: 20,
				mass:   5,
			},
			wantLabel: thoughtful.PackageLabelStandard,
		},
		{
			name: "normal - barely",
			args: args{
				width:  100,
				height: 100,
				length: 99.99,
				mass:   5,
			},
			wantLabel: thoughtful.PackageLabelStandard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLabel := thoughtful.Sort(tt.args.width, tt.args.height, tt.args.length, tt.args.mass); gotLabel != tt.wantLabel {
				t.Errorf("Sort() = %v, want %v", gotLabel, tt.wantLabel)
			}
		})
	}
}
