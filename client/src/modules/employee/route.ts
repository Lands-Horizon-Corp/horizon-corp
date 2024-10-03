import { createRoute } from '@tanstack/react-router'

import EmployeeLandingPage from './pages'
import EmployeeDashboardPage from './pages/dashboard'
import EmployeeMembersPage from './pages/users/members'
import EmployeeBranchPage from './pages/branch'
import EmployeeReportsPage from './pages/reports'
import EmployeeNotificationsPage from './pages/notifications'
import EmployeeFootstepsPage from './pages/footsteps'
import EmployeeProfilePage from './pages/profile'
import EmployeeSettingsPage from './pages/settings'

import EmployeeLayout from '@/modules/employee/layout'
import { rootRoute } from '@/root-route'

const employeeRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'employee',
    component: EmployeeLayout,
})

const employeeLandingRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: '/',
    component: EmployeeLandingPage,
})

const employeeDashboardRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'dashboard',
    component: EmployeeDashboardPage,
})

const employeeMembersRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: '/users/members',
    component: EmployeeMembersPage,
})

const employeeBranchRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'branch',
    component: EmployeeBranchPage,
})

const employeeReportsRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'reports',
    component: EmployeeReportsPage,
})

const employeeNotificationsRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'notifications',
    component: EmployeeNotificationsPage,
})

const employeeFootstepsRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'footsteps',
    component: EmployeeFootstepsPage,
})

const employeeProfileRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'profile',
    component: EmployeeProfilePage,
})

const employeeSettingsRoute = createRoute({
    getParentRoute: () => employeeRoute,
    path: 'settings',
    component: EmployeeSettingsPage,
})

const EmployeeRoute = employeeRoute.addChildren([
    employeeLandingRoute,
    employeeDashboardRoute,
    employeeMembersRoute,
    employeeBranchRoute,
    employeeReportsRoute,
    employeeNotificationsRoute,
    employeeFootstepsRoute,
    employeeProfileRoute,
    employeeSettingsRoute,
])

export default EmployeeRoute
