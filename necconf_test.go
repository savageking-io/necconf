package necconf

import (
	log "github.com/sirupsen/logrus"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestConfig_Init(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	type fields struct {
		configDirectory string
	}
	type args struct {
		configDirectory string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Empty Config Directory", fields{""}, args{""}, true},
		{"Passing", fields{""}, args{"."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				configDirectory: tt.fields.configDirectory,
			}
			if err := c.Init(tt.args.configDirectory); (err != nil) != tt.wantErr {
				t.Errorf("Config.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_ReadConfig(t *testing.T) {
	conf0 := struct {
		Data string `yaml:"data"`
	}{}

	mockFS := fstest.MapFS{
		"broken.yaml": {
			Data: []byte("```"),
		},
		"config.yaml": {
			Data: []byte("---\ndata: test"),
		},
	}

	type fields struct {
		configDirectory string
	}
	type args struct {
		fsys     fs.FS
		filename string
		conf     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Not Initialized", fields{""}, args{mockFS, "config.yaml", conf0}, true},
		{"No Config", fields{"./"}, args{mockFS, "config.yaml", nil}, true},
		{"No FSYS", fields{"./"}, args{nil, "config.yaml", conf0}, true},
		{"Empty filename", fields{"./"}, args{mockFS, "", conf0}, true},
		{"Wrong filename", fields{"./"}, args{mockFS, "wrong.yaml", conf0}, true},
		{"Broken YAML", fields{"./"}, args{mockFS, "broken.yaml", conf0}, true},
		{"Passing", fields{"./"}, args{mockFS, "config.yaml", conf0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				configDirectory: tt.fields.configDirectory,
			}
			if err := c.ReadConfig(tt.args.fsys, tt.args.filename, tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
