import { AxiosError } from 'axios'

import { TErrorMessageExtractor } from '.'
import { IErrorResponse } from '@/server/types'
import { axiosErrorMessageExtractor } from '@/server/helpers'

export const axiosErrExtractor: TErrorMessageExtractor = [
    AxiosError<IErrorResponse>,
    (err: Error) =>
        axiosErrorMessageExtractor(err as AxiosError<IErrorResponse>),
]
