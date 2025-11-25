package boilerplate

type SimpleModuleBoilerplate struct {
	CommonModuleBoilerplate
}

func (b *SimpleModuleBoilerplate) Create() error {
	// Create init file
	err := b.CreateInitFile()
	if err != nil {
		return err
	}

	// Create module directory
	err = b.CreateModuleDir()
	if err != nil {
		return err
	}

	// Create module file
	err = b.CreateModuleFile("simple")
	if err != nil {
		return err
	}

	return nil
}
