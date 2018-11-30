package classifier

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     bool
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(1.0, 1.0, 1.0),
			want:     true,
		},
		{
			name:     "case-2",
			triangle: NewTriangle(1.0, 1.0, 0.0),
			want:     false,
		},
		{
			name:     "case-3",
			triangle: NewTriangle(-1.0, 1.0, 1.0),
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isValid()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("triangle is valid? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCanFormTriangle(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     bool
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(1.0, 1.0, 1.0),
			want:     true,
		},
		{
			name:     "case-2",
			triangle: NewTriangle(1.0, 1.0, 0.0),
			want:     false,
		},
		{
			name:     "case-3",
			triangle: NewTriangle(-1.0, 1.0, 1.0),
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.canFormTriangle()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("can triangle be formed? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsEquilateral(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     string
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(21.0, 21.0, 21.0),
			want:     "Equilateral",
		},
		{
			name:     "case-2",
			triangle: NewTriangle(21.0, 22.0, 21.0),
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isEquilateral()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("is triangle Equilateral? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsIsosceles(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     string
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(21.0, 22.0, 21.0),
			want:     "Isosceles",
		},
		{
			name:     "case-2",
			triangle: NewTriangle(3.0, 4.0, 6.0),
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isIsosceles()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("is triangle Isosceles? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsRight(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     bool
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(1.0, 1.0, 1.0),
			want:     false,
		},
		{
			name:     "case-2",
			triangle: NewTriangle(3.0, 4.0, 6.0),
			want:     false,
		},
		{
			name:     "case-3",
			triangle: NewTriangle(18, 80, 82),
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isRight()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("is triangle Right? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsScaleneOrRightScalene(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     string
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(18, 80, 82),
			want:     "Right Scalene",
		},
		{
			name:     "case-2",
			triangle: NewTriangle(3.0, 4.0, 6.0),
			want:     "Scalene",
		},
		{
			name:     "case-3",
			triangle: NewTriangle(1, 1, 1),
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isScaleneOrRightScalene()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("is triangle Scalene Or Right Scalene? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsIsoscelesOrRightIsosceles(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     string
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(18, 80, 82),
			want:     "",
		},
		{
			name:     "case-2",
			triangle: NewTriangle(21.0, 22.0, 21.0),
			want:     "Isosceles",
		},
		{
			name:     "case-3",
			triangle: NewTriangle(5.0, 5.0, 7.071067812),
			want:     "Right Isosceles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.isIsoscelesOrRightIsosceles()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("is triangle Isosceles Or Right Isosceles? = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	tests := []struct {
		name     string
		triangle *Triangle
		want     string
		err      error
	}{
		{
			name:     "case-1",
			triangle: NewTriangle(18, 80, 82),
			want:     "Right Scalene",
			err:      nil,
		},
		{
			name:     "case-2",
			triangle: NewTriangle(-1.0, 1.0, 1.0),
			want:     "",
			err:      errors.New("given sides must be positive"),
		},
		{
			name:     "case-3",
			triangle: NewTriangle(2.0, 6.0, 200.0),
			want:     "",
			err:      errors.New(fmt.Sprintf("sides with length %f, %f, %f can't form a triangle", 2.0, 6.0, 200.0)),
		},
		{
			name:     "case-4",
			triangle: NewTriangle(21.0, 22.0, 21.0),
			want:     "Isosceles",
			err:      nil,
		},
		{
			name:     "case-5",
			triangle: NewTriangle(3.0, 4.0, 6.0),
			want:     "Scalene",
			err:      nil,
		},
		{
			name:     "case-6",
			triangle: NewTriangle(5.0, 5.0, 7.071067812),
			want:     "Right Isosceles",
			err:      nil,
		},
		{
			name:     "case-7",
			triangle: NewTriangle(7e-61, 6e-61, 5e-61),
			want:     "Right Scalene",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.triangle.GetType()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("error is = %v, want = %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("the type of triangle for the given input is = %v, want = %v", got, tt.want)
			}
		})
	}
}
