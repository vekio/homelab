package secrets

type secretsManager struct {
	name string
	dir  string
	file string
}

type HomelabSecrets struct {
	Authelia autheliaSecrets `yaml:"authelia"`
	Gitea    giteaSecrets    `yaml:"gitea"`
	Immich   immichSecrets   `yaml:"immich"`
	Lldap    lldapSecrets    `yaml:"lldap"`
	Traefik  traefikSecrets  `yaml:"traefik"`
}

type autheliaSecrets struct {
	JWTSecret                        string `yaml:"jwt_secret"`
	IdentityProviderOIDCHMACSecret   string `yaml:"identity_providers_oidc_hmac_secret"`
	IdentityProviderIssuerPrivateKey string `yaml:"identity_providers_issuer_private_key"`
	SessionSecret                    string `yaml:"session_secret"`
	StorageEncryptionKey             string `yaml:"storage_encryption_key"`
}

type giteaSecrets struct {
	OIDCSecret string
}

type immichSecrets struct {
	DBPass     string `yaml:"db_pass"`
	OIDCSecret string
}

type lldapSecrets struct {
	JWTSecret    string `yaml:"jwt_secret"`
	LDAPUserPass string `yaml:"ldap_user_pass"`
}

type traefikSecrets struct {
}
