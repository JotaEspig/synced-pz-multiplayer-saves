package config

import (
	"log"
	"os/exec"
	"syncedpz/pkg/utils"

	"github.com/go-git/go-git/v5/plumbing/transport"
)

var (
	PZ_DataPath string
	PZ_ExePath  string
	PZ_SteamID  string
	GitAuth     transport.AuthMethod
	ServersPath = DataPath + "/servers"
)

func checkExistenceOfGit() {
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("Git is not installed. Install it")
	}
}

func init() {
	utils.EnsureDir(DataPath)
	utils.EnsureDir(ServersPath)
	checkExistenceOfGit()
}
