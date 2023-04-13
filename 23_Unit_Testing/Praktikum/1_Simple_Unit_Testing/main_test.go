package main

import (
	"errors"
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success",
			args: args{
				a: 3,
				b: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiple(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success",
			args: args{
				a: 3,
				b: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Success",
			args: args{
				a: 6,
				b: 2,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "Divide by 0",
			args: args{
				a: 6,
				b: 0,
			},
			want:    0,
			wantErr: errors.New("cannot divice by zero"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if err != nil {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
