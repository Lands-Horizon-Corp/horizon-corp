import { BranchResource } from "../branch";
import { MediaResource } from "../media";
import { OwnerResource } from "../profile";
import { FilterPages } from "../table";

export interface CompanyResource {
  id: number;
  name: string;
  description?: string;
  address?: string;
  longitude?: number;
  latitude?: number;
  contactNumber: string;
  isAdminVerified: boolean;
  owner?: OwnerResource;
  media?: MediaResource;
  branches?: BranchResource[];
  createdAt: string;
  updatedAt: string;
}

export type CompanyPaginatedResource = FilterPages<CompanyResource>
