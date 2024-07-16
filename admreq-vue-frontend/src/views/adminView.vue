<template>
    <v-row class="pa-5">
        <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in AdminStore.adminTickets"
                    :key="ticket.id"
                    class="mt-3"
                    :to="'/full/' + ticket.id + '/admin'"
                >
                    <v-card-item>
                        <template v-slot:prepend>
                            <v-icon v-if="ticket.important" color="red" icon="mdi-comment-alert"></v-icon>
                            <v-icon v-if="!ticket.important" color="green" icon="mdi-comment-check"></v-icon>
                        </template>
                        <template v-slot:append>
                            <v-icon v-if="ticket.finished" color="green" icon="mdi-clock-check"></v-icon>
                            <v-icon v-if="!ticket.finished" color="purple" icon="mdi-clock"></v-icon>
                        </template>
                        <v-card-subtitle>Пользователь: {{ ticket.userInitials }} ({{ ticket.userDepartment }}) / Исполнитель: {{ ticket.employeeInitials }} / Дата обращения: {{ AuthStore.myDateTimeFormat(ticket.createdAt) }}</v-card-subtitle>
                        <v-card-text>Текст обращения: {{ ticket.text }}</v-card-text>
                    </v-card-item>
        </v-card>
    </v-row>
</template>

<script setup>
import { useAuthStore } from '../stores/AuthStore';
import { useAdminStore } from '../stores/AdminStore';
const AuthStore = useAuthStore();
const AdminStore = useAdminStore();
AdminStore.getAdminTickets(AuthStore.credentials.token);
</script>