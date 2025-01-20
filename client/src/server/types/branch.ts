import { IMediaResource } from "./media"; 
import { IMemberResource } from "./member";
import { ICompanyResource } from "./company"; 
import { IEmployeeResource } from "./employee"; 
import { IPaginatedResult } from "./paginated-result"; 

export interface IBranchRequest {
  name: string;
  address?: string;
  longitude?: number;
  latitude?: number;
  email: string;
  contactNumber: string;
  isAdminVerified: boolean;
  mediaId?: number;
  companyId: number;
}

export interface IBranchResource {
  id: number;
  name: string;
  address?: string;
  longitude?: number;
  latitude?: number;
  email: string;
  contactNumber: string;
  isAdminVerified: boolean;
  media?: IMediaResource;
  company?: ICompanyResource;
  employees?: IEmployeeResource[];
  members?: IMemberResource[];
  createdAt: string;
  updatedAt: string;
}

export type IBranchPaginatedResource = IPaginatedResult<IBranchResource>