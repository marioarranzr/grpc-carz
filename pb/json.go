package pb

// we can redefine the way something is json marshalled
func (loc *Location) MarshalJSON() ([]byte, error) {
	return []byte(`"LOCATION"`), nil
}
