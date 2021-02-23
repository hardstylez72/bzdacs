import { createValidRoute } from './routes_create_util';
import {createValidSystem} from "../home/tree_view_system_create_util";
import {createValidNamespace} from "../home/tree_view_namespace_create_util";
import {clickRoutesNode, expandNamespaceNode} from "../home/tree_view_namespace_util";
import {deleteRoute} from "./routes_delete_util";

describe('Routes', () => {

  it('Creates route and validates it is shown in a list', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const namespace = createValidNamespace(system);
    expandNamespaceNode(system, namespace)
    clickRoutesNode(system, namespace)
    createValidRoute(namespace);
  });

  it('Creates route then deletes it', () => {
    cy.visit('/?lang=en');
    const system = createValidSystem();
    const namespace = createValidNamespace(system);
    expandNamespaceNode(system, namespace)
    clickRoutesNode(system, namespace)
    const route = createValidRoute(namespace);
    deleteRoute(route)
  });
  //
  // it.only('Creates route then edits it', () => {
  //   cy.visit('/?lang=en');
  //   const have = createValidSystem();
  //   const want = buildSystem();
  //   const got = editSystem(have, want);
  //   assertSystemIsShown(got);
  // });
  //
  // it.only('Acceptance test', () => {
  //   cy.visit('/?lang=en');
  //   const have = createValidSystem();
  //   assertSystemIsShown(have);
  //   const want = buildSystem();
  //   const got = editSystem(have, want);
  //   assertSystemIsShown(got);
  //   deleteSystem(got);
  //   assertSystemIsNotShown(got);
  // });
});
