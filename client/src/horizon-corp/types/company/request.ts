export interface CompanyRequest {
  id?: number;
  name: string;
  description?: string;
  address?: string;
  longitude?: number;
  latitude?: number;
  contactNumber: string;
  ownerId?: number;
  mediaId?: number;
  isAdminVerified?: boolean
}
