package vec3

import (
	"math"
	"reflect"
	"testing"
)

func TestVec3_Plus(t *testing.T) {
	const d float64 = 0.00000001
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		v *Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		{
			name:   "add two vector",
			fields: fields{0.1, 0.2, 0.5},
			args:   args{&Vec3{0.1, 0.3, 0.6}},
			want:   &Vec3{0.2, 0.5, 1.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Plus(tt.args.v)
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.Plus() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestVec3_Inverse(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vec3
	}{
		{
			name:   "inverse vector",
			fields: fields{0.4, -0.6, 0.6},
			want:   &Vec3{-0.4, 0.6, -0.6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := a.Inverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Scale(t *testing.T) {
	const d float64 = 0.00000001
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		f float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		{
			name:   "scale vector",
			fields: fields{0.4, -0.6, 0.6},
			args:   args{3},
			want:   &Vec3{1.2, -1.8, 1.8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Scale(tt.args.f)
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.Scale() = %v, want %v", got, tt.want)
			}

		})
	}
}
