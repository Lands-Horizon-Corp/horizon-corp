import axios from 'axios'

export const handleAxiosError = (error: unknown): string => {
    if (axios.isAxiosError(error)) {
        if (!error.response) {
            return "Network error. Please check your connection."
        }
        const { response } = error

        switch (response.status) {
            case 404:
                return 'Sorry, the requested resource is not found'
            case 500:
                return 'Sorry, the server encountered an error'
            default:
                return (
                    response.data?.message || response.data?.error ||
                    'An error occurred. Please try again'
                )
        }
    } else {
        const errorMessage =
            (error as Error)?.message || 'An unknown error occurred.'
        return errorMessage
    }
}
