import { System } from '../../../../../src/views/system/service';
import { fillForm } from './tree_view_namespace_util';
import {Namespace} from "../../../../../src/views/namespace/service";
import {clickNamespaceOptions} from "./tree_view_namespace_util";

const clickNamespaceEditOption = () => {
  cy.getBySel('namespace-edit-option').click();
};

const assertForm = (namespace: Namespace) => {
  cy.getBySel('name-input').should('contain.value', namespace.name);
};

const clickUpdateButton = () => {
  cy.getBySel('update-namespace-btn').click();
};

export const editNamespace = (system: System, namespace: Namespace, editedNamespace: Namespace): Namespace => {
  clickNamespaceOptions(system, namespace)
  clickNamespaceEditOption();
  assertForm(namespace);
  fillForm(editedNamespace);
  clickUpdateButton();
  return editedNamespace;
};
