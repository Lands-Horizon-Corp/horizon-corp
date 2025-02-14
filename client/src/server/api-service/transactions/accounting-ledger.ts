// import { TEntityId } from "@/server/types";
import { IAccountingLedgerRequest, IAccountingLedgerPaginatedResource, IAccountingLedgerResource } from "@/server/types/accounts/accounting-ledger";
import qs from 'query-string'
import APIService from "../api-service";

export default class IAccountingLedgerService {
    private static readonly BASE_ENDPOINT = '/accounting-ledger';

    private static async makeRequest<T>(apiCall: () => Promise<{ data: T }>): Promise<T> {
        try {
            const response = await apiCall();
            return response.data;
        } catch (error) {
            console.error('API Request Failed:', error);
            throw error;
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
            filters?: string;
            preloads?: string[];
            pagination?: { pageIndex: number; pageSize: number };
            sort?: string;
        }
    ): string {
        return qs.stringify(
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
        );
    }

    public static async create(
        ledgerData: IAccountingLedgerRequest,
        preloads?: string[]
    ): Promise<IAccountingLedgerResource> {
        const url = this.buildUrl('', { preloads });
        return this.makeRequest(() =>
            APIService.post<IAccountingLedgerRequest, IAccountingLedgerResource>(url, ledgerData)
        );
    }

    public static async getLedgers({
        sort,
        filters,
        preloads,
        pagination,
    }: {
        sort?: string;
        filters?: string;
        preloads?: string[];
        pagination?: { pageIndex: number; pageSize: number };
    }) {
        const url = this.buildUrl('', { filters, preloads, pagination, sort });
        return this.makeRequest(() =>
            APIService.get<IAccountingLedgerPaginatedResource>(url)
        );
    }

    // public static async delete(id: TEntityId): Promise<void> {
    //     const url = this.buildUrl(`/${id}`, {});
    //     return this.makeRequest(() => APIService.delete(url));
    // }

    // public static async update(
    //     id: TEntityId,
    //     ledgerData: IAccountingLedgerRequest,
    //     preloads?: string[]
    // ): Promise<IAccountingLedgerRequest> {
    //     const url = this.buildUrl(`/${id}`, { preloads });
    //     return this.makeRequest(() =>
    //         APIService.put<IAccountingLedgerRequest, IAccountingLedgerRequest>(url, ledgerData)
    //     );
    // }


}
