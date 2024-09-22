import { UserBase } from "./base";
import { Branch } from "./branch";
import { Role } from "./role";

export interface Member extends UserBase {
  branchId?: string;
  branch?: Branch;
}

export interface MemberRole {
  memberId: string;
  member: Member;
  roleId: string;
  role: Role;
  assignedAt: Date;
}