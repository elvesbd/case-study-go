package ports

// "context"

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
