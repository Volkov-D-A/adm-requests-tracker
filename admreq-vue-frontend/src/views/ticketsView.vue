<template>
    <v-row>
        <v-col cols="9">
            <v-row class="pa-5">
                <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in TicketStore.userTickets"
                    :key="ticket.id"
                    class="mt-3"
                    :to="'/full/' + ticket.id"
                >
                    <v-card-item>
                        <v-card-subtitle>Исполнитель: {{ ticket.employeeInitials }} / Дата обращения: {{ AuthStore.myDateTimeFormat(ticket.createdAt) }}</v-card-subtitle>
                        <v-card-text>Текст обращения: {{ ticket.text }}</v-card-text>
                    </v-card-item>
                    
                </v-card>
            </v-row> 
        </v-col>
        <v-col>
            <div class="pa-3">
                <v-card class="pa-3" variant="elevated" elevation="16" color="teal-lighten-4">
                    <span class="ma-2">Направить обращение:</span><br><br>
                    <v-form fast-fail @submit.prevent="TicketStore.createTicket(text, AuthStore.credentials.token)">
                        <v-textarea counter auto-grow variant="outlined" background-color="blue-lighten-5" v-model="text" label="Текст обращения"></v-textarea>
                        <v-btn type="submit" color="teal-darken-1" block class="mt-2">Отправить</v-btn>
                    </v-form>
                </v-card>
            </div>
        </v-col>
    </v-row>
</template>

<script setup>
import { useTicketStore } from '../stores/TicketStore';
import { useAuthStore } from '../stores/AuthStore';
import { ref } from 'vue';
const TicketStore = useTicketStore();
const AuthStore = useAuthStore();
const text = ref("");
TicketStore.getUserTickets(AuthStore.credentials.token)
</script>