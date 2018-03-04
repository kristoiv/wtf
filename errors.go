package wtf

const (
	ErrItemNotFound      = Error("item not found")
	ErrItemTitleRequired = Error("items must have a title")
	ErrItemIDRequired    = Error("item id required")
)

type Error string

func (e Error) Error() string { return string(e) }
