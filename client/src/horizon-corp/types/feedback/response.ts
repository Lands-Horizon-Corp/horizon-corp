import { PaginatedResult } from "../paginated-result"

export interface FeedbackResource {
  id: number
  email: string
  description: string
  feedbackType: string
  createdAt: string
  updatedAt: string
}


export type FeedbackPaginatedResource = PaginatedResult<FeedbackResource>
