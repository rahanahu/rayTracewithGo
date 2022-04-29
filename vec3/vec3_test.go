package vec3

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

const d float64 = 0.00000001

func TestVec3_Plus(t *testing.T) {

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

func TestVec3_Minus(t *testing.T) {
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
			name:   "sub two vector",
			fields: fields{0.1, 0.2, 0.5},
			args:   args{&Vec3{0.1, 0.3, 0.6}},
			want:   &Vec3{0, -0.1, -0.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Minus(tt.args.v)
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.Minus() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestVec3_LengthSquared(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "square length",
			fields: fields{3, 4, 5},
			want:   50,
		},
		{
			name:   "square length",
			fields: fields{3.1, 4.5, 5.3},
			want:   57.95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.LengthSquared()
			if !(math.Abs(got-tt.want) < d) {
				t.Errorf("Vec3.LengthSquared() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestVec3_Length(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "square length",
			fields: fields{3, 4, 5},
			want:   math.Sqrt(50),
		},
		{
			name:   "square length",
			fields: fields{3.1, 4.5, 5.3},
			want:   math.Sqrt(57.95),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Length()
			if !(math.Abs(got-tt.want) < d) {
				t.Errorf("Vec3.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Times(t *testing.T) {
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
			name:   "multiple two vector",
			fields: fields{0.1, 0.2, 0.5},
			args:   args{&Vec3{0.1, 0.3, 0.6}},
			want:   &Vec3{0.01, 0.06, 0.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Times(tt.args.v)
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.Times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Dot(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Dot1",
			fields: fields{3, 4, 5},
			want:   50,
		},
		{
			name:   "Dot2",
			fields: fields{3.1, 4.5, 5.3},
			want:   57.95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Dot()
			if !(math.Abs(got-tt.want) < d) {
				t.Errorf("Vec3.Dot() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestVec3_Cross(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		t *Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		{
			name:   "cross two vector",
			fields: fields{3, 4, 1},
			args:   args{&Vec3{3, 7, 5}},
			want:   &Vec3{13, -12, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.Cross(tt.args.t)
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_UnitVector(t *testing.T) {
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
			name:   "unit vector 1",
			fields: fields{1, 2, 3},
			want:   &Vec3{1 / 3.741657386774, 2 / 3.741657386774, 3 / 3.741657386774},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			got := a.UnitVector()
			if !((math.Abs(got.X-tt.want.X) < d) &&
				(math.Abs(got.Y-tt.want.Y) < d) &&
				(math.Abs(got.Z-tt.want.Z) < d)) {
				t.Errorf("Vec3.UnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_WriteColor(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		{
			name:    "write color",
			fields:  fields{0.4, 0.2, 0.4},
			wantOut: "102 51 102¥n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			out := &bytes.Buffer{}
			a.WriteColor(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Vec3.WriteColor() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
