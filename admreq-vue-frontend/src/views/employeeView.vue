<template>
    <v-row>
        <v-col cols="3"
            v-for="ticket in UserStore.getEmployeeTickets()"
            :key="ticket.id"
        >
            <v-card class="ma-3">
                <v-card-item>
                    <v-card-subtitle>{{ ticket.userId }}</v-card-subtitle>
                </v-card-item>
                <v-card-text>
                    <div>{{ ticket.text }}</div>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="teal-darken-1" variant="outlined" class="mt-2" text="Подробнее"></v-btn>
                    <v-btn color="deep-orange-darken-1" variant="outlined" class="mt-2" text="Завершить"></v-btn>
                </v-card-actions>
                </v-card>
            </v-col>
        
        </v-row>
    <!--<div>{{ UserStore.getEmployeeTickets() }}</div>-->
</template>

<script setup>
import { useUserStore } from '../stores/UserStore';
import { ref } from 'vue';
const UserStore = useUserStore();
UserStore.getTickets(UserStore.credentials.token)
</script>