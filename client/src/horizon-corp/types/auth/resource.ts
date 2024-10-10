export interface ChangePasswordResource {
  message: string
  updatedAt: Date
}

export interface ForgotPasswordResource {
  message: string
}

export interface SignInResource {
  email: string
  token: string
  loggedInAt: Date
}

export interface SignOutResource {
  message: string
  loggedOutAt: string
}

export interface SignUpResource {
  email: string
  token: string
  loggedInAt: Date
}