export function IngredientsAccountingPage() {
  return (
    <section className="page">
      <h1 className="title">IngredientsAccountingPage</h1>
      <div className="info-block">
        <p className="common-text">Общее кол-во выведенных позиций: [100]</p>
        <p className="common-text">Общее кол-во ингредиентов: [100]</p>
        <p className="common-text">Общее кол-во украшений: [100]</p>
        <p className="common-text">Общая закупочная стоимость: [100]</p>
      </div>
      <select className="select">
        <option value="1">1</option>
        <option value="2">2</option>
        <option value="3">3</option>
      </select>
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
              </tr>
            </thead>
            <tbody className="table__body">
              <tr className="table__row">
                <td className="table__data">1</td>
                <td className="table__data">2</td>
                <td className="table__data">3</td>
                <td className="table__data">4</td>
                <td className="table__data">5</td>
                <td className="table__data">6</td>
                <td className="table__data">7</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>
  );
}
