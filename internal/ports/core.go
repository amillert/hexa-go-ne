package ports

type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Subtration(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}
