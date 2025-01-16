import { MediaResource } from "../media";
import { CompanyResource } from "../company";
import { EmployeeResource, MemberResource } from "../profile";

import { PaginatedResult } from "../paginated-result";

export interface BranchResource {
  id: number;
  name: string;
  address?: string;
  longitude?: number;
  latitude?: number;
  email: string;
  contactNumber: string;
  isAdminVerified: boolean;
  media?: MediaResource;
  company?: CompanyResource;
  employees?: EmployeeResource[];
  members?: MemberResource[];
  createdAt: string;
  updatedAt: string;
}

export type BranchPaginatedResource = PaginatedResult<BranchResource>