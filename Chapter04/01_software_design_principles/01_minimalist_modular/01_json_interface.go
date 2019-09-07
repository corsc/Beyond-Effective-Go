package _1_minimalist_modular

type JSONObject interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
