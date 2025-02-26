import {
    ITransactionTypePaginatedResource,
    ITransactionTypeRequest,
    ITransactionTypeResource,
} from '../../types/transactions/transaction-type'
import qs from 'query-string'
import APIService from '../api-service'
import { TEntityId } from '../../types'
import { downloadFile } from '@/server/helpers'

/**
 * Service class to handle CRUD operations for Transaction Types.
 */
export default class TransactionTypeService {
    private static readonly BASE_ENDPOINT = '/transaction-type'

    /**
     * Centralized request handling for better error management.
     */
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

    /**
     * Constructs the request URL with optional preloads.
     */
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

    /**
     * Creates a new transaction type.
     */
    public static async create(
        transactionTypeData: ITransactionTypeRequest,
        preloads?: string[]
    ): Promise<ITransactionTypeResource> {
        const url = this.buildUrl('', { preloads })
        return this.makeRequest(() =>
            APIService.post<ITransactionTypeRequest, ITransactionTypeResource>(
                url,
                transactionTypeData
            )
        )
    }

    /**
     * Deletes a transaction type by ID.
     */
    public static async delete(id: TEntityId): Promise<void> {
        const url = this.buildUrl(`/${id}`, {})
        return this.makeRequest(() => APIService.delete(url))
    }

    /**
     * Updates a transaction type by ID.
     */
    public static async update(
        id: TEntityId,
        transactionTypeData: ITransactionTypeRequest,
        preloads?: string[]
    ): Promise<ITransactionTypeResource> {
        const url = this.buildUrl(`/${id}`, { preloads })
        return this.makeRequest(() =>
            APIService.put<ITransactionTypeRequest, ITransactionTypeResource>(
                url,
                transactionTypeData
            )
        )
    }

    /**
     * Fetches transaction types with optional filters, preloads, sorting, and pagination.
     */
    public static async getTransactionTypes({
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
        const url = this.buildUrl(``, { filters, preloads, pagination, sort })

        return this.makeRequest(() =>
            APIService.get<ITransactionTypePaginatedResource>(url)
        )
    }

    /**
     * Deletes multiple transaction types by IDs.
     */
    public static async deleteMany(ids: TEntityId[]): Promise<void> {
        const endpoint = `${TransactionTypeService.BASE_ENDPOINT}/bulk-delete`
        const payload = { ids }

        await APIService.delete<void>(endpoint, payload)
    }

    /**
     * Export all transaction types.
     */
    public static async exportAll(): Promise<void> {
        const url = this.buildUrl(`/export`, {})
        await downloadFile(url, 'all_transaction_type_export.xlsx')
    }

    /**
     * Export filtered transaction types.
     */
    public static async exportAllFiltered(filters?: string): Promise<void> {
        const url = this.buildUrl(`/export-search?filter=${filters || ''}`, {})
        await downloadFile(url, 'filtered_transaction_type_export.xlsx')
    }

    /**
     * Export selected transaction types by IDs.
     */
    public static async exportSelected(ids: TEntityId[]): Promise<void> {
        const url = qs.stringifyUrl(
            {
                url: `${TransactionTypeService.BASE_ENDPOINT}/export-selected`,
                query: { ids },
            },
            { skipNull: true }
        )

        await downloadFile(url, 'selected_transaction_type_export.xlsx')
    }
    /**
     * Export the current page of transaction types.
     */
    public static async exportCurrentPage(page: number): Promise<void> {
        const url = this.buildUrl(`/export-current-page/${page}`, {})
        await downloadFile(
            url,
            `current_page_transaction_type_${page}_export.xlsx`
        )
    }
}
