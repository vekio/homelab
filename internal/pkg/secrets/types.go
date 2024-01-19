package secrets

type secretConfig struct {
	Authelia authelia `yaml:"authelia"`
	Gitea    gitea    `yaml:"gitea"`
	Immich   immich   `yaml:"immich"`
	Lldap    lldap    `yaml:"lldap"`
	Traefik  traefik  `yaml:"traefik"`
}

type authelia struct {
	JWTSecret                        string `yaml:"jwt_secret"`
	IdentityProviderOIDCHMACSecret   string `yaml:"identity_providers_oidc_hmac_secret"`
	IdentityProviderIssuerPrivateKey string `yaml:"identity_providers_issuer_private_key"`
	SessionSecret                    string `yaml:"session_secret"`
	StorageEncryptionKey             string `yaml:"storage_encryption_key"`
}

type gitea struct {
	OIDCSecret string
}

type immich struct {
	DBPass     string `yaml:"db_pass"`
	OIDCSecret string
}

type lldap struct {
	JWTSecret    string `yaml:"jwt_secret"`
	LDAPUserPass string `yaml:"ldap_user_pass"`
}

type traefik struct {
}
