package services

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/cli/conf"
	"github.com/vekio/homelab/cli/secrets"
	"github.com/vekio/homelab/cli/utils"
)

func InitAuthelia() error {
	autheliaConf := conf.Config.DirPath() + "/" + AUTHELIA

	err := _fs.CreateDir(autheliaConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	err = initAutheliaConfig(autheliaConf)
	if err != nil {
		return err
	}

	err = initAutheliaSecrets(autheliaConf)
	if err != nil {
		return err
	}

	return nil
}

func initAutheliaConfig(autheliaConf string) error {
	err := _fs.CreateDir(autheliaConf+"/config", os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	giteaOIDCHash, err := utils.Bcrypt(secrets.Secrets.Gitea.OIDCSecret)
	if err != nil {
		return err
	}

	immichOIDCHash, err := utils.Bcrypt(secrets.Secrets.Immich.OIDCSecret)
	if err != nil {
		return err
	}

	data := map[string]string{
		"DOMAIN":             os.Getenv("DOMAIN"),
		"SLD":                os.Getenv("SLD"),
		"TLD":                os.Getenv("TLD"),
		"SMTP_USERNAME":      os.Getenv("SMTP_USERNAME"),
		"SMTP_HOST":          os.Getenv("SMTP_HOST"),
		"SMTP_PORT":          os.Getenv("SMTP_PORT"),
		"SMTP_FROM":          os.Getenv("SMTP_FROM"),
		"GITEA_OIDC_SECRET":  giteaOIDCHash,
		"IMMICH_OIDC_SECRET": immichOIDCHash,
	}

	src := repoConfig + "/" + AUTHELIA + "/config/configuration.yml"
	dst := autheliaConf + "/config/configuration.yml"
	if err := utils.ParseConfig(src, dst, data); err != nil {
		return err
	}

	return nil
}

func initAutheliaSecrets(autheliaConf string) error {
	secretsDir := autheliaConf + "/secrets"
	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_JWT_SECRET_FILE",
		secrets.Secrets.Authelia.JWTSecret); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_IDENTITY_PROVIDERS_OIDC_HMAC_SECRET_FILE",
		secrets.Secrets.Authelia.IdentityProviderOIDCHMACSecret); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_IDENTITY_PROVIDERS_OIDC_ISSUER_PRIVATE_KEY_FILE",
		secrets.Secrets.Authelia.IdentityProviderIssuerPrivateKey); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_SESSION_SECRET_FILE",
		secrets.Secrets.Authelia.SessionSecret); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_STORAGE_ENCRYPTION_KEY_FILE",
		secrets.Secrets.Authelia.StorageEncryptionKey); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD_FILE",
		secrets.Secrets.Lldap.LDAPUserPass); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/AUTHELIA_NOTIFIER_SMTP_PASSWORD_FILE",
		os.Getenv("SMTP_PASSWORD")); err != nil {
		return err
	}

	return nil
}
