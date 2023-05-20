package pointer

type Result struct {
	Foo string `json:"foo"`
}

type PtResult struct {
	Foo *string `json:"foo"`
}
