package services

type Compose interface {
	Config()
	Down()
	Logs()
	Pull()
	Restart()
	Stop()
	Up()
}

// func Config(composeFiles ...string) error {

// 	// Concatenate compose files with -f flag "-f file1.yml -f file2.yml -f file3.yml"
// 	args := ""
// 	for _, file := range composeFiles {
// 		args += fmt.Sprintf("-f %s ", file)
// 	}

// 	if err := utils.ExComposeCmd("config", args); err != nil {
// 		return err
// 	}

// 	return nil
// }
