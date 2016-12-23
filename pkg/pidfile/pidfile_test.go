package pidfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePID(t *testing.T) {
	dir, _ := ioutil.TempDir("", "agent_test")
	defer os.RemoveAll(dir)

	pidFilePath := filepath.Join(dir, "this_should_be_created", "agent.pid")
	err := WritePID(pidFilePath)
	assert.Nil(t, err)
	data, err := ioutil.ReadFile(pidFilePath)
	assert.Nil(t, err)
	pid, err := strconv.Atoi(string(data))
	assert.Nil(t, err)
	assert.Equal(t, pid, os.Getpid())
}

func TestIsProcess(t *testing.T) {
	assert.True(t, isProcess(os.Getpid()))
}