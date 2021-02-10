<template>
  <div class="tree-view-container">
    <v-treeview
      data-test="tree-view"
      v-model="tree"
      :items="systems"
      activatable
      :active.sync="active"
      item-key="name"
      open-on-click
      return-object
      @update:active="nodeWasSelected"
    >
      <template v-slot:prepend="{ item, open }">
        <v-icon :color="item.color ? item.color : ''">{{ iconMap.get(item.type) }}</v-icon>
      </template>
      <template v-slot:label="{ item, open }">
        <span :data-test="item.id">{{item.name}}</span>
      </template>
      <template v-slot:append="{ item, open }">
        <SystemMenu :data-test="item.id+'_system_options'" v-if="item.type === 'system'" :system="item.origin"/>
      </template>
    </v-treeview>
    <CreateSystemDialog v-model="showCreateSystemDialog" @itemCreated="systemCreated"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { System } from '@/views/system/service';
import { Node, NodeType } from '../entity';
import CreateSystemDialog from '../../system/components/CreateSystemDialog.vue';
import SystemMenu from './SystemMenu.vue';

@Component({
  components: {
    CreateSystemDialog,
    SystemMenu,
  },
})

export default class TreeView extends Vue {
  showCreateSystemDialog = false

  iconMap = new Map<NodeType, string>()

  active = []

  createNewNamespaceBtn: Node = {
    id: 'createNewNamespaceBtn',
    name: 'Новое пространство имен',
    type: 'createNewNamespaceBtn',
    color: 'green',
  }

  createNewSystemBtn: Node ={
    id: 'createNewSystemBtn',
    name: 'Новая система',
    type: 'createNewSystemBtn',
    color: 'green',
  }

  namespaceChildren: Node[] = [
    {
      name: 'routes',
      type: 'routes',
    }, {
      name: 'groups',
      type: 'groups',
    }, {
      name: 'users',
      type: 'users',
    },
  ]

  systemCreated(system: System) {
    this.$store.direct.dispatch.system.GetList();
  }

  get systems(): readonly Node[] {
    const base = this.$store.direct.getters.system.getItems;

    const nodes = base.map((sys) => {
      const node: Node = {
        id: sys.name,
        name: sys.name,
        type: 'system',
        origin: sys,
      };
      return node;
    });
    nodes.push(this.createNewSystemBtn);
    return nodes;
  }

  nodeWasSelected(nodes: Node[]) {
    if (this.isButtonCreateSystemWasPressed(nodes)) {
     this.showCreateSystemDialog = true;
      this.clearActiveNodes();
    }
  }

  clearActiveNodes() {
    this.active = [];
  }

  isButtonCreateSystemWasPressed(nodes: Node[]): boolean {
    if (!nodes) {
      return false;
    }
    if (nodes.length === 1) {
      if (nodes[0].type === 'createNewSystemBtn') {
        return true;
      }
    }
    return false;
  }

created() {
  this.$store.direct.dispatch.system.GetList();
  this.iconMap.set('groups', 'mdi-account-group');
  this.iconMap.set('namespace', 'mdi-card');
  this.iconMap.set('routes', 'mdi-map-marker-path');
  this.iconMap.set('system', 'mdi-database');
  this.iconMap.set('users', 'mdi-account');
  this.iconMap.set('createNewSystemBtn', 'mdi-database-plus');
  this.iconMap.set('createNewNamespaceBtn', 'mdi-card-plus');
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
