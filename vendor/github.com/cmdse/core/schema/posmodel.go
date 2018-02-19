package schema

type PositionalModel struct {
	Binding      Binding
	IsSemantic   bool
	IsOptionPart bool
	IsOptionFlag bool
	name         string
}

func (posModel PositionalModel) String() string {
	return posModel.name
}

func (posModel PositionalModel) Equal(comparedPosModel *PositionalModel) bool {
	return posModel.name == comparedPosModel.name
}

var (
	PosModOptImplicitAssignmentLeftSide = &PositionalModel{
		Binding:      BindRight,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModOptImplicitAssignmentLeftSide",
	}
	PosModOptImplicitAssignmentValue = &PositionalModel{
		Binding:      BindLeft,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: false,
		name:         "PosModOptImplicitAssignmentValue",
	}
	PosModStandaloneOptAssignment = &PositionalModel{
		Binding:      BindNone,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModStandaloneOptAssignment",
	}
	PosModOptSwitch = &PositionalModel{
		Binding:      BindNone,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModOptSwitch",
	}
	PosModCommandOperand = &PositionalModel{
		Binding:      BindNone,
		IsSemantic:   true,
		IsOptionPart: false,
		IsOptionFlag: false,
		name:         "PosModCommandOperand",
	}
	PosModUnset = &PositionalModel{
		Binding:      BindUnknown,
		IsSemantic:   false,
		IsOptionPart: false,
		IsOptionFlag: false,
		name:         "PosModUnset",
	}
)
