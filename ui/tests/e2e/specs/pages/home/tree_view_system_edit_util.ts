import { System } from '../../../../../src/views/system/service';
import { fillForm } from './tree_view_system_util';

const clickSystemOptions = (system: System) => {
  cy.getBySel(`${system.name}_system_options`).click();
};

const clickSystemEditOption = (system: System) => {
  cy.getBySel('system-edit-option').click();
};

const assertForm = (system: System) => {
  cy.getBySel('name-input').should('contain.value', system.name);
};

const clickUpdateButton = () => {
  cy.getBySel('update-system-btn').click();
};

export const editSystem = (system: System, editedSystem: System): System => {
  clickSystemOptions(system);
  clickSystemEditOption(system);
  assertForm(system);
  fillForm(editedSystem);
  clickUpdateButton();
  return editedSystem;
};
