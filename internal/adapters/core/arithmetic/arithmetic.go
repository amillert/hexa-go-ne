package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter { return &Adapter{} }

func (_ *Adapter) Addition(a int32, b int32) (int32, error)       { return a + b, nil }
func (_ *Adapter) Subtraction(a int32, b int32) (int32, error)    { return a - b, nil }
func (_ *Adapter) Multiplication(a int32, b int32) (int32, error) { return a * b, nil }
func (_ *Adapter) Division(a int32, b int32) (int32, error)       { return a / b, nil }
