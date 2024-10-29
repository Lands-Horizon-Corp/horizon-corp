import z from 'zod'

import { contactNumberSchema } from './common'
import {
  emailSchema,
  passwordSchema,
  lastNameSchema,
  userNameSchema,
  firstNameSchema,
  birthDateSchema,
  middleNameSchema,
  userAccountTypeSchema,
  permanentAddressSchema,
} from '@/validations/common'

export const signUpFormSchema = z
  .object({
    email: emailSchema,
    username: userNameSchema,
    firstName: firstNameSchema,
    middleName: middleNameSchema,
    lastName: lastNameSchema,
    birthDate: birthDateSchema,
    contactNumber: contactNumberSchema,
    permanentAddress: permanentAddressSchema,
    password: passwordSchema,
    confirmPassword: passwordSchema,
    acceptTerms: z
      .boolean()
      .default(false)
      .refine(
        (val) => {
          return val === true
        },
        {
          message: 'You must accept the terms and conditions',
        }
      ),
    accountType: userAccountTypeSchema,
  })
  .refine(({ password, confirmPassword }) => password === confirmPassword, {
    message: "Password doesn't match",
    path: ['confirm_password'],
  })
