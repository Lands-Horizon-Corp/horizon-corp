import Notification from "@/components/notifications";
import { Role } from "@/types";


const sampleRoles: Role[] = [
  {
      id: "1",
      name: "Admin",
      description: "Administrator role with full permissions.",
      createdAt: new Date("2025-01-01T10:00:00Z"),
      updatedAt: new Date("2025-01-05T12:00:00Z"),
      deletedAt: undefined,
      permissions: [
          {
              id: "101",
              name: "Manage Users",
              description: "Allows managing user accounts.",
              read: true,
              readDescription: "Can view user accounts.",
              update: true,
              updateDescription: "Can update user accounts.",
              create: true,
              createDescription: "Can create user accounts.",
              roleId: "1",
              createdAt: new Date("2025-01-01T10:30:00Z"),
              updatedAt: new Date("2025-01-05T12:30:00Z"),
              deletedAt: undefined,
          },
          {
              id: "102",
              name: "Manage Settings",
              description: "Allows managing system settings.",
              read: true,
              readDescription: "Can view system settings.",
              update: true,
              updateDescription: "Can update system settings.",
              create: true,
              createDescription: "Can create new settings.",
              roleId: "1",
              createdAt: new Date("2025-01-02T09:00:00Z"),
              updatedAt: new Date("2025-01-06T11:00:00Z"),
              deletedAt: undefined,
          },
      ],
  },
  {
      id: "2",
      name: "Editor",
      description: "Editor role with limited permissions.",
      createdAt: new Date("2025-01-02T08:00:00Z"),
      updatedAt: new Date("2025-01-06T14:00:00Z"),
      deletedAt: undefined,
      permissions: [
          {
              id: "201",
              name: "Edit Articles",
              description: "Allows editing and creating articles.",
              read: true,
              readDescription: "Can view articles.",
              update: true,
              updateDescription: "Can edit articles.",
              create: true,
              createDescription: "Can create articles.",
              roleId: "2",
              createdAt: new Date("2025-01-03T09:00:00Z"),
              updatedAt: new Date("2025-01-06T15:00:00Z"),
              deletedAt: undefined,
          },
          {
              id: "202",
              name: "Publish Articles",
              description: "Allows publishing articles.",
              read: true,
              readDescription: "Can view publishing status.",
              update: true,
              updateDescription: "Can update publishing status.",
              create: false,
              createDescription: "Cannot create new publishing rules.",
              roleId: "2",
              createdAt: new Date("2025-01-04T10:00:00Z"),
              updatedAt: new Date("2025-01-06T16:00:00Z"),
              deletedAt: undefined,
          },
      ],
  },
  {
      id: "3",
      name: "Viewer",
      description: "Viewer role with read-only permissions.",
      createdAt: new Date("2025-01-03T07:00:00Z"),
      updatedAt: new Date("2025-01-06T13:00:00Z"),
      deletedAt: undefined,
      permissions: [
          {
              id: "301",
              name: "View Dashboard",
              description: "Allows viewing the dashboard.",
              read: true,
              readDescription: "Can view the dashboard.",
              update: false,
              updateDescription: "Cannot update the dashboard.",
              create: false,
              createDescription: "Cannot create new dashboard widgets.",
              roleId: "3",
              createdAt: new Date("2025-01-04T11:00:00Z"),
              updatedAt: new Date("2025-01-06T17:00:00Z"),
              deletedAt: undefined,
          },
      ],
  },
];

const AdminNotificationsPageTable = () => {
 
  return (
      <div className="flex w-full max-w-full pt-10 flex-col items-center px-4 pb-6 sm:px-8">
          <Notification userId={101} role={sampleRoles[0]}  />
      </div>
  )
}

export default AdminNotificationsPageTable
