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
          :data-test="item.testId+'_system_options'"
          v-if="item.type === 'system'"
          :system="item.system"
          @systemUpdated="systemUpdated"
          @systemDeleted="systemDeleted"
        />
        <NamespaceMenu
          v-if="item.type === 'namespace'"
          :data-test="item.testId+'_namespace_options'"
          :namespace="item.namespace"
          :system="item.system"
          @deleted="namespaceDeleted"
          @updated="namespaceUpdated"
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
import { System } from '@/system/entity';
import { Namespace } from '@/namespace/entity';
import { Node, NodeType } from '../entity';
import { QueryParams } from '../helper';
import CreateSystemDialog from '../../system/components/CreateDialog.vue';
import CreateNamespaceDialog from '../../namespace/components/CreateDialog.vue';
import SystemMenu from './SystemMenu.vue';
import NamespaceMenu from './NamespaceMenu.vue';

@Component({
  components: {
    CreateSystemDialog,
    CreateNamespaceDialog,
    SystemMenu,
    NamespaceMenu,
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
      name: this.$t('new-namespace').toString(),
      type: 'createNewNamespaceBtn',
      color: 'green',
      testId: `${system.name}_createNewNamespaceBtn`,
      system,
    };
  }

  createNewSystemBtn: Node ={
    id: 'createNewSystemBtn',
    name: this.$t('new-system').toString(),
    type: 'createNewSystemBtn',
    testId: 'createNewSystemBtn',
    color: 'green',
  }

  inc = 1

  generateRandom(): string {
    this.inc += 1;
    return this.inc.toString();
  }

  getNamespaceChildren(system: System, namespace: Namespace): Node[] {
  return [
    {
      id: this.generateRandom(),
      name: 'routes',
      type: 'routes',
      system,
      namespace,
      testId: `${system.name}_${namespace.name}_routes`,
    }, {
      id: this.generateRandom(),
      name: 'groups',
      type: 'groups',
      system,
      namespace,
      testId: `${system.name}_${namespace.name}_groups`,
    }, {
      id: this.generateRandom(),
      name: 'users',
      type: 'users',
      system,
      namespace,
      testId: `${system.name}_${namespace.name}_users`,
    },
    {
      id: this.generateRandom(),
      name: 'tags',
      type: 'tags',
      system,
      namespace,
      testId: `${system.name}_${namespace.name}_tags`,
    },
  ];
  }

  async loadChildren(node: Node) {
    if (node.type === 'system' && node.system) {
      node.children = await this.getNamespacesBelongToSystem(node.system);
    }
   if (node.type === 'namespace' && node.namespace && node.system) {
     node.children = this.getNamespaceChildren(node.system, node.namespace);
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

  async namespaceDeleted(namespaceId: number, systemId: number) {
     this.items.forEach((system) => {
      if (system.system?.id === systemId) {
        // eslint-disable-next-line no-unused-expressions
        system.children?.forEach((namespace, namespaceIndex) => {
          if (namespace.namespace?.id === namespaceId) {
            // eslint-disable-next-line no-unused-expressions
            system.children?.splice(namespaceIndex, 1);
          }
        });
      }
      return system;
    });
  }

  async namespaceUpdated(ns: Namespace, system: System) {
     this.items.forEach((systemNode) => {
      if (systemNode.system?.id === system.id) {
        // eslint-disable-next-line no-unused-expressions
        systemNode.children = systemNode.children?.map((namespaceNode) => {
          if (namespaceNode.namespace) {
            if (namespaceNode.namespace.id === ns.id) {
              namespaceNode = this.convertNamespaceToNode(ns, system);
            }
          }
          return namespaceNode;
        });
      }

      return systemNode;
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
    const nodes = systems.map((system) => this.convertSystemToNode(system));
    nodes.push(this.createNewSystemBtn);
    return nodes;
  }

  convertSystemToNode(system: System): Node {
    return {
      id: system.name,
      name: system.name,
      type: 'system',
      system,
      testId: system.name,
      children: [],
    };
  }

  convertNamespaceToNode(namespace: Namespace, system: System): Node {
    return {
      namespace,
      system,
      type: 'namespace',
      name: namespace.name,
      id: `${system.name}_${namespace.name}`,
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
      this.showCreateNamespaceDialog = true;
      this.clearActiveNodes();
    }

    const node = nodes[0];
    if (node.type === 'routes') {
      await this.pushToPage('Routes', node);
    } else if (node.type === 'groups') {
      await this.pushToPage('Groups', node);
    } else if (node.type === 'users') {
      await this.pushToPage('Users', node);
    } else if (node.type === 'tags') {
      await this.pushToPage('Tags', node);
    }
    this.clearActiveNodes();
  }

  async pushToPage(name: string, node: Node) {
    if (this.$route.name === name) return;

    return this.$router.push({
      name,
      // @ts-ignore
      query: { ...this.formQueryParams(node.system, node.namespace), lang: this.$route.query.lang },
    });
  }

  formQueryParams(system: System, namespace: Namespace): QueryParams {
    return {
      systemId: system.id,
      systemName: system.name,
      namespaceId: namespace.id,
      namespaceName: namespace.name,
    };
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
  this.iconMap.set('tags', 'mdi-tag-multiple');

  this.items = await this.getSystems();
}

  tree= []
}
</script>

<style scoped lang="scss">
.tree-view-container {
  background: aliceblue;
}
</style>

<i18n>
{
  "en": {
    "new-system": "New system",
    "new-namespace": "New service"
  },
  "ru": {
    "new-system": "Новая система",
    "new-namespace": "Новый сервис"
  }
}
</i18n>
