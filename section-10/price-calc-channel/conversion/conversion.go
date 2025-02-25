package conversion

import "strconv"

// StringsToFloats converts slice of strings to slice of floats
func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, s := range strings {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}
	return floats, nil
}
