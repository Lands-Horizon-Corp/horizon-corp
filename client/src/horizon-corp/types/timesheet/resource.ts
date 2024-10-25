import { MediaResource } from "../media";

export interface TimesheetResource {
  id: number;
  employeeId: number;
  timeIn: Date;
  timeOut?: Date;
  mediaInId?: number;
  mediaOutId?: number;
  createdAt: Date;
  updatedAt: Date;
  mediaIn?: MediaResource;
  mediaOut?: MediaResource;
}