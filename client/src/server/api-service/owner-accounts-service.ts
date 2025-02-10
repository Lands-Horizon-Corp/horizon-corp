import {
    IAccountsPaginatedResource,
    IAccountsRequest,
    IAccountsResource,
} from '../types/accounts/accounts'
import qs from 'query-string'
import APIService from './api-service'
import { TEntityId } from '../types'

/**
 * Service class to handle CRUD operations for Accounts.
 */
export default class AccountsService {
    private static readonly BASE_ENDPOINT = '/account'

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
                    filters, // Spread filters to ensure dynamic query keys
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true, skipEmptyString: true } // Avoid sending empty strings or nulls
        )
    }
    /**
     * Creates a new account.
     */
    public static async create(
        accountsData: IAccountsRequest,
        preloads?: string[]
    ): Promise<IAccountsResource> {
        const url = this.buildUrl('', { preloads })
        return this.makeRequest(() =>
            APIService.post<IAccountsRequest, IAccountsResource>(
                url,
                accountsData
            )
        )
    }

    /**
     * Deletes an account by ID.
     */
    public static async delete(id: TEntityId): Promise<void> {
        const url = this.buildUrl(`/${id}`, {})
        return this.makeRequest(() => APIService.delete(url))
    }

    /**
     * Updates an account by ID.
     */
    public static async update(
        id: TEntityId,
        accountData: IAccountsRequest,
        preloads?: string[]
    ): Promise<IAccountsResource> {
        const url = this.buildUrl(`/${id}`, { preloads })
        return this.makeRequest(() =>
            APIService.put<IAccountsRequest, IAccountsResource>(
                url,
                accountData
            )
        )
    }

    public static async getAccounts({
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
            APIService.get<IAccountsPaginatedResource>(url)
        )
    }

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
}
