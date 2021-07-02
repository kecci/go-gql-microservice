package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = os.Setenv("APPENV", "development")
	trueCfg, _ := New("go-gql-microservice")

	failYamlRepoDir := filepath.Join(os.Getenv("GOPATH"), "src/github.com/kecci", "configYamlFail")
	failYamlDir := filepath.Join(failYamlRepoDir, "files/etc", "configYamlFail")
	failYamlPath := filepath.Join(failYamlDir, "configYamlFail.development.yaml")

	_, err := os.Stat(failYamlDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(failYamlDir, os.ModeDir|os.ModePerm)
		if err != nil {
			t.Fatalf("can't create directory for config test file, err %v", err)
			return
		}
	}
	defer func() {
		_ = os.RemoveAll(failYamlRepoDir)
	}()

	_, err = os.Stat(failYamlPath)
	if os.IsExist(err) {
		_ = os.Remove(failYamlPath)
	}
	failYamlFile, err := os.Create(failYamlPath)
	if err != nil {
		t.Fatalf("can't create config test file, err %v", err)
		return
	}
	_, err = failYamlFile.WriteString("fail fail yaml | asdlkfj\n  fail\n  yaml\nfasldkj")
	if err != nil {
		t.Fatalf("can't write to config test file, err %v", err)
		return
	}
	err = failYamlFile.Sync()
	if err != nil {
		t.Fatalf("can't save config test file, err %v", err)
		return
	}
	err = failYamlFile.Close()
	if err != nil {
		t.Fatalf("can't close config test file, err %v", err)
		return
	}
	defer func() { _ = os.Remove(failYamlPath) }()

	tests := []struct {
		name     string
		repoName string
		want     *Config
		wantErr  bool
		prep     func()
	}{
		{
			name:     "normal",
			repoName: "go-gql-microservice",
			want:     trueCfg,
			wantErr:  false,
		},
		{
			name:     "not found config",
			repoName: "not found",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "yaml invalid",
			repoName: "configYamlFail",
			want:     nil,
			wantErr:  true,
			prep: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prep != nil {
				tt.prep()
			}
			got, err := New(tt.repoName)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
