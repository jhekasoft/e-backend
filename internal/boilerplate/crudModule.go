package boilerplate

type CRUDModuleBoilerplate struct {
	CommonModuleBoilerplate
}

func (b *CRUDModuleBoilerplate) Create() error {
	tmplTypeName := "crud"
	return b.CommonCreate(tmplTypeName)
}
