package secrets

import (
	"fmt"

	_rsa "github.com/vekio/crypto/rsa"
	_sgen "github.com/vekio/rand/secretgen"
)

// generateAutheliaSecrets generates secrets required for Authelia.
func generateAutheliaSecrets() (autheliaSecrets, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to generate JWT secret: %w", err)
	}

	identityProviderOIDCHMACSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to generate OIDC HMAC secret: %w", err)
	}

	keyPair, err := _rsa.GenerateKeyPair(4096)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to generate RSA key pair: %w", err)
	}
	identityProviderIssuerPrivateKey, err := _rsa.PrivateKeyData(keyPair)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to extract private key data: %w", err)
	}

	sessionSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to generate session secret: %w", err)
	}

	storageEncryptionKey, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return autheliaSecrets{}, fmt.Errorf("autheliaSecrets: failed to generate storage encryption key: %w", err)
	}

	autheliaSecrets := autheliaSecrets{
		JWTSecret:                        jwtSecret,
		IdentityProviderOIDCHMACSecret:   identityProviderOIDCHMACSecret,
		IdentityProviderIssuerPrivateKey: identityProviderIssuerPrivateKey,
		SessionSecret:                    sessionSecret,
		StorageEncryptionKey:             storageEncryptionKey,
	}

	return autheliaSecrets, nil
}

// generateLldapSecrets generates secrets required for LLDAP.
func generateLldapSecrets() (lldapSecrets, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return lldapSecrets{}, fmt.Errorf("lldapSecrets: failed to generate JWT secret: %w", err)
	}

	ldapUserPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return lldapSecrets{}, fmt.Errorf("lldapSecrets: failed to generate LDAP user password: %w", err)
	}

	lldapSecrets := lldapSecrets{
		JWTSecret:    jwtSecret,
		LDAPUserPass: ldapUserPass,
	}

	return lldapSecrets, nil
}

// generateGiteaSecrets generates secrets required for Gitea.
func generateGiteaSecrets() (giteaSecrets, error) {
	dbPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return giteaSecrets{}, fmt.Errorf("giteaSecrets: failed to generate db pass: %w", err)
	}

	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return giteaSecrets{}, fmt.Errorf("giteaSecrets: failed to generate OIDC secret: %w", err)
	}

	giteaSecrets := giteaSecrets{
		DBPass:     dbPass,
		OIDCSecret: oidcSecret,
	}

	return giteaSecrets, nil
}

// generateImmichSecrets generates secrets required for Immich.
func generateImmichSecrets() (immichSecrets, error) {
	dbPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return immichSecrets{}, fmt.Errorf("immichSecrets: failed to generate db pass: %w", err)
	}

	oidcSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return immichSecrets{}, fmt.Errorf("immichSecrets: failed to generate OIDC secret: %w", err)
	}

	immichSecrets := immichSecrets{
		DBPass:     dbPass,
		OIDCSecret: oidcSecret,
	}

	return immichSecrets, nil
}
