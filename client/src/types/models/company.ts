import { BaseModel } from './base'
import { Branch } from './branch'
import { Owner } from './owner'

export interface Company extends BaseModel {
    id: string
    name: string
    description: string
    ownerId: string
    owner: Owner
    branches: Branch[]
}
