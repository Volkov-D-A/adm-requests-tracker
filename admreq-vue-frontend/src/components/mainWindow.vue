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
              :subtitle='UserStore.credentials.login'
              :title='UserStore.credentials.firstName+" "+UserStore.credentials.lastName' 
            ></v-list-item>
          </v-list>
  
          <v-divider></v-divider>
  
          <v-list density="compact" nav>
            <v-list-item prepend-icon="mdi-card-text-outline" title="Обращения" to="/"></v-list-item>
            <v-list-item prepend-icon="mdi-account-multiple" title="Пользователи" to="/users"></v-list-item>
          </v-list>
        </v-navigation-drawer>
  
        <v-main style="height: 100vh;">
          <RouterView/>
        </v-main>
      </v-layout>
    </v-card>
</template>

<script setup>

import { RouterView } from 'vue-router';
import { useUserStore } from '../stores/UserStore';
const UserStore = useUserStore();

</script>