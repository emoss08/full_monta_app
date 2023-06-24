/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import { RouteObject } from "react-router-dom";
import { lazy } from "react";

const HomePage = lazy(() => import("../pages"));
const LoginPage = lazy(() => import("../pages/LoginPage"));
const ErrorPage = lazy(() => import("../pages/ErrorPage"));
const LogoutPage = lazy(() => import("../pages/LogoutPage"));
const UserManagementPage = lazy(
  () => import("../pages/admin/users/UserManagement")
);
const UserSettingsPage = lazy(() => import("../pages/users/UserSettings"));
const ResetPasswordPage = lazy(() => import("../pages/ResetPasswordPage"));

export type RouteObjectWithPermission = RouteObject & {
  title: string;
  group: string;
  path: string;
  excludeFromMenu?: boolean;
  permission?: string;
};

export const routes: RouteObjectWithPermission[] = [
  {
    title: "Home",
    group: "main",
    path: "/",
    element: <HomePage />,
  },
  // Authentication Pages
  {
    title: "Login",
    group: "auth",
    path: "/login",
    element: <LoginPage />,
    excludeFromMenu: true,
  },
  {
    title: "Logout",
    group: "auth",
    path: "/logout",
    element: <LogoutPage />,
    excludeFromMenu: true,
  },
  {
    title: "Reset Password",
    group: "auth",
    path: "/reset-password",
    element: <ResetPasswordPage />,
    excludeFromMenu: true,
  },
  // User Management
  {
    title: "User Management",
    group: "admin",
    path: "/admin/users",
    element: <UserManagementPage />,
    permission: "admin.users.view",
  },
  {
    title: "User Settings",
    group: "user",
    path: "/account/settings/:userId",
    element: <UserSettingsPage />,
    excludeFromMenu: true,
  },
  {
    title: "Error",
    group: "error",
    path: "*",
    element: <ErrorPage />,
    excludeFromMenu: true,
  },
];
