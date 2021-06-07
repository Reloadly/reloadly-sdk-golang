package Err

import (
	"reflect"
	"testing"
)

func TestErrorResponse_Error(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Error Test",
			fields: fields{
				Message:   "Oooops! An Error occured",
			},
			want:   "Oooops! An Error occured",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				Message:   tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse_GetErrorCode(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Error Code",
			fields: fields{
				ErrorCode: "0",
			},
			want:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				TimeStamp: tt.fields.TimeStamp,
				Message:   tt.fields.Message,
				Path:      tt.fields.Path,
				ErrorCode: tt.fields.ErrorCode,
				InfoLink:  tt.fields.InfoLink,
				Details:   tt.fields.Details,
			}
			if got := e.GetErrorCode(); got != tt.want {
				t.Errorf("GetErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse_GetErrorDetails(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name:   "Error Details",
			fields: fields{
				Details:   nil,
			},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				Details:   tt.fields.Details,
			}
			if got := e.GetErrorDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrorDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse_GetErrorInfo(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Test Error Info",
			fields: fields{
				InfoLink:  "https://linktoerror.com",
			},
			want: "https://linktoerror.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				InfoLink:  tt.fields.InfoLink,
			}
			if got := e.GetErrorInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrorInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse_GetErrorPath(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Error Path",
			fields: fields{
				Path: "Error/path",
			},
			want: "Error/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				Path:      tt.fields.Path,
			}
			if got := e.GetErrorPath(); got != tt.want {
				t.Errorf("GetErrorPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse_GetErrorTimeStamp(t *testing.T) {
	type fields struct {
		TimeStamp string
		Message   string
		Path      string
		ErrorCode string
		InfoLink  interface{}
		Details   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Error TimeStamp",
			fields: fields{
				TimeStamp: "2009-11-10 23:00:00 +0000 UTC m=+0.000000000",
			},
			want: "2009-11-10 23:00:00 +0000 UTC m=+0.000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorResponse{
				TimeStamp: tt.fields.TimeStamp,
			}
			if got := e.GetErrorTimeStamp(); got != tt.want {
				t.Errorf("GetErrorTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
