<template>
  <v-app v-if=!AuthStore.authorized>
    <authForm/>
  </v-app>
  <v-app v-else>
    <MainWindow/>
  </v-app>
</template>

<script setup>
import { useAuthStore } from './stores/AuthStore';
const AuthStore = useAuthStore();

import authForm from './components/authForm.vue';
import MainWindow from './components/mainWindow.vue';
</script>

<style scoped>

</style>
