package services

// func InitTraefik(repoConfig, envConfig string) error {
// 	// Create acme.json (also directories)
// 	acmeFile := fmt.Sprintf("%s/%s/acme.json", envConfig, "certificates")
// 	if err := _file.Touch(acmeFile, os.FileMode(0600)); err != nil {
// 		return err
// 	}

// 	// Copy config folder
// 	traefikConfig := fmt.Sprintf("%s/config", envConfig)
// 	if err := _dir.Copy(repoConfig, traefikConfig); err != nil {
// 		return err
// 	}

// 	// Parse traefik.yml
// 	traefikYMLFile := fmt.Sprintf("%s/traefik.yml", traefikConfig)
// 	data := map[string]string{
// 		"DOMAIN":             os.Getenv("DOMAIN"),
// 		"TRAEFIK_CERT_EMAIL": os.Getenv("TRAEFIK_CERT_EMAIL"),
// 	}

// 	if err := parseConfigFile(traefikYMLFile, data); err != nil {
// 		return err
// 	}

// 	return nil
// }
