<template>
    <v-row>
        <v-col cols="3"
            v-for="ticket in UserStore.ticketsNotEmployee"
            :key="ticket.id"
        >
            <v-card class="ma-3">
                <v-card-item>
                    <v-card-subtitle>{{ ticket.userId }}</v-card-subtitle>
                </v-card-item>
                <v-card-text>
                    <div>{{ ticket.text }}</div>
                    <v-form fast-fail @submit.prevent="UserStore.setEmployee(ticket.id, ticket.employeeId, UserStore.credentials.token)">
                        <!-- <v-text-field v-model="ticket.employeeId" label="Id исполнителя"></v-text-field> -->
                        <v-select v-model="ticket.employeeId" :items="AdminStore.getEmployeeItems()"></v-select>
                        <v-btn type="submit" color="primary" block class="mt-2">Назначить</v-btn>
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-btn color="teal-darken-1" variant="outlined" class="mt-2" text="Подробнее"></v-btn>
                    
                </v-card-actions>
                </v-card>
            </v-col>
        
        </v-row>
    <!-- <div v-for="ticket in UserStore.ticketsNotEmployee" :key="ticket.id">
        <div>User ID: {{ ticket.userId }}</div>
        <div>Text: {{ ticket.text }}</div>
        <v-form fast-fail @submit.prevent="UserStore.setEmployee(ticket.id, ticket.employeeId, UserStore.credentials.token)">
            <v-text-field v-model="ticket.employeeId" label="Id исполнителя"></v-text-field>
            <v-btn type="submit" color="primary" block class="mt-2">Sign in</v-btn>
        </v-form>
    </div> -->

    <!-- <div>{{ UserStore.ticketsNotEmployee }}</div>
    <div>{{ AdminStore.getEmployeeItems() }}</div> -->

</template>

<script setup>
import { useUserStore } from '../stores/UserStore';
import { useAdminStore } from '../stores/AdminStore';
import { ref } from 'vue';
const emplid = ref("");
const AdminStore = useAdminStore();
const UserStore = useUserStore();
UserStore.getTickets(UserStore.credentials.token);
UserStore.setNoEmployeeTickets();
AdminStore.getUsers(UserStore.credentials.token)
</script>