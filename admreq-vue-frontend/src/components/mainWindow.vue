<template>
    <v-card>
      <v-layout class="d-flex align-center justify-center">
        <v-navigation-drawer
          expand-on-hover
          rail
        >
          <v-list>
            <v-list-item
              prepend-icon="mdi-account"
              :subtitle='AuthStore.credentials.departmentName'
              :title='AuthStore.credentials.lastname+" "+AuthStore.credentials.firstname[0]+"."+AuthStore.credentials.surname[0]+"."' 
            ></v-list-item>
          </v-list>
  
          <v-divider></v-divider>
  
          <v-list density="compact" nav>
            <v-list-item prepend-icon="mdi-card-text-outline" title="Обращения" to="/"></v-list-item>
            <v-list-item v-if="role === 'admin'" prepend-icon="mdi-account-multiple" title="Пользователи" to="/users"></v-list-item>
            <v-list-item v-if="role != 'user'" prepend-icon="mdi-pipe-wrench" title="В работе" to="/works"></v-list-item>
            <v-list-item v-if="role === 'admin'" prepend-icon="mdi-sitemap" title="Управление" to="/admin"></v-list-item>
            <v-list-item v-if="role === 'admin'" prepend-icon="mdi-archive" title="Архив" to="/archive"></v-list-item>
            <v-list-item prepend-icon="mdi-exit-run" title="Выход" @click="exit()"></v-list-item>
          </v-list>
        </v-navigation-drawer>
  
        <v-main style="min-height: 100vh;">
          <RouterView/>
        </v-main>
      </v-layout>
    </v-card>
</template>

<script setup>
import { useRouter, RouterView } from 'vue-router';

import { useAuthStore } from '../stores/AuthStore';
const AuthStore = useAuthStore();
const router = useRouter();
const role = AuthStore.credentials.Role

function exit() {
  router.push('/')
  AuthStore.authorized = false
}
</script>