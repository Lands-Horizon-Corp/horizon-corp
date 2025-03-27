import { ICheckClearingDepositRequest } from '@/server/types/transactions/check-clearing-deposit'
import APIService from '../api-service'

export default class CheckClearingService {
    private static readonly BASE_ENDPOINT = '/check-clearing'

    public static async create(
        paymentData: ICheckClearingDepositRequest
    ): Promise<ICheckClearingDepositRequest> {
        const response = await APIService.post<ICheckClearingDepositRequest>(
            CheckClearingService.BASE_ENDPOINT,
            paymentData
        )

        return response.data as ICheckClearingDepositRequest
    }
}
