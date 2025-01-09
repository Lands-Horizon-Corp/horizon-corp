import UseServer from '../../request/server'
import { TimesheetResource, TimeInRequest, TimeOutRequest } from '../../types'

/**
 * Service class to handle timesheet-specific operations.
 */
export default class TimesheetService {
    private static readonly BASE_ENDPOINT = '/timesheet'

    /**
     * Records a time-in for the employee.
     *
     * @param {TimeInRequest} timeInData - The data for the time-in event.
     * @returns {Promise<TimesheetResource>} - A promise that resolves to the recorded timesheet resource.
     */
    public static async timeIn(
        timeInData: TimeInRequest
    ): Promise<TimesheetResource> {
        const endpoint = `${TimesheetService.BASE_ENDPOINT}/time-in`
        const response = await UseServer.post<TimeInRequest, TimesheetResource>(
            endpoint,
            timeInData
        )
        return response.data
    }

    /**
     * Records a time-out for the employee.
     *
     * @param {TimeOutRequest} timeOutData - The data for the time-out event.
     * @returns {Promise<TimesheetResource>} - A promise that resolves to the updated timesheet resource with time-out recorded.
     */
    public static async timeOut(
        timeOutData: TimeOutRequest
    ): Promise<TimesheetResource> {
        const endpoint = `${TimesheetService.BASE_ENDPOINT}/time-out`
        const response = await UseServer.post<
            TimeOutRequest,
            TimesheetResource
        >(endpoint, timeOutData)
        return response.data
    }

    /**
     * Retrieves the current timesheet entry for the logged-in employee.
     *
     * @returns {Promise<TimesheetResource | null>} - A promise that resolves to the current timesheet resource, or null if no active entry.
     */
    public static async getCurrentEmployeeTime(): Promise<TimesheetResource | null> {
        const endpoint = `${TimesheetService.BASE_ENDPOINT}/current`
        const response = await UseServer.get<TimesheetResource | null>(endpoint)
        return response.data
    }
}
