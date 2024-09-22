import { BaseModel } from "./base";
import { Role } from "./role";

export interface Permission extends BaseModel {
  id: string;
  name: string;
  description: string;
  read: boolean;
  readDescription: string;
  update: boolean;
  updateDescription: string;
  create: boolean;
  createDescription: string;
  roleId: string;

  role?: Role;
}