// services/MemberService.ts
import qs from 'query-string';

import { downloadFile } from '@/horizon-corp/helpers';
import UseServer from '@/horizon-corp/request/server';

import {
    MediaRequest,
    MemberRequest,
    MemberResource,
    MemberPaginatedResource,
} from '@/horizon-corp/types';

/**
 * Service class to handle CRUD operations for members.
 */
export default class MemberService {
    private static readonly BASE_ENDPOINT = '/member';

    /**
     * Retrieves a member by its ID with optional preloads.
     *
     * @param {number} id - The ID of the member to retrieve.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<MemberResource>} - A promise that resolves to the member resource.
     */
    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<MemberResource> {
        const url = qs.stringifyUrl({
            url: `${MemberService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        });

        const response = await UseServer.get<MemberResource>(url, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        });

        return response.data;
    }

    /**
     * Deletes a member by its ID.
     *
     * @param {number} id - The ID of the member to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${MemberService.BASE_ENDPOINT}/${id}`;
        await UseServer.delete<void>(endpoint);
    }

    /**
     * Updates an existing member by its ID with optional preloads.
     *
     * @param {number} id - The ID of the member to update.
     * @param {MemberRequest} memberData - The updated data for the member.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<MemberResource>} - A promise that resolves to the updated member resource.
     */
    public static async update(
        id: number,
        memberData: MemberRequest,
        preloads?: string[]
    ): Promise<MemberResource> {
        const url = qs.stringifyUrl({
            url: `${MemberService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        });

        const response = await UseServer.put<MemberRequest, MemberResource>(
            url,
            memberData,
            {
                headers: {
                    Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
                },
            }
        );
        return response.data;
    }

    /**
     * Filters members based on provided options with optional pagination.
     *
     * @param {Object} [props] - The options for filtering members.
     * @param {string} [props.filters] - The filters to apply for exporting members.
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<MemberPaginatedResource>} - A promise that resolves to paginated member resources.
     */
    public static async getMembers(props?: {
        sort?: string;
        filters?: string;
        preloads?: string[];
        pagination?: { pageIndex: number; pageSize: number };
    }): Promise<MemberPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {};

        const url = qs.stringifyUrl(
            {
                url: `${MemberService.BASE_ENDPOINT}`,
                query: {
                    sort,
                    preloads,
                    filter: filters,
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true }
        );

        const response = await UseServer.get<MemberPaginatedResource>(url);
        return response.data;
    }

    /**
     * Exports all members.
     *
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAll(): Promise<void> {
        const url = `${MemberService.BASE_ENDPOINT}/export`;
        await downloadFile(url, 'all_members_export.csv');
    }

    /**
     * Exports filtered members.
     *
     * @param {string} [filters] - The filters to apply for exporting members.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : '';
        const url = `${MemberService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`;
        await downloadFile(url, 'filtered_members_export.csv');
    }

    /**
     * Exports selected members.
     *
     * @param {number[]} ids - The IDs of the selected members to export.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No member IDs provided for export.');
        }

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&');
        const url = `${MemberService.BASE_ENDPOINT}/export-selected?${query}`;
        await downloadFile(url, 'selected_members_export.csv');
    }

    /**
     * Deletes multiple members by their IDs.
     *
     * @param {number[]} ids - The IDs of the members to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async deleteMany(ids: number[]): Promise<void> {
        const endpoint = `${MemberService.BASE_ENDPOINT}/bulk-delete`;

        const payload = { ids };

        await UseServer.delete<void>(endpoint, payload);
    }

    /**
     * Uploads a profile picture for the member.
     *
     * @param {number} id - The ID of the member.
     * @param {MediaRequest} data - The request payload containing the media information.
     * @param {string[]} [preloads=["Media"]] - Optional array of relations to preload.
     * @returns {Promise<MemberResource>} - A promise that resolves to the updated member resource.
     */
    public static async profilePicture(
        id: number,
        data: MediaRequest,
        preloads: string[] = ['Media']
    ): Promise<MemberResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || '';
        const separator = preloadParams ? '?' : '';
        const endpoint = `${MemberService.BASE_ENDPOINT}/profile-picture/${id}${separator}${preloadParams}`;
        const response = await UseServer.post<MediaRequest, MemberResource>(
            endpoint,
            data
        );
        return response.data;
    }
}
