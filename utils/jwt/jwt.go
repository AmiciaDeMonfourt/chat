package jwt

func GenerateToken(userid int64) (string, error) {
	return "dummy-token", nil
}

func GetIdFromClaims(tokenString string) (int64, error) {
	return 1, nil
}
