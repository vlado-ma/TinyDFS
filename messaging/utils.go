package messaging

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/vlado-github/tinydfs/logging"
)

func decodeMessage(message *Message, dec *json.Decoder) error {
	err := dec.Decode(&message)
	if err != nil {
		logging.AddError("Error: Decoding message.", err.Error())
	}
	return err
}

// Converts Json to byte stream and sends over the newtwork connection
// that is set in Encoder
func encodeMessage(message *Message, enc *json.Encoder) error {
	err := enc.Encode(message)
	if err != nil {
		logging.AddError("Error: Encoding message.", err.Error())
	}
	return err
}

func getCurrentDirectory() string {
	var pathToDir = "C://tinydfs_storage//"
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logging.AddError("Path of working directory not found.", err.Error())
	} else {
		pathToDir = path
	}
	return pathToDir
}
