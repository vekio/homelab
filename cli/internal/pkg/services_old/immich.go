package services

// import (
// 	"os"

// 	_fs "github.com/vekio/fs"
// 	"github.com/vekio/homelab/cli/conf"
// 	"github.com/vekio/homelab/cli/secrets"
// 	"github.com/vekio/homelab/cli/utils"
// )

// func InitImmich() error {
// 	immichConf := conf.Config.DirPath() + "/" + IMMICH

// 	err := _fs.CreateDir(immichConf, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	dataDir := immichConf + "/data"
// 	if err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
// 		return err
// 	}

// 	cacheDir := immichConf + "/cache"
// 	if err = _fs.CreateDir(cacheDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
// 		return err
// 	}

// 	err = initImmichSecrets(immichConf)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func initImmichSecrets(immichConf string) error {
// 	secretsDir := immichConf + "/secrets"
// 	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	if err = utils.WriteSecret(secretsDir+"/IMMICH_DB_PASS",
// 		secrets.Secrets.Immich.DBPass); err != nil {
// 		return err
// 	}

// 	return nil
// }
