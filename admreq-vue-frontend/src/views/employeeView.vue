<template>
    <v-row class="pa-5">
        <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in EmployeeStore.employeeTickets"
                    :key="ticket.id"
                    class="mt-3"
                    :to="'/full/' + ticket.id + '/employee'"
                >
                    <v-card-item>
                        <template v-slot:prepend>
                            <v-icon v-if="ticket.important" color="red" icon="mdi-comment-alert"></v-icon>
                            <v-icon v-if="!ticket.important" color="green" icon="mdi-comment-check"></v-icon>
                        </template>
                        <v-card-subtitle>Пользователь: {{ ticket.userInitials }} / Исполнитель: {{ ticket.employeeInitials }} / Дата обращения: {{ AuthStore.myDateTimeFormat(ticket.createdAt) }}</v-card-subtitle>
                        <v-card-text>Текст обращения: {{ ticket.text }}</v-card-text>
                    </v-card-item>
        </v-card>
    </v-row>
</template>

<script setup>
import { useEmployeeStore } from '../stores/EmployeeStore';
import { useAuthStore } from '../stores/AuthStore';
const EmployeeStore = useEmployeeStore();
const AuthStore = useAuthStore();

EmployeeStore.getEmployeeTickets(AuthStore.credentials.token)
</script>