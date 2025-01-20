import { IPaginatedResult } from './paginated-result'

export interface IFeedbackRequest {
    email: string
    description: string
    feedbackType: string
}

export interface IFeedbackResource {
    id: number
    email: string
    description: string
    feedbackType: string
    createdAt: string
    updatedAt: string
}

export type IFeedbackPaginatedResource = IPaginatedResult<IFeedbackResource>
