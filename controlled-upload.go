package controlledupload

import (
	"net/http"
	"time"
)

func ControlledUpload(data []byte, serveOverSeconds int, w http.ResponseWriter) error {
	var chunks [][]byte
	var bytesPerChunk = len(data) / serveOverSeconds
	var numChunks = len(data) / bytesPerChunk

	// split the file into chunks
	for i := range numChunks {
		// check if this is the last iter
		if i == numChunks-1 {
			chunks = append(chunks, data[i*bytesPerChunk:])
			break
		}

		chunks = append(chunks, data[i*bytesPerChunk:(i+1)*bytesPerChunk])
	}

	// serve the chunks over time
	for _, chunk := range chunks {
		_, err := w.Write(chunk)
		if err != nil {
			return err
		}

		// sleep for a second
		time.Sleep(time.Second)
	}

	return nil
}
