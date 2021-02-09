<template>
    <v-menu  right>
      <template v-slot:activator="{ on, attrs }">
        <v-btn dark icon v-bind="attrs" v-on="on">
          <v-icon color="grey">mdi-dots-vertical</v-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item @click="clickedUpdate">
          <v-list-item-title>Редактировать</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-pencil</v-icon>
          </v-list-item-icon>
        </v-list-item>
        <v-list-item @click="clickedDelete">
          <v-list-item-title>Удалить</v-list-item-title>
          <v-list-item-icon>
            <v-icon>mdi-delete</v-icon>
          </v-list-item-icon>
        </v-list-item>
      </v-list>
    </v-menu>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';
import { System } from '@/views/system/service';

@Component

export default class SystemMenu extends Vue {
  @Prop() system!: System

  clickedDelete() {
    this.$emit('clickedDelete', this.system);
  }

  clickedUpdate() {
    this.$emit('clickedUpdate', this.system);
  }
}
</script>

<style scoped lang="scss">

</style>
