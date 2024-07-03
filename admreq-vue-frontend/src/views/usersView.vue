<template>
    <v-row>
        <v-col cols="9">
            <div class="pa-3">
                <v-list lines="one">
                    <v-list-item 
                        v-for="user in UsersStore.users"
                        :key="user.uuid"
                    >
                        <v-list-item-title>
                            {{ user.firstname }} {{ user.lastname }} ({{ user.Role }})
                        </v-list-item-title>
                        <v-list-item-subtitle>
                            {{ user.uuid }}
                        </v-list-item-subtitle>

                    </v-list-item>
                    <!-- {{ AdminStore.users }} -->
                </v-list>
            </div>
        </v-col>
        <v-col>
            <div class="pa-3">
                <v-card class="pa-3" variant="elevated" elevation="16" color="teal-lighten-4">
                <v-form fast-fail @submit.prevent="UsersStore.createUser(firstname, lastname, surname, department, username, password, role, AuthStore.credentials.token)">
                    <v-text-field v-model="firstname" label="Имя"></v-text-field>
                    <v-text-field v-model="lastname" label="Фамилия"></v-text-field>
                    <v-text-field v-model="surname" label="Отчество"></v-text-field>
                    <v-text-field v-model="department" label="Отдел"></v-text-field>
                    <v-text-field v-model="username" label="Логин"></v-text-field>
                    <v-text-field v-model="password" label="Пароль"></v-text-field>
                    <v-text-field v-model="role" label="Роль"></v-text-field>
                    <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
                </v-form>
                <span class="d-flex align-center justify-center ma-3 text-red">{{ UsersStore.usersErrors }}</span>
                </v-card>
            </div>        
        </v-col>
    </v-row>
</template>

<script setup>
import { useAuthStore } from '../stores/AuthStore';
import { useUsersStore } from '../stores/UsersStore';
import { ref } from 'vue';

const AuthStore = useAuthStore();
const UsersStore = useUsersStore();

UsersStore.getUsers(AuthStore.credentials.token);

const firstname = ref("");
const lastname = ref("");
const surname = ref("");
const department = ref("");
const username = ref("");
const password = ref("");
const role = ref("");
</script>