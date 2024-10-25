import { AccountType, MediaResource } from '..'

export type AccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

export interface UserData {
  id: number
  username: string
  firstName: string
  middleName?: string
  lastName: string
  permanentAddress: string
  description: string
  birthDate: Date
  createdAt: Date
  email: string
  contactNumber: string
  accountType: AccountType
  status: AccountStatus
  isEmailVerified: boolean
  isContactVerified: boolean
  isSkipVerification: boolean
  media?: MediaResource
}
