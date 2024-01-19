package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name: "LoadConfig",
			want: &Config{
				AppName: "monorepo",
				Env:     "development",
				Database: Database{
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "postgres",
					Name:     "monorepo",
				},
				Server: Server{
					Port:            "8080",
					ShutdownTimeout: 10,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file
			tempFile, err := ioutil.TempFile("", "test_config_*.yml")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write the test configuration to the temporary file
			configContent := []byte("app_name: monorepo\nenv: development\ndatabase:\n  host: localhost\n  port: 5432\n")
			if _, err := tempFile.Write(configContent); err != nil {
				t.Fatal(err)
			}

			// Close the file to save changes
			tempFile.Close()

			// Load the configuration using the temporary file
			got, err := LoadConfig(tempFile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadDoctorConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *DoctorConfig
		wantErr bool
	}{
		{
			name: "LoadConfig",
			want: &DoctorConfig{
				AppName: "monorepo",
				Database: Database{
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "postgres",
					Name:     "monorepo",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file
			tempFile, err := ioutil.TempFile("", "test_config_*.yml")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Write the test configuration to the temporary file
			configContent := []byte("app_name: monorepo\nenv: development\ndatabase:\n  host: localhost\n  port: 5432\n")
			if _, err := tempFile.Write(configContent); err != nil {
				t.Fatal(err)
			}

			// Close the file to save changes
			tempFile.Close()

			// Load the configuration using the temporary file
			got, err := LoadDoctorConfig(tempFile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDoctorConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadDoctorConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
