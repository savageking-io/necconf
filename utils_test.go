package necconf

import "testing"

func TestExtractDirectoryAndFilenameFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"Not YAML", args{"."}, "", "", true},
		{"Only Filename", args{"config.yml"}, ".", "config.yml", false},
		{"Relative path to directory", args{"./../"}, "", "", true},
		{"Absolute path to directory", args{"/some/directory"}, "", "", true},
		{"Full path", args{"/some/directory/config.yml"}, "/some/directory", "config.yml", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ExtractDirectoryAndFilenameFromPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractDirectoryAndFilenameFromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractDirectoryAndFilenameFromPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractDirectoryAndFilenameFromPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
