import {
  getAllIngredientsFilters,
  TGetAllComponentsResults,
} from "../entities";
import { apiClient } from "./apiClient";

export async function fetchAllPurchasedComponents(
  filters: getAllIngredientsFilters
): Promise<TGetAllComponentsResults> {
  const { data } = await apiClient.get(
    `/${filters?.component_category}/purchased`,
    {
      params: filters,
    }
  );

  return data;
}
