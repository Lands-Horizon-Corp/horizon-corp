package config

func GenerateOTP() string {
	return "123456"
}

func SendOTP(contactNumber, otp string) error {
	return nil
}

func VerifyOTP(contactNumber, otp string) bool {
	return true
}
