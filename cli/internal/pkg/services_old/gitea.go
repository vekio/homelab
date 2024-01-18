package services

// import (
// 	"os"

// 	_fs "github.com/vekio/fs"
// 	"github.com/vekio/homelab/cli/conf"
// )

// func InitGitea() error {
// 	giteaConf := conf.Config.DirPath() + "/" + GITEA

// 	err := _fs.CreateDir(giteaConf, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	dataDir := giteaConf + "/data"
// 	err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
