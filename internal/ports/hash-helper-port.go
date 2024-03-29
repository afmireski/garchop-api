package ports

type HashHelperPort interface {
	GenerateHash(password string, salt int) (string, error)

	CompareHash(password string, hash string) bool
}