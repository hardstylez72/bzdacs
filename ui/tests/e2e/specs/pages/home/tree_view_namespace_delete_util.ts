import {System} from "../../../../../src/views/system/service";
import {Namespace} from "../../../../../src/views/namespace/service";
import {clickNamespaceOptions} from "./tree_view_namespace_util";

const clickSystemDeleteOption = () => {
  cy.getBySel('namespace-delete-option').click();
};

const clickDeleteButton = () => {
  cy.getBySel('delete-btn').click();
};

export const deleteNamespace = (system: System, namespace: Namespace) => {
  clickNamespaceOptions(system, namespace)
  clickSystemDeleteOption()
  clickDeleteButton()
};
