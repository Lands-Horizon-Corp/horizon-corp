export interface IContactsResource {
    id: number
    first_name: string
    last_name: string
    email: string
    description?: string
    created_at: string
    updated_at: string
}

export interface IContactsRequest {
    firstName: string
    lastName: string
    email: string
    contactNumber: string
    description?: string
}