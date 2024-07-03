<template>
    <v-row>
        <v-col cols="3"
            v-for="ticket in EmployeeStore.employeeTickets"
            :key="ticket.id"
        >
            <v-card class="ma-3">
                <v-card-item>
                    <v-card-subtitle>{{ ticket.userInitials }}</v-card-subtitle>
                </v-card-item>
                <v-card-text>
                    <div>{{ ticket.text }}</div>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="teal-darken-1" variant="outlined" class="mt-2" text="Подробнее"></v-btn>
                </v-card-actions>
                </v-card>
            </v-col>
        
        </v-row>
</template>

<script setup>
import { useEmployeeStore } from '../stores/EmployeeStore';
import { useAuthStore } from '../stores/AuthStore';
const EmployeeStore = useEmployeeStore();
const AuthStore = useAuthStore();

EmployeeStore.getEmployeeTickets(AuthStore.credentials.token)
</script>