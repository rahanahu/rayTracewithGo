package ray

import (
	"goRayTrace/vec3"
	"reflect"
	"testing"
)

func TestRay_At(t *testing.T) {
	type fields struct {
		Origin    vec3.Vec3
		Direction vec3.Vec3
	}
	type args struct {
		t float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *vec3.Vec3
	}{
		{
			name: "at",
			fields: fields{
				Origin:    vec3.Vec3{X: 1, Y: 2, Z: 3},
				Direction: vec3.Vec3{X: 1, Y: 0, Z: 0}},
			args: args{1},
			want: &vec3.Vec3{X: 2, Y: 2, Z: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			if got := a.At(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ray.At() = %v, want %v", got, tt.want)
			}
		})
	}
}
