import { MediaResource } from "../media";
import { CompanyResource } from "../company";
import { EmployeeResource, MemberResource } from "../profile";

import { FilterPages } from "../table";

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

export type BranchPaginatedResource = FilterPages<BranchResource>