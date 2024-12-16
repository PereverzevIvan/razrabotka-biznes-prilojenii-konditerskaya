import { useState } from "react";
import { Button } from "../../components/Button";
import { roles } from "../../configs";
import { useAuthContext, useToastContext } from "../../contexts";
import { TPurchasedComponent } from "../../entities";
import { PurchasedComponentEditModal } from "../PurchasedComponentEditModal/PurchasedComponentEditModal";
import "./IngredientsTable.scss";

export function IngredientsTable({
  componentsDatas,
}: {
  componentsDatas: TPurchasedComponent[] | undefined;
}) {
  if (componentsDatas && componentsDatas.length > 0) {
    return renderTable(componentsDatas);
  } else {
    return <p>Нет данных</p>;
  }
}

function renderTable(componentsDatas: TPurchasedComponent[]) {
  const { role } = useAuthContext();
  const [showEditModal, setShowEditModal] = useState(false);

  return (
    <>
      {/* <PurchasedComponentEditModal */}
      {/*   show={showEditModal} */}
      {/*   handleClose={() => setShowEditModal(false)} */}
      {/*   purComponent={editingPurComponent} */}
      {/* ></PurchasedComponentEditModal> */}
      <div className="table-container">
        <div className="table-container">
          <table className="table">
            <thead className="table__head">
              <tr className="table__row">
                <th className="table__header">Артикул</th>
                <th className="table__header">Наименование</th>
                <th className="table__header">Кол-во</th>
                <th className="table__header">Единица измерения</th>
                <th className="table__header">Закупочная стоимость</th>
                <th className="table__header">Основной поставщик</th>
                <th className="table__header">Срок годности</th>
                {(role == roles.director || role == roles.purchase_manager) && (
                  <th className="table__header">Действия</th>
                )}
              </tr>
            </thead>
            <tbody className="table__body">
              {renderValues(componentsDatas)}
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
}

function renderValues(componentsDatas: TPurchasedComponent[]) {
  const { role } = useAuthContext();
  const { addToast } = useToastContext();

  function handleEdit(component: TPurchasedComponent) {
    setEditingPurComponent(component);
  }

  function handleDelete(purComponent: TPurchasedComponent) {
    if (purComponent.quantity != 0) {
      addToast(
        `Нельзя удалить элемент ${purComponent.component.article}, так как его количество не равно 0`,
        "error"
      );
      return;
    }
    addToast(`Вы хотите удалить элемент ${purComponent.id}`, "error");
  }

  return componentsDatas?.map((purComponent: TPurchasedComponent) => (
    <tr className="table__row" key={purComponent.id}>
      <td className="table__data">{purComponent.component.article}</td>
      <td className="table__data">{purComponent.component.name}</td>
      <td className="table__data">{purComponent.quantity}</td>
      <td className="table__data">{purComponent.component.measure_unit}</td>
      <td className="table__data">{purComponent.purchase_price}₽</td>
      <td className="table__data">
        {purComponent.supplier ? purComponent.supplier?.name : "Нет данных"}
      </td>
      <td className="table__data">
        {purComponent.shelf_life ? purComponent.shelf_life : "Нет данных"}
      </td>
      {(role == roles.director || role == roles.purchase_manager) && (
        <td className="table__data">
          <Button onClick={() => handleEdit(purComponent)}>Изменить</Button>
          <Button onClick={() => handleDelete(purComponent)} color="red">
            Удалить
          </Button>
        </td>
      )}
    </tr>
  ));
}
