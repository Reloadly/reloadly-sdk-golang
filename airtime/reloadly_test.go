package airtime_test

import (
	"reflect"
	"strconv"
	"testing"

	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
)

func TestNewClient(t *testing.T) {
	type args struct {
		clientId     string
		clientSecret string
		sandbox      bool
		opts         []reloadly.ClientOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *reloadly.Client
		wantErr bool
	}{
		{
			name: "client",
			args: args{
				clientId: "Nfgh-1234",
				clientSecret: "xyz-090",
				sandbox: true,
			},
			want: nil,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reloadly.NewClient(tt.args.clientId, tt.args.clientSecret, tt.args.sandbox, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterByPage(t *testing.T) {
	cases := [] struct{
		page int
		ExpectedPage int
	}{
		{
			page: 7,
			ExpectedPage: 7,
		},
	}



	for _, c := range cases {
		res := reloadly.FilterByPage(c.page)
		o := &reloadly.FilterOptions{}
		res(o)

		if res != nil{
			if strconv.Itoa(c.ExpectedPage) != o.Page{
				t.Fatalf("Expected Filter Page to be %s but got %s", strconv.Itoa(c.page), o.Page)
			}

		}


	}
}

func TestFilterBySize(t *testing.T) {
	cases := [] struct{
		Size int
		ExpectedSize int
	}{
		{
			Size: 7,
			ExpectedSize: 7,
		},
	}



	for _, c := range cases {
		res := reloadly.FilterBySize(c.Size)
		o := &reloadly.FilterOptions{}
		res(o)


		if res != nil{
			if strconv.Itoa(c.ExpectedSize) != o.Size{
				t.Fatalf("Expected Filter Page to be %s but got %s", strconv.Itoa(c.Size), o.Size)
			}

		}


	}
}