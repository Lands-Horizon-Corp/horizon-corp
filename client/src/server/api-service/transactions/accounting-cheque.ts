import { TEntityId } from '@/server/types'
import APIService from '../api-service'
import qs from 'query-string'
import { Cheque } from '@/server/types/cheque'

export default class ChequeService {
    private static readonly BASE_ENDPOINT = '/cheque'

    private static async makeRequest<T>(
        apiCall: () => Promise<{ data: T }>
    ): Promise<T> {
        try {
            const response = await apiCall()
            return response.data
        } catch (error) {
            console.error('API Request Failed:', error)
            throw error
        }
    }

    private static buildUrl(
        endpoint: string,
        {
            filters,
            preloads,
            pagination,
            sort,
        }: {
            filters?: string
            preloads?: string[]
            pagination?: { pageIndex: number; pageSize: number }
            sort?: string
        }
    ): string {
        return qs.stringifyUrl(
            {
                url: `${this.BASE_ENDPOINT}${endpoint}`,
                query: {
                    sort,
                    preloads,
                    filters,
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true, skipEmptyString: true }
        )
    }

    public static async create(
        chequeData: Cheque,
        preloads?: string[]
    ): Promise<Cheque> {
        const url = this.buildUrl('', { preloads })
        return this.makeRequest(() =>
            APIService.post<Cheque, Cheque>(url, chequeData)
        )
    }

    public static async delete(id: TEntityId): Promise<void> {
        const url = this.buildUrl(`/${id}`, {})
        return this.makeRequest(() => APIService.delete(url))
    }

    public static async update(
        id: TEntityId,
        chequeData: Cheque,
        preloads?: string[]
    ): Promise<Cheque> {
        const url = this.buildUrl(`/${id}`, { preloads })
        return this.makeRequest(() =>
            APIService.put<Cheque, Cheque>(url, chequeData)
        )
    }

    public static async getCheques({
        sort,
        filters,
        preloads,
        pagination,
    }: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }) {
        const url = this.buildUrl('', { filters, preloads, pagination, sort })
        return this.makeRequest(() => APIService.get<{ data: Cheque[] }>(url))
    }

    public static async deleteMany(ids: TEntityId[]): Promise<void> {
        const endpoint = `${ChequeService.BASE_ENDPOINT}/bulk-delete`
        const payload = { ids }
        await APIService.delete<void>(endpoint, payload)
    }
}
