package parse

import (
	"testing"
)

func TestTransformRawGitToClean(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test base case",
			args: args{
				url: "github.com:ali-wetrill/hello-react.git",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "test web url",
			args: args{
				url: "https://github.com/ali-wetrill/hello-react",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "test web url trailing slash",
			args: args{
				url: "https://github.com/ali-wetrill/hello-react/",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "test git username prefix",
			args: args{
				url: "git@github.com:ali-wetrill/hello-react.git",
			},
			want: "git@github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "http",
			args: args{
				url: "https://github.com/ali-wetrill/hello-react.git",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "test bitbucket web url",
			args: args{
				url: "https://bitbucket.org/classranked/classranked/",
			},
			want: "bitbucket.org:classranked/classranked.git",
		},
		{
			name: "test bitbucket ssh url",
			args: args{
				url: "git@bitbucket.org:classranked/classranked.git",
			},
			want: "git@bitbucket.org:classranked/classranked.git",
		},
		{
			name: "test bitbucket ssh url without git  username",
			args: args{
				url: "bitbucket.org:classranked/classranked.git",
			},
			want: "bitbucket.org:classranked/classranked.git",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformRawGitToClean(tt.args.url); got != tt.want {
				t.Errorf("TransformRawGitToClean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHTTPURLFromOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ssh url w/o git username",
			args: args{
				origin: "github.com:ali-wetrill/hello-react.git",
			},
			want: "http://github.com/ali-wetrill/hello-react",
		},
		{
			name: "ssh url w/ git username",
			args: args{
				origin: "git@github.com:ali-wetrill/hello-react.git",
			},
			want: "http://github.com/ali-wetrill/hello-react",
		},
		{
			name: "http url",
			args: args{
				origin: "http://github.com/ali-wetrill/hello-react",
			},
			want: "http://github.com/ali-wetrill/hello-react",
		},
		{
			name: "https url",
			args: args{
				origin: "https://github.com/ali-wetrill/hello-react",
			},
			want: "http://github.com/ali-wetrill/hello-react",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHTTPURLFromOrigin(tt.args.origin); got != tt.want {
				t.Errorf("GetHTTPURLFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHTTPSURLFromOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "ssh url w/o git username",
			args: args{
				origin: "github.com:ali-wetrill/hello-react.git",
			},
			want: "https://github.com/ali-wetrill/hello-react",
		},
		{
			name: "ssh url w/ git username",
			args: args{
				origin: "git@github.com:ali-wetrill/hello-react.git",
			},
			want: "https://github.com/ali-wetrill/hello-react",
		},
		{
			name: "http url",
			args: args{
				origin: "http://github.com/ali-wetrill/hello-react",
			},
			want: "https://github.com/ali-wetrill/hello-react",
		},
		{
			name: "https url",
			args: args{
				origin: "https://github.com/ali-wetrill/hello-react",
			},
			want: "https://github.com/ali-wetrill/hello-react",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHTTPSURLFromOrigin(tt.args.origin); got != tt.want {
				t.Errorf("GetHTTPSURLFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSSHURLFromOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "ssh url w/o git username",
			args: args{
				origin: "github.com:ali-wetrill/hello-react.git",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "ssh url w/ git username",
			args: args{
				origin: "git@github.com:ali-wetrill/hello-react.git",
			},
			want: "git@github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "http url",
			args: args{
				origin: "http://github.com/ali-wetrill/hello-react",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
		{
			name: "https url",
			args: args{
				origin: "https://github.com/ali-wetrill/hello-react",
			},
			want: "github.com:ali-wetrill/hello-react.git",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSSHURLFromOrigin(tt.args.origin); got != tt.want {
				t.Errorf("GetSSHURLFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepoNameFromOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "ssh url w/o git username",
			args: args{
				origin: "github.com:ali-wetrill/hello-react.git",
			},
			want: "hello-react",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRepoNameFromOrigin(tt.args.origin); got != tt.want {
				t.Errorf("GetRepoNameFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}
