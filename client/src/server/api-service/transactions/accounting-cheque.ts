import { TEntityId } from "@/server/types";
import APIService from "../api-service";
import qs from "query-string";
import { downloadFile } from "@/server/helpers";
import { Cheque } from "@/server/types/cheque";

export default class ChequeService {
    private static readonly BASE_ENDPOINT = '/cheque';

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
        );
    }

    public static async create(chequeData: Cheque, preloads?: string[]): Promise<Cheque> {
        const url = this.buildUrl('', { preloads });
        return this.makeRequest(() => APIService.post<Cheque, Cheque>(url, chequeData));
    }

    public static async delete(id: TEntityId): Promise<void> {
        const url = this.buildUrl(`/${id}`, {});
        return this.makeRequest(() => APIService.delete(url));
    }

    public static async update(id: TEntityId, chequeData: Cheque, preloads?: string[]): Promise<Cheque> {
        const url = this.buildUrl(`/${id}`, { preloads });
        return this.makeRequest(() => APIService.put<Cheque, Cheque>(url, chequeData));
    }

    public static async getCheques({
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
        return this.makeRequest(() => APIService.get<{ data: Cheque[] }>(url));
    }

    public static async deleteMany(ids: TEntityId[]): Promise<void> {
        const endpoint = `${ChequeService.BASE_ENDPOINT}/bulk-delete`;
        const payload = { ids };
        await APIService.delete<void>(endpoint, payload);
    }

    public static async exportAll(): Promise<void> {
        const url = this.buildUrl('/export', {});
        await downloadFile(url, 'all_cheques_export.xlsx');
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const url = this.buildUrl(`/export-search?filter=${filters || ''}`, {});
        await downloadFile(url, 'filtered_cheques_export.xlsx');
    }

    public static async exportSelected(ids: TEntityId[]): Promise<void> {
        const url = qs.stringifyUrl(
            {
                url: `${ChequeService.BASE_ENDPOINT}/export-selected`,
                query: { ids },
            },
            { skipNull: true }
        );
        await downloadFile(url, 'selected_cheques_export.xlsx');
    }

    public static async exportCurrentPage(page: number): Promise<void> {
        const url = this.buildUrl(`/export-current-page/${page}`, {});
        await downloadFile(url, `current_page_cheques_${page}_export.xlsx`);
    }
}
