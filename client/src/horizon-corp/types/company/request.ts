export interface CompanyRequest {
    name: string
    description?: string
    address?: string
    longitude?: number
    latitude?: number
    contactNumber: string
    ownerId?: number
    companyId?: number
    mediaId?: number
    isAdminVerified?: boolean
}
