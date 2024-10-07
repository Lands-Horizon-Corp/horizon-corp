export interface ContactsRequest {
  first_name: string; // Required, max 255 characters
  last_name: string;  // Required, max 255 characters
  email: string;      // Required, must be a valid email, max 255 characters
  description?: string; // Optional, max 3000 characters
}