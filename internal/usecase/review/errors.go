package review

import "errors"

var (
	ErrNotFound                 = errors.New("review not found")
	ErrNotAllowedToCreateReview = errors.New("you are not allowed to create reviews for this apartment")
)
