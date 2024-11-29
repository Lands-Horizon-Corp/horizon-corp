import { AccountStatus } from "../auth";
import { BranchResource } from "../branch";
import { CompanyResource } from "../company";
import { FootstepResource } from "../footstep";
import { GenderResource } from "../gender";
import { MediaResource } from "../media";
import { RolesResource } from "../roles";
import { TimesheetResource } from "../timesheet";

export interface MemberResource {
  accountType: string;
  id: number;
  firstName: string;
  lastName: string;
  middleName: string;
  permanentAddress: string;
  description: string;
  birthDate: string;
  username: string;
  email: string;
  isEmailVerified: boolean;
  isContactVerified: boolean;
  status: AccountStatus;
  contactNumber: string;
  longitude?: number;
  latitude?: number;
  media?: MediaResource;
  branch?: BranchResource;
  role?: RolesResource;
  genderId?: number;
  gender?: GenderResource;
  footsteps?: FootstepResource[];
  createdAt: string;
  updatedAt: string;
}

export interface OwnerResource {
  accountType: string;
  id: number;
  firstName: string;
  lastName: string;
  middleName: string;
  permanentAddress: string;
  description: string;
  birthDate: string;
  username: string;
  email: string;
  contactNumber: string;
  isEmailVerified: boolean;
  isContactVerified: boolean;
  status: AccountStatus;
  media?: MediaResource;
  companies?: CompanyResource[];
  genderId?: number;
  gender?: GenderResource;
  footsteps?: FootstepResource[];
  createdAt: string;
  updatedAt: string;
}

export interface AdminResource {
  accountType: string;
  id: number;
  firstName: string;
  lastName: string;
  middleName?: string;
  permanentAddress?: string;
  description?: string;
  birthDate: string;
  username: string;
  email: string;
  contactNumber: string;
  isEmailVerified: boolean;
  isContactVerified: boolean;
  isSkipVerification: boolean;
  status: AccountStatus;
  media?: MediaResource;
  role?: RolesResource;
  genderId?: number;
  gender?: GenderResource;
  footsteps?: FootstepResource[];
  createdAt: string;
  updatedAt: string;
}

export interface EmployeeResource {
  accountType: string;
  id: number;
  firstName: string;
  lastName: string;
  middleName?: string;
  permanentAddress?: string;
  description?: string;
  birthDate: string;
  username: string;
  email: string;
  contactNumber: string;
  isEmailVerified: boolean;
  isContactVerified: boolean;
  isSkipVerification: boolean;
  status: AccountStatus;
  media?: MediaResource;
  branch?: BranchResource;
  longitude?: number;
  latitude?: number;
  timesheets?: TimesheetResource[];
  role?: RolesResource;
  genderId?: number;
  gender?: GenderResource;
  footsteps?: FootstepResource[];
  createdAt: string;
  updatedAt: string;
}
