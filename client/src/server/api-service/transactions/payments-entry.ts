import APIService from '../api-service'
import {
    IPaymentsEntryRequest,
    IPaymentsEntryResource,
} from '@/server/types/transactions/payments-Entry'

export default class PaymentsEntryService {
    private static readonly BASE_ENDPOINT = '/payments'

    public static async getAll(): Promise<IPaymentsEntryResource[]> {
        const response = await APIService.get<IPaymentsEntryResource[]>(
            PaymentsEntryService.BASE_ENDPOINT
        )
        return response.data
    }

    public static async create(
        paymentData: IPaymentsEntryRequest
    ): Promise<IPaymentsEntryResource> {
        const response = await APIService.post<
            IPaymentsEntryRequest,
            IPaymentsEntryResource
        >(PaymentsEntryService.BASE_ENDPOINT, paymentData)
        return response.data
    }
}
