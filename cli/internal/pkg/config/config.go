package config

var Config C

func init() {

}

// TODO validar si Settings es correcto tras cargar el fichero
func (s C) isValid() (bool, error) {
	return false, nil
}

// TODO cargar las variables de entorno
func (s C) loadEnvVars() (bool, error) {
	return false, nil
}
