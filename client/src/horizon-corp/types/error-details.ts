export interface ErrorDetails {
    id?: number
    error : string
    name: string
    stack?: string
    response?: unknown
    status?: number
}
