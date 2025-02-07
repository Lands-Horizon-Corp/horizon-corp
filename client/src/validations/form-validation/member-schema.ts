import z from 'zod'
import {
    emailSchema,
    lastNameSchema,
    passwordSchema,
    userNameSchema,
    birthDateSchema,
    firstNameSchema,
    middleNameSchema,
    contactNumberSchema,
    permanentAddressSchema,
} from '../common'

export const createMemberSchema = z.object({
    email: emailSchema,
    username: userNameSchema,
    firstName: firstNameSchema,
    middleName: middleNameSchema.optional(),
    lastName: lastNameSchema,
    birthDate: birthDateSchema,
    contactNumber: contactNumberSchema,
    permanentAddress: permanentAddressSchema,
    password: passwordSchema,
    confirmPassword: passwordSchema,
})
