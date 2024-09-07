import { useState } from "#app";

interface AuthState {
  isAuthenticated: boolean;
}

// Initialize or use existing auth state
export const useAuth = () => {
  const auth = useState<AuthState>("auth", () => ({
    isAuthenticated: false,
  }));
  return auth;
};
