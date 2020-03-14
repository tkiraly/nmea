package nmea

import "fmt"

func Parse(sentence string, v interface{}) error {
	switch t := v.(type) {
	case *GPGGA:
		return parsegpgga(sentence, t)
	case *GPRMC:
		return parsegprmc(sentence, t)
	default:
		return fmt.Errorf("unknown sentence type: %T", t)
	}
}

func Build(sentence interface{}) (string, error) {
	switch v := sentence.(type) {
	case GPGGA:
		return buildgpgga(v)
	case *GPGGA:
		return buildgpgga(*v)
	case GPRMC:
		return buildgprmc(v)
	case *GPRMC:
		return buildgprmc(*v)
	default:
		return "", fmt.Errorf("unknown sentence type: %T", sentence)
	}
}
