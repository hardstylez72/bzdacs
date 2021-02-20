import { createValidSystem } from './tree_view_system_create_util';
import { assertSystemIsNotShown, assertSystemIsShown, buildSystem } from './tree_view_system_util';
import { editSystem } from './tree_view_system_edit_util';
import { deleteSystem } from './tree_view_system_delete_util';

describe('TreeView system', () => {
  it('Creates system and validates it is shown in a list', () => {
    cy.visit('/?lang=en');
    createValidSystem();
  });

  it('Creates system then deletes it', () => {
    cy.visit('/?lang=en');
    const have = createValidSystem();
    assertSystemIsShown(have);
    deleteSystem(have);
    assertSystemIsNotShown(have);
  });

  it('Creates system then edits it', () => {
    cy.visit('/?lang=en');
    const have = createValidSystem();
    const want = buildSystem();
    const got = editSystem(have, want);
    assertSystemIsShown(got);
  });

  it.only('Acceptance test', () => {
    cy.visit('/?lang=en');
    const have = createValidSystem();
    assertSystemIsShown(have);
    const want = buildSystem();
    const got = editSystem(have, want);
    assertSystemIsShown(got);
    deleteSystem(got);
    assertSystemIsNotShown(got);
  });
});
