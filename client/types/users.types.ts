// types/models.ts

export interface User {
  id: number;
  createdAt: string; // ISO 8601 formatted date-time string
  updatedAt: string; // ISO 8601 formatted date-time string
  deletedAt?: string; // Optional ISO 8601 formatted date-time string
  firstName: string;
  lastName: string;
  permanentAddress?: string;
  description?: string;
  birthdate: string; // ISO 8601 formatted date
  username: string;
  email: string;
}

export interface UserRegister {
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  birthdate: Date;
  permanentAddress: string;
  password: string;
  confirmPassword: string;
}

export interface UserLogin {
  username: string;
  password: string;
}
