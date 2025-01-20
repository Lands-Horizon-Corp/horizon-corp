import { IMediaResource } from "./media"

export interface ITimeInRequest {
    timeIn: Date
    mediaIn: IMediaResource
}

export interface ITimeOutRequest {
    timeOut: Date
    mediaOut: IMediaResource
}

export interface ITimesheetResource {
    id: number
    employeeId: number
    timeIn: Date
    timeOut?: Date
    mediaInId?: number
    mediaOutId?: number
    createdAt: Date
    updatedAt: Date
    mediaIn?: IMediaResource
    mediaOut?: IMediaResource
}
