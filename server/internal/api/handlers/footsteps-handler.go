package handlers

// This will be a middleware

type FootstepHandler struct {
}

func NewFootstepHandler() *FootstepHandler {
	return &FootstepHandler{}
}

func (f *FootstepHandler) Create(activity string) {

}
