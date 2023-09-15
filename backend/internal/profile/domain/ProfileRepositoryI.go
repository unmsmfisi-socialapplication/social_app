
package domain

type ProfileRepositoryI interface {
    UpdateProfile(profile *Profile) (*Profile, error)
}
