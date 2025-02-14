<template>
    <v-row class="pa-5">
        <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in ArchiveStore.archiveTickets"
                    :key="ticket.id"
                    class="mt-3"
                    :to="'/full/' + ticket.id + '/archive'"
                >
                    <v-card-item>
                        <template v-slot:prepend>
                            <v-icon v-if="ticket.important" color="red" icon="mdi-comment-alert"></v-icon>
                            <v-icon v-if="!ticket.important" color="green" icon="mdi-comment-check"></v-icon>
                        </template>
                        <v-card-subtitle>Пользователь: {{ ticket.userInitials }} ({{ ticket.userDepartment }}) / Исполнитель: {{ ticket.employeeInitials }} / Дата обращения: {{ AuthStore.myDateTimeFormat(ticket.createdAt) }}<span v-if="ticket.finishBefore != null"> / Срок исполнения: {{ AuthStore.myDateFormat(ticket.finishBefore) }}</span></v-card-subtitle>
                        <v-card-text>Текст обращения: {{ ticket.text }}</v-card-text>
                    </v-card-item>
        </v-card>
    </v-row>
</template>

<script setup>
import { useArchiveStore } from '../stores/ArchiveStore';
import { useAuthStore } from '../stores/AuthStore';
const ArchiveStore = useArchiveStore();
const AuthStore = useAuthStore();

ArchiveStore.getArchiveTickets(AuthStore.credentials.token)
</script>