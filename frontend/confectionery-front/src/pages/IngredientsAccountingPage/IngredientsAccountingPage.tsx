import { useQuery } from "@tanstack/react-query";
import { IngredientsTable } from "../../modules/IngredientsTable";
import { getAllIngredientsFilters, TPurchasedComponent } from "../../entities";
import { useState, useEffect } from "react";
import { useToastContext } from "../../contexts";
import { fetchAllIngredients } from "./api/ingredientsApi";
import "./IngredientsAccountingPage.scss";

export function IngredientsAccountingPage() {
  const { addToast } = useToastContext();

  const [filters, setFilters] = useState<getAllIngredientsFilters>({
    component_category: "ingredients",
  });

  const [componentsData, setComponentsData] = useState<TPurchasedComponent[]>(
    []
  );
  const {
    data: GetAllComponentsResult,
    isPending,
    isError,
    error,
  } = useQuery({
    queryKey: ["purchased-components", filters],
    queryFn: () => fetchAllIngredients(filters),
  });

  function handleChange(event: any) {
    const { name, value } = event.target;
    setFilters((prevData) => ({ ...prevData, [name]: value }));
  }

  useEffect(() => {
    if (isError) {
      addToast(error.message, "error");
    }
  }, [isError]);

  return (
    <section className="page">
      <h1 className="title">IngredientsAccountingPage</h1>
      <div className="info-block">
        <p className="common-text">
          Общее кол-во выведенных позиций: {GetAllComponentsResult?.total_rows}
        </p>
        <p className="common-text">
          Общее кол-во компонентов: {GetAllComponentsResult?.total_count}
        </p>
        <p className="common-text">
          Общая закупочная стоимость: {GetAllComponentsResult?.total_price}₽
        </p>
      </div>

      <div className="h-flex filters">
        <div className="h-flex filters__container">
          <label htmlFor="select-component_category">Вид компонента:</label>
          <select
            id="select-component_category"
            name="component_category"
            className="drop-down-list"
            onChange={handleChange}
          >
            <option value="ingredients" defaultChecked={true}>
              Ингредиенты
            </option>
            <option value="decorations">Декорации для тортов</option>
          </select>
        </div>

        <div className="h-flex filters__container">
          <label htmlFor="select-sort">Вид компонента:</label>
          <select
            id="select-sort"
            name="sort"
            className="drop-down-list"
            onChange={handleChange}
          >
            <option value="" defaultChecked={true}>
              без сортировки
            </option>
            <option value="quantity">По количеству</option>
            <option value="shelf_life">По сроку годности</option>
          </select>
        </div>
      </div>

      {isPending ? (
        <p>Loading...</p>
      ) : (
        <IngredientsTable componentsDatas={GetAllComponentsResult?.Data} />
      )}
    </section>
  );
}
