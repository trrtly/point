package hashids

import "github.com/speps/go-hashids"

// Config contains the information needed to generate hashids
type Config struct {
	// Alphabet is the alphabet used to generate new ids
	Alphabet string `envconfig:"POINT_HD_ALPHABET"`

	// MinLength is the minimum length of a generated id
	MinLength int `envconfig:"POINT_HD_MINLENGTH"`

	// Salt is the secret used to make the generated id harder to guess
	Salt string `envconfig:"POINT_HD_SALT"`
}

var DefaultHd *HD

// HD defines a hashids
type HD struct {
	*hashids.HashID
}

// New generate a new hashids instance.
func New(cfg *Config) (*HD, error) {
	hd := hashids.NewData()
	if cfg.Alphabet != "" {
		hd.Alphabet = cfg.Alphabet
	}
	hd.MinLength = cfg.MinLength
	hd.Salt = cfg.Salt

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	return &HD{h}, nil
}
