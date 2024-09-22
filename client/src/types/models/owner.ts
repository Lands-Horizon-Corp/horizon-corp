import { UserBase } from "./base";
import { Role } from "./role";

export interface Owner extends UserBase { }

export interface OwnerRole {
  ownerId: string;
  owner: Owner;
  roleId: string;
  role: Role;
  assignedAt: Date;
}