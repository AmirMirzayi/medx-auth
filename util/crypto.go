package util

func Hash(password string) string {
	return password //TODO: implement bcrypt
}

func Verify(password string, hash string) bool {
	return Hash(password) == hash
}
