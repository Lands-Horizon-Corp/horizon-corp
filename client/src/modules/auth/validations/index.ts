import z from 'zod'

export const AuthModeSchema = z.enum(['Member', 'Owner', 'Admin', 'Employee'], {
    required_error: 'mode is required',
    message: 'Valid options are Member, Owner, Admin, Employee',
    invalid_type_error: 'Valid options are Member, Owner, Admin, Employee',
})
