import { BaseModel } from './base'
import { Permission } from './Permission'

export interface Role extends BaseModel {
    id: string
    name: string
    description: string
    permissions: Permission[]
}
