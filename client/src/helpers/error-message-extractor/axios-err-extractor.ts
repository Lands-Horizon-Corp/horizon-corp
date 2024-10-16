import { AxiosError } from "axios";

import { TErrorMessageExtractor } from ".";
import { ErrorResponse } from "@/horizon-corp/types";
import { axiosErrorMessageExtractor } from "@/horizon-corp/helpers";

export const axiosErrExtractor : TErrorMessageExtractor = [
    AxiosError<ErrorResponse>,
    (err: Error) =>
        axiosErrorMessageExtractor(err as AxiosError<ErrorResponse>),
]
