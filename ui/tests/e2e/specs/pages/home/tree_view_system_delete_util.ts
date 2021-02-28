import { System } from '../../../../../src/system/entity';

const clickSystemOptions = (system: System) => {
  cy.getBySel(`${system.name}_system_options`).click();
};

const clickSystemDeleteOption = () => {
  cy.getBySel('system-delete-option').click();
};

const clickDeleteButton = () => {
  cy.getBySel('delete-btn').click();
};

export const deleteSystem = (system: System) => {
  clickSystemOptions(system);
  clickSystemDeleteOption();
  clickDeleteButton();
};
