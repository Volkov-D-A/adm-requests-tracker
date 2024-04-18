<template>
    <v-row>
        <v-col cols="9">
            <div class="pa-3">
                <v-list lines="one">
                    <v-list-item 
                        v-for="user in AdminStore.users"
                        :key="user.uuid"
                    >
                        <v-list-item-title>
                            {{ user.firstName }} {{ user.lastName }} ({{ user.Role }})
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
                <v-form fast-fail @submit.prevent="AdminStore.createUser(firstname, lastname, username, password, role, UserStore.credentials.token)">
                    <v-text-field v-model="firstname" label="Имя"></v-text-field>
                    <v-text-field v-model="lastname" label="Фамилия"></v-text-field>
                    <v-text-field v-model="username" label="Логин"></v-text-field>
                    <v-text-field v-model="password" label="Пароль"></v-text-field>
                    <v-text-field v-model="role" label="Роль"></v-text-field>
                    <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
                </v-form>
                <span class="d-flex align-center justify-center ma-3 text-red">{{ AdminStore.usersErrors }}</span>
                </v-card>
            </div>        
        </v-col>
    </v-row>
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