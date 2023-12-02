package helpers

import (
	"net/url"
	"strconv"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

func ParsePaginationParams(query url.Values) (domain.PostPaginationParams, error) {
	var params domain.PostPaginationParams

	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	if pageStr == "" && limitStr == "" {
		return params, nil
	}

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return params, err
		}
		params.Page = page
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return params, err
		}
		params.Limit = limit
	}

	return params, nil
}
