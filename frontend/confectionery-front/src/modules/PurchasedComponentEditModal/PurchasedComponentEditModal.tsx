import { useForm } from "react-hook-form";
import { Modal } from "../../components/Modal";
import { useState } from "react";
import { TPurchasedComponent } from "../../entities";
import { useQuery } from "@tanstack/react-query";

type TModalProps = {
  show: boolean;
  handleClose: () => void;
  purComponent?: TPurchasedComponent;
  title?: string;
  sub?: boolean;
};

export function PurchasedComponentEditModal(props: TModalProps) {
  const form = useForm<TPurchasedComponent>({
    defaultValues: props.purComponent,
  });

  return (
    <Modal
      show={props.show}
      handleClose={props.handleClose}
      title={props.title ? props.title : "Редактировать компонент"}
    >
      <form onSubmit={form.handleSubmit((data) => console.log(data))}>
        <select {...form.register("component_id")}></select>
        <input
          {...form.register("component_id")}
          type="text"
          placeholder="Название компонента"
        />
        <input
          {...form.register("supplier_id")}
          type="text"
          placeholder="Название компонента"
        />
        <input
          {...form.register("purchase_price")}
          type="number"
          placeholder="Цена"
        />
        <input
          {...form.register("shelf_life")}
          type="number"
          placeholder="Количество"
        />
        <button type="submit">Сохранить</button>
      </form>
    </Modal>
  );
}
