import { AxiosError } from 'axios'

export const axiosErrorMessageExtractor = (
    error: AxiosError<{ message?: string; error?: string }>
): string => {
    if (!error.response) {
        return 'Network error. Please check your connection.'
    }
    const { response } = error

    switch (response.status) {
        case 404:
            return 'Sorry, the requested resource is not found'
        case 500:
            return 'Sorry, the server encountered an error'
        default:
            return (
                response.data?.message ||
                response.data?.error ||
                'An error occurred. Please try again'
            )
    }
}
