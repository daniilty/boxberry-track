package response

import (
	"fmt"
	"io"
	"net/http"
)

// Validate - validate if response has right statusCode.
func Validate(resp *http.Response) error {
	if resp.StatusCode > http.StatusNoContent {
		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		return fmt.Errorf("request failed with status: %d, response body: %s", resp.StatusCode, string(bb))
	}

	return nil
}
