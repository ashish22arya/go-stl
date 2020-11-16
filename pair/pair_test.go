package pair

import (
	"reflect"
	"testing"
)

func Test_pair_First(t *testing.T) {
	type fields struct {
		first  interface{}
		second interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Test case 1", fields: fields{first:2, second:5}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pair{
				first:  tt.fields.first,
				second: tt.fields.second,
			}
			if got := p.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pair.First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pair_Second(t *testing.T) {
	type fields struct {
		first  interface{}
		second interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Test case 1", fields: fields{first:2, second:5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pair{
				first:  tt.fields.first,
				second: tt.fields.second,
			}
			if got := p.Second(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pair.Second() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePair(t *testing.T) {
	type args struct {
		first  interface{}
		second interface{}
	}
	tests := []struct {
		name string
		args args
		want pair
	}{
		{name: "Test case 1", args: args{first:2, second:5}, want: pair{first:2, second:5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePair(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyPair(t *testing.T) {
	type args struct {
		p pair
	}
	tests := []struct {
		name string
		args args
		want pair
	}{
		{name: "Test case 1", args: args{p : pair{first:2, second:5}}, want: pair{first:2, second:5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyPair(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
