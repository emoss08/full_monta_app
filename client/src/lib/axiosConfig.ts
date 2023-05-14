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

import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;
export const TOKEN_KEY = import.meta.env.VITE_TOKEN_KEY;

const getToken = (): string | null => {
  return localStorage.getItem(TOKEN_KEY);
};

axios.interceptors.request.use(
  req => {
    req.baseURL = API_URL;
    const token = getToken();
    if (token) {
      req.headers.Authorization = `Bearer ${token}`;
    }
    console.info(`Making request to ${req.baseURL}${req.url}`);
    return req
  },
  (error: any) => {
    return Promise.reject(error);
  }
)

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      localStorage.removeItem(TOKEN_KEY)
    }
    return Promise.reject(error);
  }
);

export default axios;
