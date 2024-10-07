export interface RolesRequest {
  name: string;
  description?: string;
  apiKey: string;
  readRole?: boolean;
  writeRole?: boolean;
  updateRole?: boolean;
  deleteRole?: boolean;
  readErrorDetails?: boolean;
  writeErrorDetails?: boolean;
  updateErrorDetails?: boolean;
  deleteErrorDetails?: boolean;
  readGender?: boolean;
  writeGender?: boolean;
  updateGender?: boolean;
  deleteGender?: boolean;
}