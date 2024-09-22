import { UserBase } from './base'
import { Role } from './role'

export interface Employee extends UserBase {}

export interface EmployeeRole {
    employeeId: string
    employee: Employee
    roleId: string
    role: Role
    assignedAt: Date
}
