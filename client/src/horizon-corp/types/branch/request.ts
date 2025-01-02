export interface BranchRequest {
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