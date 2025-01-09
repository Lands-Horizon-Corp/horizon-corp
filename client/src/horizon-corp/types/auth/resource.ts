import { AccountType } from '..'
import { IUserBase } from '../common'

export interface UserData extends IUserBase {
    accountType: AccountType
}
