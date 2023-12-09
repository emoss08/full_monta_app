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

import { EquipmentClassChoiceProps } from "@/lib/choices";
import { StatusChoiceProps } from "@/types/index";
import { BaseModel } from "./organization";

export interface EquipmentType extends BaseModel {
  id: string;
  status: StatusChoiceProps;
  name: string;
  description?: string | null;
  costPerMile?: string | null;
  equipmentClass: EquipmentClassChoiceProps;
  fixedCost?: string | null;
  variableCost?: string | null;
  height?: string | null;
  length?: string | null;
  width?: string | null;
  weight?: string | null;
  idlingFuelUsage?: string | null;
  exemptFromTolls: boolean;
}

export type EquipmentTypeFormValues = Omit<
  EquipmentType,
  "organization" | "businessUnit" | "created" | "modified" | "id"
>;

export interface EquipmentManufacturer extends BaseModel {
  id: string;
  status: StatusChoiceProps;
  name: string;
  description?: string | null;
}

export type EquipmentManufacturerFormValues = Pick<
  EquipmentManufacturer,
  "name" | "description" | "status"
>;

export interface Trailer extends BaseModel {
  id: string;
  code: string;
  isActive: boolean;
  planningComment?: string;
  equipmentType: string | null;
  make?: string;
  model?: string;
  year: number;
  vinNumber?: string;
  fleetCode?: string | null;
  tagIdentifier?: string;
  state?: string;
  licensePlateNumber?: string;
  licensePlateState?: string;
  licensePlateExpirationDate?: string | null;
  lastInspection?: string;
  length?: string | null;
  width?: string | null;
  height?: string | null;
  axles: number;
  owner?: string;
  isLeased: boolean;
  leasedDate?: string | null;
  leaseExpirationDate?: string | null;
  timesUsed: number;
  equipTypeName: string;
}

export type TrailerFormValues = Omit<
  Trailer,
  | "id"
  | "timesUsed"
  | "equipTypeName"
  | "organization"
  | "businessUnit"
  | "created"
  | "modified"
>;
