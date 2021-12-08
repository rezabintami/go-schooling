package convertpointer

func ConvertPointerString(x *string) *string {
	if *x == "" {
		x = nil
	}
	return x
}