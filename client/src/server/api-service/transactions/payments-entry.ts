import APIService from '../api-service'
import {
    IPaymentTransactionRequest,
    IPaymentTransactionResource,
} from '@/server/types/transactions/payment-transaction'

export default class PaymentService {
    private static readonly BASE_ENDPOINT = '/payments'

    public static async getAll(): Promise<IPaymentTransactionResource[]> {
        const response = await APIService.get<IPaymentTransactionResource[]>(
            PaymentService.BASE_ENDPOINT
        )
        return response.data
    }

    public static async create(
        paymentData: IPaymentTransactionRequest
    ): Promise<IPaymentTransactionResource> {
        const response = await APIService.post<
            IPaymentTransactionRequest,
            IPaymentTransactionResource
        >(PaymentService.BASE_ENDPOINT, paymentData)
        return response.data
    }
}
