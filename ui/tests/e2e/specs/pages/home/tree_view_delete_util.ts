import { System } from '../../../../../src/views/system/service';

const clickSystemOptions = (system: System) => {
  cy.getBySel(`${system.name}_system_options`).click();
};

const clickSystemDeleteOption = () => {
  cy.getBySel('system-delete-option').click();
};

const clickDeleteButton = () => {
  cy.getBySel('delete-system-btn').click();
};

export const deleteSystem = (system: System) => {
  clickSystemOptions(system);
  clickSystemDeleteOption();
  clickDeleteButton();
};
