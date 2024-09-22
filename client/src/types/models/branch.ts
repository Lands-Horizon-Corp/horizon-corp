import { BaseModel } from './base'
import { Company } from './company'
import { Media } from './media'

export interface Branch extends BaseModel {
    id: string
    name: string
    email: string
    address: string
    contactNumber: string
    approved: boolean
    description: string
    theme: string
    companyId: string
    company: Company
    mediaId?: string
    profilePicture?: Media
}
