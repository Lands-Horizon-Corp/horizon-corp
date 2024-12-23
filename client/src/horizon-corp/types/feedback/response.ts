import { FilterPages } from "../table"

export interface FeedbackResource {
  id: number
  email: string
  description: string
  feedbackType: string
  createdAt: string
  updatedAt: string
}


export type FeedbackPaginatedResource = FilterPages<FeedbackResource>
