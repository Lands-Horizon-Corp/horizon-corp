export interface RolesResource {
  id: number;
  name: string;
  description?: string;
  apiKey?: string;
  createdAt: Date;
  updatedAt: Date;

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