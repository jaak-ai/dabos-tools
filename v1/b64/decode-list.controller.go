package b64

func DecodeList(textList ...string) ([][]byte, error) {
	var textListBytes [][]byte

	for _, videoString := range textList {
		videoByte, err := Decode(videoString)
		if err != nil {
			return nil, err
		}
		textListBytes = append(textListBytes, videoByte)
	}

	return textListBytes, nil
}
