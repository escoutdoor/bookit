package converter

import (
	"github.com/escoutdoor/bookit/internal/model"
	repomodel "github.com/escoutdoor/bookit/internal/repository/review/model"
)

func ToReviewFromRepository(rev *repomodel.Review) *model.Review {
	return &model.Review{
		ID:          rev.ID,
		Content:     rev.Content,
		Rating:      rev.Rating,
		RenterID:    rev.RenterID,
		ApartmentID: rev.ApartmentID,
		CreatedAt:   rev.CreatedAt,
	}
}

func ToReviewsFromRepository(reporevs []*repomodel.Review) []*model.Review {
	var revs []*model.Review
	for _, v := range reporevs {
		revs = append(revs, ToReviewFromRepository(v))
	}

	return revs
}
