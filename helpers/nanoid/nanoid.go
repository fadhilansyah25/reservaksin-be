package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateNanoId() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	return nanoId, nil
}
