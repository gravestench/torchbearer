package procedure

type Dependency = ManagesProcedures

type ManagesProcedures interface {
	Register(ProcedureGenerator) error
	Deregister(ProcedureGenerator) error
	Begin(name string) (Procedure, error)
	End(instance Procedure) error
}

type HasProcedures interface {
	Procedures() []ProcedureGenerator
}
