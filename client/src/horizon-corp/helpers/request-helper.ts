import axios from 'axios'
import { ErrorDetails } from '@/horizon-corp/types'

export const handleAxiosError = (error: unknown): string => {
    console.log(error)
    if (axios.isAxiosError(error)) {
        if (error.response) {
            return (error.response.data as ErrorDetails).message
        }
        if (error.message?.toLocaleLowerCase().includes('network error')) {
            return 'Network Error. Please try again'
        }
        return error.message ?? 'An error occurred without a message'
    } else {
        return 'An unexpected error occurred. Please try again.'
    }
}
