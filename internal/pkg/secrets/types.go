package secrets

type s struct {
	Authelia Authelia `yaml:"authelia"`
	Gitea    Gitea    `yaml:"gitea"`
	Immich   Immich   `yaml:"immich"`
	Lldap    Lldap    `yaml:"lldap"`
	Traefik  Traefik  `yaml:"traefik"`
}

type Authelia struct {
	JWTSecret                        string `yaml:"jwt_secret"`
	IdentityProviderOIDCHMACSecret   string `yaml:"identity_providers_oidc_hmac_secret"`
	IdentityProviderIssuerPrivateKey string `yaml:"identity_providers_issuer_private_key"`
	SessionSecret                    string `yaml:"session_secret"`
	StorageEncryptionKey             string `yaml:"storage_encryption_key"`
}

type Gitea struct {
	OIDCSecret string
}

type Immich struct {
	DBPass     string `yaml:"db_pass"`
	OIDCSecret string
}

type Lldap struct {
	JWTSecret    string `yaml:"jwt_secret"`
	LDAPUserPass string `yaml:"ldap_user_pass"`
}

type Traefik struct {
}
