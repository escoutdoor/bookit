package avatar

import (
	"context"
	"errors"
	"net/http"

	"github.com/escoutdoor/bookit/internal/controller/http/resp"
)

const formAvatarKey = "avatar"

func (c *controller) Upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024)

	f, hdr, err := r.FormFile(formAvatarKey)
	if err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			resp.Error(w, http.StatusBadRequest, "your avatar is too large to upload (over 2 mb)")
			return
		}

		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	defer f.Close()

	ctx := context.Background()
	url, err := c.uc.Upload(ctx, f, hdr.Size)
	if err != nil {
		c.log.Error("failed to upload avatar", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		URL string `json:"url"`
	}

	resp.JSON(w, http.StatusOK, response{
		URL: url,
	})
}
