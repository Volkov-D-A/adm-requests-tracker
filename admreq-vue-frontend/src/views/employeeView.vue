<template>
    <v-row class="pa-5">
        <v-alert v-if="EmployeeStore.UM" type="warning" variant="outlined" density="compact" text="Имеются непрочитанные комментарии"></v-alert>
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
                        <template v-slot:append>
                            <v-icon v-if="ticket.unreadMessages" color="red" icon="mdi-comment-alert" class="mr-3"></v-icon>
                        </template>
                        <v-card-subtitle>Пользователь: {{ ticket.userInitials }} ({{ ticket.userDepartment }}) / Исполнитель: {{ ticket.employeeInitials }} / Дата обращения: {{ AuthStore.myDateTimeFormat(ticket.createdAt) }}<span v-if="ticket.finishBefore != null"> / Срок исполнения: {{ AuthStore.myDateFormat(ticket.finishBefore) }}</span></v-card-subtitle>
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