<template>
    <v-row class="ma-2">
        <v-alert v-if="AdminStore.UM" type="warning" variant="outlined" density="compact" text="Имеются непрочитанные комментарии"></v-alert>
    </v-row>
    <v-row class="ma-0 pa-0">
        <v-col cols="3"></v-col>
        <v-col cols="6">
            <v-btn-toggle v-model="toggle">
                <v-btn color="teal-darken-1" @click="AdminStore.filterAll()">все</v-btn>
                <v-btn color="teal-darken-1" @click="AdminStore.filterNotEmployee()">без исполнителя</v-btn>
                <v-btn color="teal-darken-1" @click="AdminStore.filterFinished()">завершенные</v-btn>
                <!-- <v-btn color="teal-darken-1">просроченные</v-btn> -->
            </v-btn-toggle>
        </v-col>
        <v-col cols="3"></v-col>
    </v-row>
    <v-row class="pl-5 pr-5">
        <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in AdminStore.showTickets"
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
                            <v-icon v-if="ticket.unreadMessages" color="red" icon="mdi-comment-alert" class="mr-3"></v-icon>
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
import { ref } from 'vue';
const AuthStore = useAuthStore();
const AdminStore = useAdminStore();
const toggle = ref(0)
AdminStore.getAdminTickets(AuthStore.credentials.token);
</script>