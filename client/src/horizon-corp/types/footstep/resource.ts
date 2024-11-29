import { AdminResource, EmployeeResource, MemberResource, OwnerResource } from "../profile";

export interface FootstepResource {
  id: number;
  description?: string;
  activity: string;
  accountType: string;
  createdAt: string;
  updatedAt: string;

  admin?: AdminResource;
  employee?: EmployeeResource;
  owner?: OwnerResource;
  member?: MemberResource;
}