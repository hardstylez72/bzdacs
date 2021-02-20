import { createValidSystem } from './tree_view_system_create_util';
import { createValidNamespace } from './tree_view_namespace_create_util';
import {deleteNamespace} from "./tree_view_namespace_delete_util";
import {
  assertNamespaceIsNotShown,
  assertNamespaceIsShown,
  buildNamespace
} from "./tree_view_namespace_util";

import {editNamespace} from "./tree_view_namespace_edit_util";
import {assertSystemIsNotShown, assertSystemIsShown} from "./tree_view_system_util";
import {deleteSystem} from "./tree_view_system_delete_util";

describe('TreeView namespace', () => {
  it('Creates namespace and validates it is shown in a list', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const namespace = createValidNamespace(system);
  });

  it('Creates namespace then deletes it', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const namespace = createValidNamespace(system);
    deleteNamespace(system, namespace)
    assertNamespaceIsNotShown(namespace)
  });

  it('Creates system then edits it', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const have = createValidNamespace(system);
    const want = buildNamespace();
    const got = editNamespace(system, have, want);
    assertNamespaceIsShown(got);
  });

  it.only('Acceptance test', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const have = createValidNamespace(system);
    const want = buildNamespace();
    const got = editNamespace(system, have, want);
    assertNamespaceIsShown(got);
    deleteNamespace(system, got);
    assertNamespaceIsNotShown(got);
    assertSystemIsShown(system);
    deleteSystem(system);
    assertSystemIsNotShown(system);
  });
});
