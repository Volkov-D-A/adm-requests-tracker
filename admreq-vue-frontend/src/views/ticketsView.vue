<template>
    <v-row>
        <v-col cols="9">
            <v-row>
            
                <v-col cols="3"
                    v-for="ticket in UserStore.getUserTickets()"
                    :key="ticket.id"
                >
                    <v-card class="ma-3">
                        <v-card-item>
                            <v-card-subtitle>{{ ticket.userId }}</v-card-subtitle>
                        </v-card-item>

                    <v-card-text>
                        <div>{{ ticket.text }}</div>
                        <br>
                        <div class="align-end">Исполнитель: {{ ticket.employeeId}}</div>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="teal-darken-1" variant="outlined" block class="mt-2">Подробнее</v-btn>
                    </v-card-actions>
                    </v-card>
                </v-col>
            
            </v-row>
                    <!-- <div>{{ UserStore.getUserTickets() }}</div> -->
        </v-col>
        <v-col>
            <div class="pa-3">
                <v-card class="pa-3" variant="elevated" elevation="16" color="teal-lighten-4">
                    <span class="ma-2">Направить обращение:</span><br><br>
                    <v-form fast-fail @submit.prevent="UserStore.createTicket(text, UserStore.credentials.token)">
                        <v-textarea counter auto-grow variant="outlined" background-color="blue-lighten-5" v-model="text" label="Текст обращения"></v-textarea>
                        <v-btn type="submit" color="teal-darken-1" block class="mt-2">Отправить</v-btn>
                    </v-form>
                </v-card>
                
        <!-- <span class="d-flex align-center justify-center ma-3 text-red">{{ AdminStore.usersErrors }}</span> -->
    </div>
        </v-col>
    </v-row>
</template>

<script setup>
import { useUserStore } from '../stores/UserStore';
import { ref } from 'vue';
const UserStore = useUserStore();
const text = ref("");
UserStore.getTickets(UserStore.credentials.token)
</script>