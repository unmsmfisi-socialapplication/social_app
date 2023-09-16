package test

import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"

type ProfileRepositoryTest struct {}

func NewProfileRepositoryTest() *ProfileRepositoryTest {
	return &ProfileRepositoryTest{}
}

func (pr *ProfileRepositoryTest) UpdateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return profile, nil
}
