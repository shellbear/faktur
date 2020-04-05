package data

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Image can be an URL that will be used to fetch the image or the base64 encoded string representation of it.
type Image string

func (i *Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(*i)
}

func (i *Image) UnmarshalJSON(data []byte) error {
	content, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// If Image is not an URL then we assume that it's base64 encoded and we convert it to html data tag.
	if !strings.HasPrefix(content, "http://") && !strings.HasPrefix(content, "https://") {
		*i = Image("data:;base64," + content)
	} else {
		*i = Image(content)
	}

	return nil
}
