<template>
  <v-app v-if=!UserStore.authorized>
    <authForm/>
  </v-app>
  <v-app v-else>
    <MainWindow/>
  </v-app>
</template>

<script setup>
import { useUserStore } from './stores/UserStore';
const UserStore = useUserStore();

import authForm from './components/authForm.vue';
import MainWindow from './components/mainWindow.vue';
</script>

<style scoped>

</style>
