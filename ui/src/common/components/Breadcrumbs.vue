<template>
  <div>
    <v-breadcrumbs
      :items="items"
      divider="->"
    >
      <template v-slot:item="{item}">
        <span style="font-size: 20px"><h3>{{item.text}}</h3></span>
      </template>
    </v-breadcrumbs>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';

@Component

export default class Breadcrumbs extends Vue {
  @Prop() items!: []
}
</script>

<style scoped lang="scss">

</style>
