import { apiClient } from "../../../configs";
import {
  getAllIngredientsFilters,
  TGetAllComponentsResults,
} from "../../../entities/component";

export async function fetchAllIngredients(
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
