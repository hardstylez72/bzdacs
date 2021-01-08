<template>
  <v-card>
    <v-tabs v-model="tab" align-with-title @change="tabChanged">
      <v-tabs-slider/>
      <v-tab v-for="item in tabs" :key="item">{{ item }}</v-tab>
    </v-tabs>

    <v-tabs-items v-model="tab">
      <v-tab-item v-for="item in tabs" :key="item">
        <div v-if="item === tabs[0]">
          <TabRouteTable/>
        </div>
        <div v-if="item === tabs[1]">
          <TapGroupTable/>
        </div>
        <div v-if="item === tabs[2]">
          <TapUserTable/>
        </div>
        <div v-if="item === tabs[3]">
          <TabTagsTable/>
        </div>
      </v-tab-item>
    </v-tabs-items>

  </v-card>
</template>

<script lang="ts">
import {
  Component, Prop, Vue, Watch,
} from 'vue-property-decorator';

import TabRouteTable from '@/views/route/components/TabRouteTable.vue';
import TapGroupTable from '@/views/group/components/TapGroupTable.vue';
import TapUserTable from '@/views/user/components/TapUserTable.vue';
import TabTagsTable from '@/views/tag/components/TabTagsTable.vue';

@Component({
  components: {
    TabRouteTable,
    TapGroupTable,
    TapUserTable,
    TabTagsTable,
  },

})
export default class MainTabs extends Vue {
  @Prop() private msg!: string;

  tab = 0;

  tabs: string[] = [
    'routes', 'groups', 'users', 'tags',
  ];

  @Watch('$route')
  RouteUpdate() {
    const tabNumber = this.getTabNumberFromUrlQueryParams();
    if (this.tab !== tabNumber) {
      this.tab = tabNumber;
    }
  }

  mounted() {
    const tabNumber = this.getTabNumberFromUrlQueryParams();
    if (this.tab !== tabNumber) {
      this.tab = tabNumber;
    }

    const { tab } = this.$route.query;
    if (tab !== this.tabs[tabNumber]) {
      this.$router.push({ query: { tab: this.tabs[tabNumber] } });
    }
  }

  tabChanged(tabNumber: number) {
    if (tabNumber !== this.getTabNumberFromUrlQueryParams()) {
      this.$router.push({ query: { tab: this.tabs[tabNumber] } });
    }
  }

  getTabNumberFromUrlQueryParams(): number {
    const { tab } = this.$route.query;
    let tabNumber = 0;
    if (tab) {
      this.tabs.some((itemTabName: string, itemTabNumber: number) => {
        if (itemTabName === tab) {
          tabNumber = itemTabNumber;
          return true;
        }
        return false;
      });
    }
    return tabNumber;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style  lang="scss">

</style>
