import { AccountType } from '../common'
import { MediaRequest } from '../media'

export interface OwnerRequest {
    accountType: AccountType
    username: string
    firstName: string
    lastName: string
    middleName?: string
    email: string
    password: string
    confirmPassword: string
    birthDate: Date
    contactNumber: string
    permanentAddress: string
    media?: MediaRequest

    emailTemplate?: string
    contactTemplate?: string
}
