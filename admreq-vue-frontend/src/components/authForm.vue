<template>
    <div class="d-flex align-center justify-center" style="height: 100vh">
        <v-sheet width="400" class="mx-auto">
            <v-form fast-fail @submit.prevent="AuthStore.getAuth(username, password)">
                <v-text-field v-model="username" label="Имя пользователя"></v-text-field>

                <v-text-field type="password" v-model="password" label="Пароль пользователя"></v-text-field>

                <v-btn type="submit" color="primary" block class="mt-2">Вход</v-btn>
            </v-form>
            <span class="d-flex align-center justify-center ma-3 text-red">{{ AuthStore.userError }}</span>
        </v-sheet>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../stores/AuthStore';

const AuthStore = useAuthStore();
const username = ref("");
const password = ref("");
</script>