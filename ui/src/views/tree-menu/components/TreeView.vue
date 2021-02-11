<template>
  <div class="tree-view-container">
    <v-treeview
      data-test="tree-view"
      :items="items"
      activatable
      :active.sync="active"
      open-on-click
      return-object
      item-key="id"
      :multiple-active="false"
      @update:active="nodeWasSelected"
      :load-children="loadChildren"
    >
      <template v-slot:prepend="{ item, open }">
        <v-icon :color="item.color ? item.color : ''">{{ iconMap.get(item.type) }}</v-icon>
      </template>
      <template v-slot:label="{ item, open }">
        <span :data-test="item.testId">{{item.name}}</span>
      </template>
      <template v-slot:append="{ item, open }">
        <SystemMenu
          :data-test="item.id+'_system_options'"
          v-if="item.type === 'system'"
          :system="item.system"
          @systemUpdated="systemUpdated"
          @systemDeleted="systemDeleted"
        />
      </template>
    </v-treeview>
    <CreateSystemDialog
      v-model="showCreateSystemDialog"
      @itemCreated="systemCreated"
    />
    <CreateNamespaceDialog
      v-model="showCreateNamespaceDialog"
      :system-id="activeSystemId"
      @itemCreated="namespaceCreated"
    />
  </div>

</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { System } from '@/views/system/service';
import { Namespace } from '@/views/namespace/service';
import { Node, NodeType } from '../entity';
import CreateSystemDialog from '../../system/components/CreateDialog.vue';
import CreateNamespaceDialog from '../../namespace/components/CreateDialog.vue';
import SystemMenu from './SystemMenu.vue';

@Component({
  components: {
    CreateSystemDialog,
    CreateNamespaceDialog,
    SystemMenu,
  },
})

export default class TreeView extends Vue {
  showCreateSystemDialog = false

  showCreateNamespaceDialog = false

  activeSystemId = -1

  iconMap = new Map<NodeType, string>()

  items: Node[] = []

  active: Node[]= []

  createNewNamespaceBtn(system: System): Node {
    return {
      id: this.generateRandom(),
      name: 'Новое пространство имен',
      type: 'createNewNamespaceBtn',
      color: 'green',
      testId: `${system.name}_createNewNamespaceBtn`,
      system,
    };
  }

  createNewSystemBtn: Node ={
    id: 'createNewSystemBtn',
    name: 'Новая система',
    type: 'createNewSystemBtn',
    testId: 'createNewSystemBtn',
    color: 'green',
  }

  inc = 1

  generateRandom(): string {
    this.inc += 1;
    return this.inc.toString();
  }

  namespaceChildren(system: System, namespace: Namespace): Node[] {
  return [
    {
      id: this.generateRandom(),
      name: 'routes',
      type: 'routes',
      system,
      namespace,
      testId: 'routes',
    }, {
      id: this.generateRandom(),
      name: 'groups',
      type: 'groups',
      system,
      namespace,
      testId: 'groups',
    }, {
      id: this.generateRandom(),
      name: 'users',
      type: 'users',
      system,
      namespace,
      testId: 'users',
    },
  ];
  }

  async loadChildren(node: Node) {
    if (node.type === 'system' && node.system) {
      node.children = await this.getNamespacesBelongToSystem(node.system);
    }
   if (node.type === 'namespace' && node.namespace && node.system) {
     node.children = this.namespaceChildren(node.system, node.namespace);
   }
  }

  async namespaceCreated(namespace: Namespace, systemId: number) {
    this.items = this.items.map((system) => {
      if (system.system?.id === systemId) {
        // eslint-disable-next-line no-unused-expressions
        system.children?.pop();
        // eslint-disable-next-line no-unused-expressions
        system.children?.push(this.convertNamespaceToNode(namespace, system.system));
        // eslint-disable-next-line no-unused-expressions
        system.children?.push(this.createNewNamespaceBtn(system.system));
      }
      return system;
    });
  }

  async systemCreated() {
    this.items = await this.getSystems();
  }

  async systemUpdated() {
    this.items = await this.getSystems();
  }

  async systemDeleted() {
    this.items = await this.getSystems();
  }

  async getSystems(): Promise<Node[]> {
    const systems = await this.$store.direct.dispatch.system.GetList();

    const nodes = systems.map((system) => {
      const node: Node = {
        id: system.name,
        name: system.name,
        type: 'system',
        system,
        testId: system.name,
        children: [],
      };
      return node;
    });
    nodes.push(this.createNewSystemBtn);
    return nodes;
  }

  convertNamespaceToNode(namespace: Namespace, system: System): Node {
    return {
      namespace,
      system,
      type: 'namespace',
      name: namespace.name,
      id: this.generateRandom(),
      testId: `${system.name}_${namespace.name}`,
      children: [],
    };
  }

  async getNamespacesBelongToSystem(system: System): Promise<Node[]> {
    const namespaces = await this.$store.direct.dispatch.namespace.GetListBySystemId(system.id);
    const nodes = namespaces.map((namespace) => this.convertNamespaceToNode(namespace, system));
    nodes.push(this.createNewNamespaceBtn(system));
    return nodes;
  }

  async nodeWasSelected(nodes: Node[]) {
    if (!nodes) {
      return false;
    }

    if (!nodes.length) {
      return false;
    }

    if (this.IsPressedButtonWithType(nodes, 'createNewSystemBtn')) {
      this.showCreateSystemDialog = true;
      this.clearActiveNodes();
      return;
    }

    if (this.IsPressedButtonWithType(nodes, 'createNewNamespaceBtn')) {
      if (nodes[0].system?.id) {
        this.activeSystemId = nodes[0].system.id;
      }
      console.log('nodes[0].system?.name', nodes[0].system?.name);
      this.showCreateNamespaceDialog = true;
      this.clearActiveNodes();
    }

    if (nodes[0].type === 'users') {
      await this.$router.push({ query: { users: '34' } });
    }
    this.clearActiveNodes();
  }

  clearActiveNodes() {
    this.$nextTick(() => {
      this.active = [];
    });
  }

  IsPressedButtonWithType(nodes: Node[], type: NodeType): boolean {
    if (nodes.length === 1) {
      if (nodes[0].type === type) {
        return true;
      }
    }
    return false;
  }

async created() {
  this.iconMap.set('groups', 'mdi-account-group');
  this.iconMap.set('namespace', 'mdi-card');
  this.iconMap.set('routes', 'mdi-map-marker-path');
  this.iconMap.set('system', 'mdi-database');
  this.iconMap.set('users', 'mdi-account');
  this.iconMap.set('createNewSystemBtn', 'mdi-database-plus');
  this.iconMap.set('createNewNamespaceBtn', 'mdi-card-plus');

  this.items = await this.getSystems();
}

  tree= []
}
</script>

<style scoped lang="scss">
.tree-view-container {
  background: aliceblue;
  max-width: 400px;
}
</style>
