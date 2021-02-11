import { createValidSystem } from './tree_view_system_create_util';
import { createValidNamespace } from './tree_view_namespace_create_util';

describe('Guest page', () => {
  it('Creates system and validates it is shown in a list', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const namespace = createValidNamespace(system);
  });
});
