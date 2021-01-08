<template>
  <v-app>
    <v-app-bar app color="primary" dark><h1>BZDACS</h1></v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import MainTabs from '@/views/base/pages/Main.vue';

export default Vue.extend({
  name: 'App',
});
</script>

<style scoped lang="css">
.v-main {
  font-family: "Lobster";
  width: 99%;
  margin-left: 0.5%;
  margin-top: 0.5%;
}

@font-face {
  font-family: "Lobster";
  src: local("Lobster"),
  url(../public/fonts/lobster.ttf) format("truetype");
}
</style>
