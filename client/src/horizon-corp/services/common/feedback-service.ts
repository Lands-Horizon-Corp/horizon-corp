import qs from 'query-string'

import {
    FeedbackResource,
    FeedbackRequest,
    FeedbackPaginatedResource,
} from '../../types'
import UseServer from '../../request/server'
import { downloadFile } from '@/horizon-corp/helpers'

/**
 * Service class to handle CRUD operations for feedbacks.
 */
export default class FeedbackService {
    private static readonly BASE_ENDPOINT = '/feedback'

    /**
     * Retrieves a feedback by its ID with optional preloads.
     *
     * @param {number} id - The ID of the feedback to retrieve.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<FeedbackResource>} - A promise that resolves to the feedback resource.
     */
    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<FeedbackResource> {
        const url = qs.stringifyUrl({
            url: `${FeedbackService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await UseServer.get<FeedbackResource>(url)
        return response.data
    }

    /**
     * Creates a new feedback.
     *
     * @param {FeedbackRequest} feedbackData - The data for the new feedback.
     * @returns {Promise<FeedbackResource>} - A promise that resolves to the created feedback resource.
     */
    public static async create(
        feedbackData: FeedbackRequest
    ): Promise<FeedbackResource> {
        const response = await UseServer.post<
            FeedbackRequest,
            FeedbackResource
        >(FeedbackService.BASE_ENDPOINT, feedbackData)
        return response.data
    }

    /**
     * Deletes a feedback by its ID.
     *
     * @param {number} id - The ID of the feedback to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${FeedbackService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    /**
     * Filters feedbacks based on provided filters with optional preloads.
     *
     * @param {Object} props - Options for filtering feedbacks.
     * @param {string} [props.filters] - Filters to apply for searching feedbacks.
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<FeedbackPaginatedResource>} - A promise that resolves to paginated feedback resources.
     */
    public static async getFeedbacks(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<FeedbackPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl({
            url: `${FeedbackService.BASE_ENDPOINT}`,
            query: {
                sort,
                preloads,
                filter: filters,
                pageIndex: pagination?.pageIndex,
                pageSize: pagination?.pageSize,
            },
        })

        const response = await UseServer.get<FeedbackPaginatedResource>(url)
        return response.data
    }

    /**
     * Exports all feedbacks.
     *
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAll(): Promise<void> {
        const url = `${FeedbackService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_feedbacks_export.csv')
    }

    /**
     * Exports filtered feedbacks.
     *
     * @param {string} [filters] - Filters to apply for exporting feedbacks.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAllFiltered(filters?: string): Promise<void> {
        const url = qs.stringifyUrl({
            url: `${FeedbackService.BASE_ENDPOINT}/export-search`,
            query: { filter: filters },
        })

        await downloadFile(url, 'filtered_feedbacks_export.csv')
    }

    /**
     * Exports selected feedbacks.
     *
     * @param {number[]} ids - The IDs of the selected feedbacks to export.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No feedback IDs provided for export.')
        }

        const url = qs.stringifyUrl({
            url: `${FeedbackService.BASE_ENDPOINT}/export-selected`,
            query: { ids },
        })

        await downloadFile(url, 'selected_feedbacks_export.csv')
    }

    /**
     * Deletes multiple feedbacks by their IDs.
     *
     * @param {number[]} ids - The IDs of the feedbacks to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async deleteMany(ids: number[]): Promise<void> {
        const url = `${FeedbackService.BASE_ENDPOINT}/bulk-delete`
        const payload = { ids }

        await UseServer.delete<void>(url, payload)
    }
}
