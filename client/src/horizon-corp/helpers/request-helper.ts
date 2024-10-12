import axios from 'axios'
import { ErrorDetails } from '@/horizon-corp/types'

export const handleAxiosError = (error: unknown): string => {
    if (axios.isAxiosError(error)) {
        if (!error.response) {
            if (error.message === 'Network Error') {
                return 'Network error: Please check your internet connection and try again.'
            } else {
                return `Error: ${error.message}`
            }
        }
        const data = error.response.data as ErrorDetails        
        return data.message ?? error.message
    } else {
        return 'An unexpected error occurred. Please try again.'
    }
}
