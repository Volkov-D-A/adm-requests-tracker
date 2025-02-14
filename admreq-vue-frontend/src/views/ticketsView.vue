<template>
    <v-row>
        <v-col cols="9">
            <v-row class="pa-5">
                <v-alert v-if="TicketStore.UM" type="warning" variant="outlined" density="compact" text="Имеются непрочитанные комментарии"></v-alert>
                <v-card 
                    style="min-width: 100%;"
                    v-for="ticket in TicketStore.userTickets"
                    :key="ticket.id"
                    class="mt-3"
                    :to="'/full/' + ticket.id + '/user'"
                >
                    <v-card-item>
                        <template v-slot:append>
                            <v-icon v-if="ticket.unreadMessages" color="red" icon="mdi-comment-alert" class="mr-3"></v-icon>
                            <v-icon v-if="ticket.finished" color="green" icon="mdi-clock-check"></v-icon>
                            <v-icon v-if="!ticket.finished" color="purple" icon="mdi-clock"></v-icon>
                        </template>
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
                    <v-form fast-fail @submit.prevent="TicketStore.createTicket(text, AuthStore.credentials.token, department)">
                        <v-textarea counter auto-grow variant="outlined" background-color="blue-lighten-5" v-model="text" label="Текст обращения"></v-textarea>
                        <v-select
                        label="Отдел"
                        v-model="department"
                        :items="TicketStore.departments"
                        item-title="department"
                        item-value="uuid"
                        >
                        </v-select>
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
const department = ref("");
TicketStore.getUserTickets(AuthStore.credentials.token);
TicketStore.getDepartments(AuthStore.credentials.token);
</script>