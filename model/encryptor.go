package model

type EncryptorStringList struct {
	StringArray []string `json:"stringArray"`
}

type StringElement struct {
	Value string
	Index int
}
type HashedStringElement struct {
	Value [32]byte
	Index int
}
