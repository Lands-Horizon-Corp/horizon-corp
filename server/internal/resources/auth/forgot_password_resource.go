package resources

type ForgotPasswordResource struct {
	Message string `json:"message"`
}

func ToResourceForgotPassword(message string) ForgotPasswordResource {
	return ForgotPasswordResource{
		Message: message,
	}
}
