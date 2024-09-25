import { UserBase } from './base'
import { Role } from './role'

export interface Admin extends UserBase {}

export interface AdminRole {
    adminId: string
    admin: Admin
    roleId: string
    role: Role
    assignedAt: Date
}
