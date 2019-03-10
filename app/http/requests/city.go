package requests

// City Request
type City struct {
	ID   uint64 `form:"id"`
	Name string `form:"name"`
}
