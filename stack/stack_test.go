package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *stack
	}{
		{name: "Test case 1", want: &stack{nil, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_IsEmpty(t *testing.T) {
	type fields struct {
		top    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Test case 1", fields: fields{top: nil, length: 0}, want: true},
		{name: "Test case 2", fields: fields{top: &node{2, nil}, length: 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := st.IsEmpty(); got != tt.want {
				t.Errorf("stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Size(t *testing.T) {
	type fields struct {
		top    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Test case 1", fields: fields{top: nil, length: 0}, want:0},
		{name: "Test case 2", fields: fields{top: &node{2, nil}, length: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := st.Size(); got != tt.want {
				t.Errorf("stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Top(t *testing.T) {
	type fields struct {
		top    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Test case 1", fields: fields{top: nil, length: 0}, want: nil},
		{name: "Test case 2", fields: fields{top: &node{2, nil}, length: 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := st.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Push(t *testing.T) {
	type fields struct {
		top    *node
		length int
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Test case 1", fields: fields{top: nil, length: 0}, args: args{d: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			st.Push(tt.args.d)
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		top    *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Test case 1", fields: fields{top: nil, length: 0}},
		{name: "Test case 2", fields: fields{top: &node{2, nil}, length: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &stack{
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			st.Pop()
		})
	}
}
