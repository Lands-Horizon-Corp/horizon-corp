// services/MemberService.ts
import qs from 'query-string'

import UseServer from '@/horizon-corp/request/server'
import { MemberPaginatedResource } from '@/horizon-corp/types'

/**
 * Service class to handle operations for members.
 */
export default class MemberService {
    private static readonly BASE_ENDPOINT = '/member'

    /**
     * Retrieves members with optional filtering, sorting, pagination, and preloads.
     *
     * @param {Object} [props] - The options for filtering members.
     * @param {string} [props.sort] - The sort order (encoded in Base64 string).
     * @param {string} [props.filters] - The filters to apply for members (encoded in Base64 string).
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<MemberPaginatedResource>} - A promise that resolves to paginated member resources.
     */
    public static async getMembers(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<MemberPaginatedResource> {
        const { sort, filters, preloads, pagination } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${MemberService.BASE_ENDPOINT}`,
                query: {
                    sort,
                    filter: filters,
                    preloads,
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true }
        )

        const response = await UseServer.get<MemberPaginatedResource>(url)
        return response.data
    }
}
