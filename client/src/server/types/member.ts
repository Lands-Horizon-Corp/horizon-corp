import { IUserBase } from "./common";
import { IPaginatedResult } from "./paginated-result";

export interface IMemberRequest {

}

export interface IMemberResource extends IUserBase {}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>

