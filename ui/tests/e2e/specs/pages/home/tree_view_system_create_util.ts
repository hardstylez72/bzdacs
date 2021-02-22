import { System } from '../../../../../src/views/system/entity';
import { assertSystemIsShown, buildSystem, fillForm } from './tree_view_system_util';

const clickNewSystemButton = () => {
  cy.getBySel('createNewSystemBtn').click();
};

const clickSaveSystemButton = () => {
  cy.getBySel('save-system-btn').click();
};

const createSystem = (system?: System): System => {
  if (!system) {
    // eslint-disable-next-line no-param-reassign
    system = buildSystem();
  }
  clickNewSystemButton();
  fillForm(system);
  clickSaveSystemButton();
  cy.get('.v-dialog').should('not.be.visible');
  // @ts-ignore
  return system;
};

// eslint-disable-next-line import/prefer-default-export
export const createValidSystem = (system?: System): System => {
  const created = createSystem(system);
  assertSystemIsShown(created);
  return created;
};
