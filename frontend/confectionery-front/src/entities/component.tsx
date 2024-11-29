import { TSupplier } from "./supplier";

export const CComponentCategoryNames = {
  ingredients: "ingredients",
  decorations: "decorations",
};

const CAllowComponentCategoryNames = [
  CComponentCategoryNames.ingredients,
  CComponentCategoryNames.decorations,
] as const;
export type UAllowsComponentCategoryNames =
  (typeof CAllowComponentCategoryNames)[number];

export type TComponentCategory = {
  id: number;
  name: UAllowsComponentCategoryNames;
};

export type TComponentType = {
  id: number;
  component_category_id: number;
  component_category: TComponentCategory;
  name: string;
};

export type TComponent = {
  id: number;
  type_id: number;
  componentType: TComponentType;
  name: string;
  article: string;
  unit: string;
  measure_unit: string;
};

export type TPurchasedComponent = {
  id: number;
  component_id: number;
  component: TComponent;
  supplier_id: number;
  supplier: TSupplier;
  quantity: number;
  purchase_price: number;
  shelf_life: string;
};

export type TGetAllComponentsResults = {
  total_rows: number;
  Data: TPurchasedComponent[];
  total_price: number;
  total_count: number;
};

export type getAllIngredientsFilters = {
  component_category: UAllowsComponentCategoryNames;
  name?: string;
  article?: string;
  shelf_life_from?: Date;
  shelf_life_to?: Date;
  sort?: "quantity" | "shelf_life";
  offset?: number;
  limit?: number;
};
