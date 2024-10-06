package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

// initAuthelia init function for authelia service
func (h Homelab) initAuthelia() (err error) {
	err = initAutheliaConfig(h.Config, h.Secrets)
	if err != nil {
		return err
	}

	err = initAutheliaSecrets(h.Config, h.Secrets)
	if err != nil {
		return err
	}

	return nil
}

// initAutheliaConfig creates authelia config file
func initAutheliaConfig(c config.HomelabConfig, s secrets.HomelabSecrets) (err error) {
	volumeDir := config.Manager.DirPath() + "/" + services.AUTHELIA

	err = _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	configDir := volumeDir + "/config"
	err = _fs.CreateDir(configDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	giteaOIDCHash, err := Bcrypt(s.Gitea.OIDCSecret)
	if err != nil {
		return err
	}

	immichOIDCHash, err := Bcrypt(s.Immich.OIDCSecret)
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

	src := c.Repository + "/" + services.AUTHELIA + "/config/configuration.yml"
	dst := volumeDir + "/configuration.yml"
	if err := ParseTemplate(src, dst, data); err != nil {
		return err
	}

	return nil
}

func initAutheliaSecrets(hc config.HomelabConfig, s secrets.HomelabSecrets) error {
	secretsDir := secrets.Manager.DirPath() + "/secrets/" + services.AUTHELIA

	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_JWT_SECRET_FILE",
		s.Authelia.JWTSecret); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_IDENTITY_PROVIDERS_OIDC_HMAC_SECRET_FILE",
		s.Authelia.IdentityProviderOIDCHMACSecret); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_IDENTITY_PROVIDERS_OIDC_ISSUER_PRIVATE_KEY_FILE",
		s.Authelia.IdentityProviderIssuerPrivateKey); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_SESSION_SECRET_FILE",
		s.Authelia.SessionSecret); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_STORAGE_ENCRYPTION_KEY_FILE",
		s.Authelia.StorageEncryptionKey); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD_FILE",
		s.Lldap.LDAPUserPass); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/AUTHELIA_NOTIFIER_SMTP_PASSWORD_FILE",
		os.Getenv("SMTP_PASSWORD")); err != nil {
		return err
	}

	return nil
}
