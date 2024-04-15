<template>
    <div>{{ AdminStore.users }}</div>

    <div>
        <v-form fast-fail @submit.prevent="AdminStore.createUser(firstname, lastname, username, password, role, UserStore.credentials.token)">
            <v-text-field v-model="firstname" label="Имя"></v-text-field>
            <v-text-field v-model="lastname" label="Фамилия"></v-text-field>
            <v-text-field v-model="username" label="Логин"></v-text-field>
            <v-text-field v-model="password" label="Пароль"></v-text-field>
            <v-text-field v-model="role" label="Роль"></v-text-field>
            <v-btn type="submit" color="primary" block class="mt-2">Sign in</v-btn>
        </v-form>
        <span class="d-flex align-center justify-center ma-3 text-red">{{ AdminStore.usersErrors }}</span>
    </div>
</template>

<script setup>
import { useUserStore } from '../stores/UserStore';
import { useAdminStore } from '../stores/AdminStore';
import { ref } from 'vue';
const UserStore = useUserStore();
const AdminStore = useAdminStore();
AdminStore.getUsers(UserStore.credentials.token);
const firstname =ref("");
const lastname =ref("");
const username = ref("");
const password = ref("");
const role =ref("");
</script>